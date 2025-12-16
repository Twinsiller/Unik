#include "stm32f4xx.h"
#include "stm32f4xx_gpio.h"
#include "stm32f4xx_rcc.h"

void GPIO_Config(void);

int main(void)
{
    GPIO_Config();

    while (1)
    {

    	GPIO_SetBits(GPIOA, GPIO_Pin_3);  // LED1

//        if (GPIO_ReadInputDataBit(GPIOA, GPIO_Pin_0) == Bit_SET) // SW1
//        {
//            GPIO_SetBits(GPIOD, GPIO_Pin_12);  // LED1
//        }
//        else
//        {
//            GPIO_ResetBits(GPIOD, GPIO_Pin_12);
//        }
    }
}

void GPIO_Config(void)
{
    GPIO_InitTypeDef GPIO_InitStruct;

    // Включаем тактирование портов A и D
    //RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);
    RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);

//    // Настройка SW1 (PA0) как входа
//    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_0;
//    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_IN;
//    GPIO_InitStruct.GPIO_PuPd = GPIO_PuPd_DOWN; // подтяжка вниз, если кнопка замыкает на VCC
//    GPIO_Init(GPIOA, &GPIO_InitStruct);

    // Настройка LED1 (PD1) как выхода
    GPIO_InitStruct.GPIO_Pin = GPIO_Pin_1|GPIO_Pin_3|GPIO_Pin_7;
    GPIO_InitStruct.GPIO_Mode = GPIO_Mode_OUT;
    GPIO_InitStruct.GPIO_Speed = GPIO_Speed_50MHz;
    GPIO_InitStruct.GPIO_OType = GPIO_OType_PP;
    GPIO_Init(GPIOA, &GPIO_InitStruct);

}
