#ifndef FPGA_VERSION_H
#define FPGA_VERSION_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

typedef void *PredictorContext;

typedef struct {
  int Length;
  int *Array;
} Predictions;

PredictorContext NewFPGAPredictor(const char *input_file);
void DeleteFPGAPredictor(PredictorContext pred);

Predictions *PredictFPGA(PredictorContext pred);

#ifdef __cplusplus
}
#endif // __cplusplus

#endif
