// +build linux,cgo,!ppc64le darwin,cgo

package lrcn

// #include <stdlib.h>
// #include "fpga_version.h"
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/Unknwon/com"
	"github.com/pkg/errors"
	"github.com/rai-project/dlframework/framework/options"
)

type Predictor struct {
	ctx     C.PredictorContext
	options *options.Options
}

func New(opts0 ...options.Option) (*Predictor, error) {
	opts := options.New(opts0...)
	inputFile := string(opts.Graph())
	if !com.IsFile(inputFile) {
		return nil, errors.Errorf("file %s not found", inputFile)
	}

	if opts.OutputNode() == "" {
		return nil, errors.Errorf("expecting a valid (non-empty) output node name")
	}

	inputFileString := C.CString(inputFile)
	defer C.free(unsafe.Pointer(inputFileString))

	ctx := C.NewFPGAPredictor(inputFileString)

	return &Predictor{
		ctx:     ctx,
		options: opts,
	}, nil
}

func prod(arry []uint32) int64 {
	accum := int64(1)
	for _, e := range arry {
		accum *= int64(e)
	}
	return accum
}

func (p *Predictor) Predict(inputLayerName0 string, outputLayerName0 string, input []float32, shape []uint32) (Predictions, error) {
	// log.WithField("input_layer_name", inputLayerName0).
	// 	WithField("output_layer_name", outputLayerName0).
	// 	Info("performing fpga prediction")

	if inputLayerName0 == "" {
		return nil, errors.New("expecting a valid (non-empty) input layer name")
	}

	if outputLayerName0 == "" {
		return nil, errors.New("expecting a valid (non-empty) output layer name")
	}

	inputLayerName := C.CString(inputLayerName0)
	defer C.free(unsafe.Pointer(inputLayerName))

	outputLayerName := C.CString(outputLayerName0)
	defer C.free(unsafe.Pointer(outputLayerName))

	batchSize := int64(p.options.BatchSize())
	shapeLen := prod(shape)
	dataLen := int64(len(input)) / shapeLen
	if batchSize > dataLen {
		padding := make([]float32, (batchSize-dataLen)*shapeLen)
		input = append(input, padding...)
	}

	r := C.PredictFPGA(p.ctx)
	if r == nil {
		return nil, errors.New("failed to perform fpga prediction")
	}
	defer C.free(unsafe.Pointer(r.Array))
	defer C.free(unsafe.Pointer(r))

	predictions := make([]Prediction, r.Length)
	for ii, idx := range carray2slice(r.Array, r.Length) {
		predictions[ii] = Prediction{
			Index: int64(idx),
		}
	}
	return predictions, nil
}

func carray2slice(array *C.int, len0 C.int) []C.int {
	var list []C.int
	len := int(len0)
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&list)))
	sliceHeader.Cap = len
	sliceHeader.Len = len
	sliceHeader.Data = uintptr(unsafe.Pointer(array))
	return list
}

func (p Predictor) Close() {
	C.DeleteFPGAPredictor(p.ctx)
}

// func (p *Predictor) StartProfiling(name, metadata string) error {
// 	cname := C.CString(name)
// 	cmetadata := C.CString(metadata)
// 	defer C.free(unsafe.Pointer(cname))
// 	defer C.free(unsafe.Pointer(cmetadata))
// 	C.FPGAStartProfiling(p.ctx, cname, cmetadata)
// 	return nil
// }

// func (p *Predictor) EndProfiling() error {
// 	C.FPGAEndProfiling(p.ctx)
// 	return nil
// }

// func (p *Predictor) DisableProfiling() error {
// 	C.FPGADisableProfiling(p.ctx)
// 	return nil
// }

// func (p *Predictor) ReadProfile() (string, error) {
// 	cstr := C.FPGAReadProfile(p.ctx)
// 	if cstr == nil {
// 		return "", errors.New("failed to read nil profile")
// 	}
// 	defer C.free(unsafe.Pointer(cstr))
// 	return C.GoString(cstr), nil
// }
