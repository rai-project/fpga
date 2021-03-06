#ifndef FPGA_VERSION_H
#define FPGA_VERSION_H
#include <stdio.h>
#include <stdlib.h>
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <cstring>

#include <sys/mman.h>
#include <sys/stat.h>
#include <sys/time.h>
#include <sys/types.h>
#include <unistd.h>

#include <assert.h>
#include <fcntl.h>
#include <getopt.h>
#include <stdint.h>
#include <time.h>

#include <chrono> // for high_resolution_clock
#include <ap_fixed.h>
class Predictor {
public:
  int from_fpga_fd_;
  int to_fpga_fd_;
  Predictor(int from_fpga_fd, int to_fpga_fd){
	from_fpga_fd_ = from_fpga_fd;
	to_fpga_fd_ = to_fpga_fd;
  };
  ~Predictor() {
  }
};

Predictor* Newfpga(const char input_file[]);
void Deletefpga(Predictor* pred);

int *Predictfpga(Predictor* pred, const char input_image[]) ;
=======

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

