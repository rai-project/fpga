// +build linux,cgo,!ppc64le darwin,cgo

package lrcn

// #cgo LDFLAGS: -L${SRCDIR} -lstdc++
// #cgo CXXFLAGS: -std=c++11  -O3 -Wall -g
// #cgo CXXFLAGS: -Wno-sign-compare -Wno-unused-function -I${SRCDIR}/include
import "C"
