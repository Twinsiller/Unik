//
#include "define.h"
//
#include "delay.h"
#include "GPIO_LIBS.h"
#include "PID.h"
#include "Led7_emulator.h"
//

#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <math.h>
//
#define ENC1	GPIOC, GPIO_Pin_8
#define ENC2	GPIOC, GPIO_Pin_9
#define ENC3	GPIOC, GPIO_Pin_10

int SQ1 = 0;
int Et = 1000;
int Et1 = 1400;
int Et2 = 1100;
int Et3 = 700;
int stage = 1;
int pwm = 0;

int S = 0;
//
void Set_pwm(uint8_t data)
{
	TIM4->CCR1 = data;
	pwm = data;
}
//
void In_ini(void)
{
	//GPIO_SetBits(PIN_CS1);
	//
	NVIC_InitTypeDef NVIC_InitStructure;
	NVIC_InitStructure.NVIC_IRQChannel = TIM7_IRQn;
	NVIC_InitStructure.NVIC_IRQChannelPreemptionPriority = 0x0A;
	NVIC_InitStructure.NVIC_IRQChannelSubPriority = 0x00;
	NVIC_InitStructure.NVIC_IRQChannelCmd = ENABLE;
	NVIC_Init(&NVIC_InitStructure);
	//
	RCC_APB1PeriphClockCmd(RCC_APB1Periph_TIM7, ENABLE);
	TIM_TimeBaseInitTypeDef TIM_TimeBaseStructure;
	TIM_TimeBaseStructInit(&TIM_TimeBaseStructure);
	TIM_TimeBaseStructure.TIM_Prescaler = 84-1;
	TIM_TimeBaseStructure.TIM_Period = 10 - 1; //
	TIM_TimeBaseStructure.TIM_ClockDivision = TIM_CKD_DIV1;
	TIM_TimeBaseStructure.TIM_CounterMode = TIM_CounterMode_Up;
	TIM_TimeBaseInit(TIM7, &TIM_TimeBaseStructure);
	TIM_ITConfig(TIM7, TIM_IT_Update, ENABLE); // включаем прерывани€ по переполнению
	TIM_Cmd(TIM7, ENABLE);
	//
	GPIO_Ini(ENC1, GPIO_Mode_IN, GPIO_OType_PP, GPIO_Speed_100MHz, GPIO_PuPd_NOPULL);
	GPIO_Ini(ENC2, GPIO_Mode_IN, GPIO_OType_PP, GPIO_Speed_100MHz, GPIO_PuPd_UP);
//	GPIO_Ini(ENC2, GPIO_Mode_IN, GPIO_OType_PP, GPIO_Speed_100MHz, GPIO_PuPd_NOPULL);
//	GPIO_Ini(ENC3, GPIO_Mode_IN, GPIO_OType_PP, GPIO_Speed_100MHz, GPIO_PuPd_NOPULL);
}
//
void TIM7_IRQHandler(void)
{
	TIM_ClearITPendingBit(TIM7, TIM_IT_Update);
	static int NNN1 = 0;
	static int sch = 0;
	static int pin_old1 =  0;

	int pin1 = GPIO_ReadInputDataBit(ENC1);
	if ((pin1 != pin_old1) && (pin_old1 != 0))  ++NNN1;
	pin_old1 = pin1;

	if(++sch > 100000)
	{
		sch = 0;
		SQ1 = NNN1;
		NNN1 = 0;
		//
		pwm = PID(Et - SQ1, pwm);
		Set_pwm(pwm);
	}
}
//
void pwm_init(void)
{
	GPIO_Ini(GPIOD, GPIO_Pin_12,GPIO_Mode_AF,GPIO_OType_PP,GPIO_Speed_100MHz,GPIO_PuPd_UP);
	//
	GPIO_PinAFConfig(GPIOD, GPIO_PinSource12, GPIO_AF_TIM4);
	//
	TIM_TimeBaseInitTypeDef  TIM_TimeBaseStructure;
	TIM_OCInitTypeDef  TIM_OCInitStructure;
	// Clock enable
	  RCC_APB1PeriphClockCmd(RCC_APB1Periph_TIM4, ENABLE);
	  //7TIM_TimeBaseStructure(&TIM_TimeBaseStructure);
	  TIM_TimeBaseStructure.TIM_Prescaler = 84-1; //предделитель
	  TIM_TimeBaseStructure.TIM_Period = 0xFF; //определ€ет максимальное значение таймера
	  TIM_TimeBaseStructure.TIM_CounterMode = TIM_CounterMode_CenterAligned1;
	  TIM_TimeBaseInit(TIM4, &TIM_TimeBaseStructure);
	  //
	  TIM_OCStructInit(&TIM_OCInitStructure);
	  // PWM on Ch1 Init
	  TIM_OCInitStructure.TIM_OCMode = TIM_OCMode_PWM1;
	  TIM_OCInitStructure.TIM_OCPolarity = TIM_OCPolarity_High;
	  TIM_OCInitStructure.TIM_OutputState = TIM_OutputState_Enable;
	  TIM_OCInitStructure.TIM_Pulse = 0x0;
	  //
	  TIM_OC1Init(TIM4, &TIM_OCInitStructure);
	  //
	  TIM_OC1PreloadConfig(TIM4, TIM_OCPreload_Enable);
	  TIM_ARRPreloadConfig(TIM4, ENABLE);
	  //
	  TIM_Cmd(TIM4, ENABLE);
}

void keys_ini(void)//инициализаци€ кнопки
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);	//включение тактировани€ соответствующего порта
	//
	GPIO_InitTypeDef GPIO_InitStructure;	//объ€вление структуры
	//
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_IN;	//настраиваем ножку на вход
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;	//настраиваем ножку как Push-Pull
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz; //частота тактировани€ 100MHz
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_UP; //подт€жка к плюсу
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_1|GPIO_Pin_2;		//перва€ ножка
	GPIO_Init(GPIOA, &GPIO_InitStructure);			//инициализаци€
}

void dp(char *str)
{
	sprintf(str, "%6d", pwm);
	Led7_em_str(str, 0);//выводим
	sprintf(str, "%6d", SQ1);
	Led7_em_str(str, 1);//выводим
	sprintf(str, "%6d", Et);
	Led7_em_str(str, 2);//выводим
	sprintf(str, "%6d", stage);
	Led7_em_str(str, 3);//выводим

	delay_ms(500);
}
//
int main(void)
{
	char str[16];
	SystemInit();
	keys_ini();
	NVIC_PriorityGroupConfig(NVIC_PriorityGroup_4);
	//
	pwm_init();
	In_ini();
	TFT_Init(3);
	TFT_Fill_Screen(0,479,0, 319, BLACK);
	delay_Init();
	Et = Et1;
	//
	while(1)
	{
//		int er = GPIO_ReadInputDataBit(ENC2);
//		if(er)
//		{
//			Et = Et1;
//		}
//		else
//		{
//			Et = Et2;
//		}
//		dp(str);

		int er = GPIO_ReadInputDataBit(ENC2);
		if(er){
			if((GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_1) == 0) && (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_2) == 1)){
				Et = 1000;
				stage = 2;
			}
			else if ((GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_1) == 1) && (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_2) == 0)){
				Et = 1200;
				stage = 3;
			}
			else if ((GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_1) == 1) && (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_2) == 1)){
				Et = 1600;
				stage = 4;
			}
			else{
				Et = 800;
				stage = 1;
			}
		}
		else
		{
			Et = Et2;
			stage = 9;
		}
		dp(str);
	}
}
