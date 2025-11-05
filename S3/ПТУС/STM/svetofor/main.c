//
#include "define.h" //базовые библиотеки и определени€
//
void led_0_ini(void)	//инициализаци€ светодиода
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_OUT;	//настраиваем ножку на выход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_NOPULL; //без подт€жки
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_0;		//нулева€ ножка
	GPIO_Init(GPIOA, &GPIO_InitStructure);			//инициализаци€

	GPIO_SetBits(GPIOA, GPIO_Pin_0);				//установка 1 на выходе
}

void led_1_ini(void)	//инициализаци€ светодиода
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_OUT;	//настраиваем ножку на выход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_NOPULL; //без подт€жки
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_1;		//нулева€ ножка
	GPIO_Init(GPIOA, &GPIO_InitStructure);			//инициализаци€

	GPIO_SetBits(GPIOA, GPIO_Pin_1);				//установка 1 на выходе
}

void led_2_ini(void)	//инициализаци€ светодиода
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_OUT;	//настраиваем ножку на выход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_NOPULL; //без подт€жки
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_2;		//нулева€ ножка
	GPIO_Init(GPIOA, &GPIO_InitStructure);			//инициализаци€

	GPIO_SetBits(GPIOA, GPIO_Pin_2);				//установка 1 на выходе
}

void key_ini_gpiob0 (void)//инициализаци€ кнопки
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOB, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN;	//настраиваем ножку на вход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_UP; //подт€жка к плюсу
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_0;		//перва€ ножка
	GPIO_Init(GPIOB, &GPIO_InitStructure);			//инициализаци€
}

void key_ini_gpiob1 (void)//инициализаци€ кнопки
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOB, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN;	//настраиваем ножку на вход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_UP; //подт€жка к плюсу
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_1;		//перва€ ножка
	GPIO_Init(GPIOB, &GPIO_InitStructure);			//инициализаци€
}

void sleep(int time) {
	int i;
	for(i=0;i<=time;i++){

	}
}

void led0OnTick(int time) {
	GPIO_SetBits(GPIOA, GPIO_Pin_0);
	sleep(time);
	GPIO_ResetBits(GPIOA, GPIO_Pin_0);
	sleep(time);
}

void led1OnTick(int time) {
	GPIO_SetBits(GPIOA, GPIO_Pin_1);
	sleep(time);
	GPIO_ResetBits(GPIOA, GPIO_Pin_1);
	sleep(time);
}

void led2OnTick(int time) {
	GPIO_SetBits(GPIOA, GPIO_Pin_2);
	sleep(time);
	GPIO_ResetBits(GPIOA, GPIO_Pin_2);
}

int main(void)
{
	SystemInit();	//настройка тактировани€ контроллера по умолчанию
	NVIC_PriorityGroupConfig(NVIC_PriorityGroup_4);	//конфигураци€ прерываний
	//
	led_0_ini();
	led_1_ini();
	led_2_ini();
	GPIO_ResetBits(GPIOA, GPIO_Pin_0);
	GPIO_ResetBits(GPIOA, GPIO_Pin_1);
	GPIO_ResetBits(GPIOA, GPIO_Pin_2);
	int a = 10000000;
	int b = 15000000;
	int v = 60000000;
	//
	for(;;) {
		led0OnTick(a);
		led0OnTick(a);
		led0OnTick(a);
		led1OnTick(b);
		led1OnTick(b);
		led2OnTick(v);
		led1OnTick(a);
	}
}

//void led0Logic() {
//	if((GPIO_ReadInputDataBit(GPIOB, GPIO_Pin_0)== 1)||(GPIO_ReadInputDataBit(GPIOB, GPIO_Pin_1)== 1))
//			{
//				GPIO_ResetBits(GPIOA, GPIO_Pin_0);
//			} else {
//				GPIO_SetBits(GPIOA, GPIO_Pin_0);
//			}
//}
//
//void led1Logic() {
//	if((GPIO_ReadInputDataBit(GPIOB, GPIO_Pin_0)== 1)&&(GPIO_ReadInputDataBit(GPIOB, GPIO_Pin_1)== 1))
//			{
//				GPIO_ResetBits(GPIOA, GPIO_Pin_1);
//			} else {
//				GPIO_SetBits(GPIOA, GPIO_Pin_1);
//			}
//}
