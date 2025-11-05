//
#ifndef PID_H
#define PID_H
//
#include "define.h"

//
//#define KP 0.1
//#define KI 0.07
//#define KD 0.01

#define KP 0.05
#define KI 0.025
#define KD 0.025
// Увеличение K_i помогает компенсировать ошибку накопления и улучшает стабильность.
//Снижение K_p снижает реакцию на быстрые изменения, что может помочь избежать скачков скорости.
//K_d помогает предотвратить колебания, но слишком большой K_d может привести к нестабильности.

#define ACC_MAX 2500
#define ACC_MIN 0

#define REG_MAX 0xFF
#define REG_MIN 0
//
int PID(int param, int reg);

#endif
