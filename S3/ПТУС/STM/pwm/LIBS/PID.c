#include "PID.h"

int PID(int err, int reg)
{
	//static int param_old = 0;
	static int reg_old = 0;
	static int acc = 0;

	int sp, ss, sd;
	 // Пропорциональная компонента
	 sp = KP*err;
	 // Интегральная компонента
	 acc = acc+err; // Накапливаем (суммируем)
	 if(acc<ACC_MIN) acc = ACC_MIN; // Проверяем граничные значение
	 if(acc>ACC_MAX) acc = ACC_MAX;
	 ss = KI*acc;
	 // Дифференциальная компонента
	 sd = KD*(reg-reg_old);
	 reg_old = reg;

	 if ((sp+ss+sd) > REG_MAX)
	 {
		 return REG_MAX;
	 }
	 else if ((sp+ss+sd) < REG_MIN)
	 {
		 return REG_MIN;
	 }
	 return sp+ss+sd;
}
