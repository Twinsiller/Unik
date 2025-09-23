#include "stm32f4xx.h"
#include "stm32f4xx_gpio.h"
#include "stm32f4xx_rcc.h"

void GPIO_Config(void);

int main(void)
{
    GPIO_Config();

    while (1)
    {
        if (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_0) == Bit_SET) // SW1
        {
            GPIO_SetBits(GPIOD, GPIO_Pin_12);  // LED1
        }
        else
        {
            GPIO_ResetBits(GPIOD, GPIO_Pin_12);
        }
    }
}

void GPIO_Config(void)
{
    GPIO_InitTypeDef GPIO_InitStruct;

    // �������� ������������ ������ A � D
    RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);
    RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOD, ENABLE);

    // ��������� SW1 (PA0) ��� �����
    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_0;
    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_IN;
    GPIO_InitStruct.GPIO_PuPd = GPIO_PuPd_DOWN; // �������� ����, ���� ������ �������� �� VCC
    GPIO_Init(GPIOA, &GPIO_InitStruct);

    // ��������� LED1 (PD12) ��� ������
    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_12;
    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_OUT;
    GPIO_InitStruct.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_InitStruct.GPIO_OType = GPIO_OType_PP;
    GPIO_Init(GPIOD, &GPIO_InitStruct);
}



//#include "stm32f4xx.h"
//#include "stm32f4xx_gpio.h"
//#include "stm32f4xx_rcc.h"
//
//void GPIO_Config(void);
//
//int main(void)
//{
//    GPIO_Config();
//
//    while (1)
//    {
//        if (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_1) == Bit_SET) // SW1
//        {
//            GPIO_SetBits(GPIOA, GPIO_Pin_0);  // LED1
//        }
//        else
//        {
//            GPIO_ResetBits(GPIOA, GPIO_Pin_0);
//        }
//    }
//}
//
//void GPIO_Config(void)
//{
//    GPIO_InitTypeDef GPIO_InitStruct;
//
//    // �������� ������������ ������ A
//    RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);
//
//    // ��������� SW1 (PA1) ��� �����
//    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_1;
//    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_IN;
//    GPIO_InitStruct.GPIO_PuPd = GPIO_PuPd_DOWN; // �������� ����, ���� ������ �������� �� VCC
//    GPIO_Init(GPIOA, &GPIO_InitStruct);
//
//    // ��������� LED1 (PA0) ��� ������
//    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_0;
//    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_OUT;
//    GPIO_InitStruct.GPIO_Speed = GPIO_Speed_50MHz;
//    GPIO_InitStruct.GPIO_OType = GPIO_OType_PP;
//    GPIO_Init(GPIOA, &GPIO_InitStruct);
//}
