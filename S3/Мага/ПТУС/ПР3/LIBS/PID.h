//
#ifndef PID_H
#define PID_H
//
#include "define.h"

//
#define KP 0.1
#define KI 0.07
#define KD 0.01

#define ACC_MAX 1000
#define ACC_MIN 0

#define REG_MAX 0xFF
#define REG_MIN 0
//
int PID(int param, int reg);

#endif
