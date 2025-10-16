//
#include "Led7_emulator.h"
#include "GPIO_LIBS.h"
#include "define.h"
#include "delay.h"
#include <string.h>
#include <stdlib.h>
//
void AD5290_Set(uint16_t data)//
{
	if (data > 255) data = 255;
	GPIO_ResetBits(GPIOB, GPIO_Pin_1);//�������� CS
	delay_us(2);
	while(SPI_I2S_GetFlagStatus(SPI2, SPI_I2S_FLAG_TXE) == RESET);//��������, ���� ������� �������� �� ��������
	SPI_I2S_SendData(SPI2, data);//���������� ������
	while (SPI_I2S_GetFlagStatus(SPI2, SPI_I2S_FLAG_RXNE) == RESET);//��������, ���� ������� �������� �� ����������
	SPI_I2S_ReceiveData(SPI2);//���������� ������
	delay_us(2);
	GPIO_SetBits(GPIOB, GPIO_Pin_1);//��������� CS
}
//
void AD5290_Ini(void)//��������� SPI ��� ������ � ���������� ����������� ��������������
{
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOB, ENABLE); //�������� ������������ ������� �����
	RCC_APB1PeriphClockCmd(RCC_APB1Periph_SPI2, ENABLE); //�������� ������������ SPI2
	//
	GPIO_InitTypeDef GPIO_InitStructure;
	////////////////////////////////////////////
	GPIO_PinAFConfig(GPIOB, GPIO_PinSource13, GPIO_AF_SPI2);//SCK
	GPIO_PinAFConfig(GPIOB, GPIO_PinSource14, GPIO_AF_SPI2);//MISO
	GPIO_PinAFConfig(GPIOB, GPIO_PinSource15, GPIO_AF_SPI2);//MOSI
	///
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AF;
	GPIO_InitStructure.GPIO_Speed = GPIO_Speed_100MHz;
	GPIO_InitStructure.GPIO_OType = GPIO_OType_PP;
	GPIO_InitStructure.GPIO_PuPd  = GPIO_PuPd_UP;
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_13 | GPIO_Pin_14 | GPIO_Pin_15;
	GPIO_Init(GPIOB, &GPIO_InitStructure);
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_OUT;
	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_1;//CS
	GPIO_Init(GPIOB, &GPIO_InitStructure);
	GPIO_SetBits(GPIOB, GPIO_Pin_1);
	//

	SPI_InitTypeDef  SPI_InitStructure;
	// SPI-Konfiguration
	SPI_InitStructure.SPI_Direction = SPI_Direction_2Lines_FullDuplex;//������ � ��������������� ������
	SPI_InitStructure.SPI_Mode = SPI_Mode_Master;//������ ��� ������
	SPI_InitStructure.SPI_DataSize = SPI_DataSize_8b;//�������� 8bit �� ���� �������
	SPI_InitStructure.SPI_CPOL = SPI_CPOL_Low;//spi mode 0
	SPI_InitStructure.SPI_CPHA = SPI_CPHA_1Edge;//spi mode 0
	SPI_InitStructure.SPI_NSS = SPI_NSS_Soft;//����������� ���������� CS
	SPI_InitStructure.SPI_BaudRatePrescaler = SPI_BaudRatePrescaler_16;//42MHZ/16 = 2.625MHZ  ������� ������������ SPI
	SPI_InitStructure.SPI_FirstBit = SPI_FirstBit_MSB;//������� ��� ������
	SPI_InitStructure.SPI_CRCPolynomial = 7;//����������� �����
	SPI_Init(SPI2, &SPI_InitStructure);
	SPI_Cmd(SPI2, ENABLE);
	SPI_NSSInternalSoftwareConfig(SPI2, SPI_NSSInternalSoft_Set);//����������� ���������� CS
	//
	AD5290_Set(128);
}
//
void ADC_Injected_ini(void) //��������� ���
{
	GPIO_InitTypeDef      GPIO_InitStructure;
	ADC_CommonInitTypeDef ADC_CommonInitStructure;
	ADC_InitTypeDef       ADC_InitStructure;

	// Clock Enable
	RCC_APB2PeriphClockCmd(RCC_APB2Periph_ADC1, ENABLE);//�������� ������������ ���1
	RCC_AHB1PeriphClockCmd(RCC_AHB1Periph_GPIOA, ENABLE);

	GPIO_InitStructure.GPIO_Pin = GPIO_Pin_0 | GPIO_Pin_1 | GPIO_Pin_2;//0 �����, ������ ���� ���, 1 �����, ������ ���� ���, 2 �����, ������ ���� ���
	GPIO_InitStructure.GPIO_Mode = GPIO_Mode_AN;//���������� ����
	GPIO_InitStructure.GPIO_PuPd = GPIO_PuPd_NOPULL ;//��� ��������
	GPIO_Init(GPIOA, &GPIO_InitStructure);
	// ADC-Config
	ADC_CommonInitStructure.ADC_Mode = ADC_Mode_Independent;//����������� ������ ���
	ADC_CommonInitStructure.ADC_Prescaler = ADC_Prescaler_Div2;//��������� ������������ ���
	ADC_CommonInitStructure.ADC_DMAAccessMode = ADC_DMAAccessMode_Disabled;//��� ���
	ADC_CommonInitStructure.ADC_TwoSamplingDelay = ADC_TwoSamplingDelay_5Cycles;//������ ���������
	ADC_CommonInit(&ADC_CommonInitStructure);

	ADC_InitStructure.ADC_Resolution = ADC_Resolution_12b;//���������� 12 ���(0 - 4095)
	ADC_InitStructure.ADC_ScanConvMode = ENABLE;//
	ADC_InitStructure.ADC_ContinuousConvMode = ENABLE;
	ADC_InitStructure.ADC_ExternalTrigConvEdge = ADC_ExternalTrigConvEdge_None;//��� ��������
	ADC_InitStructure.ADC_DataAlign = ADC_DataAlign_Right;//������������ �� ������� ����
	ADC_InitStructure.ADC_NbrOfConversion = 1;
	ADC_Init(ADC1, &ADC_InitStructure);
	// ADC-Enable
	ADC_InjectedSequencerLengthConfig(ADC1,3);
	ADC_InjectedChannelConfig(ADC1,ADC_Channel_0, 1, ADC_SampleTime_3Cycles);
	ADC_InjectedChannelConfig(ADC1,ADC_Channel_1, 2, ADC_SampleTime_3Cycles);
	ADC_InjectedChannelConfig(ADC1,ADC_Channel_2, 3, ADC_SampleTime_3Cycles);
	//
	ADC_Cmd(ADC1, ENABLE);
	ADC_AutoInjectedConvCmd(ADC1,ENABLE);
	ADC_SoftwareStartConv(ADC1);
}
//
uint16_t Read_ADC(int num)//������ ������ � ���
{
	  if(num==0) return ADC_GetInjectedConversionValue(ADC1,ADC_InjectedChannel_1);
	  else if(num==1) return ADC_GetInjectedConversionValue(ADC1,ADC_InjectedChannel_2);
	  else if(num==2) return ADC_GetInjectedConversionValue(ADC1,ADC_InjectedChannel_3);
	  return 0;
}
//
int main(void)
{
	uint16_t data0 = 0, data1 = 0, data2 = 0;
	char str[7];
	SystemInit();
	NVIC_PriorityGroupConfig(NVIC_PriorityGroup_4);
	//
	TFT_Init(3);
	TFT_Fill_Screen(0,479,0, 319, BLACK);
	delay_Init();
	AD5290_Ini();

	//Mode_pins_ini();
	ADC_Injected_ini();

	for(;;)

	{
		data0 = Read_ADC(0);//���������� ���������� � �������������
		sprintf(str, "%5d", data0);
		Led7_em_str(str, 0);
		uint8_t to_hell = data0>>4;
		AD5290_Set(to_hell);//������������� ��������������� ���������� �� ���������(� ������ ������ �����������)
		delay_ms(100);
		//
		data1 = Read_ADC(1);//������ ���������� � ���������
		sprintf(str, "%5d", data1);
		Led7_em_str(str, 1);
		delay_ms(100);
		//
		data2 = Read_ADC(2);//������ ��� � ���������
		sprintf(str, "%5d", data2);
		Led7_em_str(str, 2);
		//
		delay_ms(100);

	}
}
