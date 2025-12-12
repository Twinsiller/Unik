#include <iom32v.h>
#include <macros.h>

void port_init(void)
{
 DDRA = 0x00;
 PORTD= 0x00;
 DDRD = 0xFF;
}

void adc_init(void)
{
 ADMUX = 0x03;
 
 ACSR = 0x80;
 ADCSRA = 0xCE;
}

#pragma interrupt_handler adc_isr:17
void adc_isr(void)
{
 int v;
 CLI();
 v = ADCL;
 v|=(int)ADCH <<8;
 v>>=2;
 
 PORTD=v;
 ADCSRA|=0x40;
 SEI();
}

void init_devices(void)
{
 CLI();
 port_init();
 adc_init();
 SEI();
}

void main(void)
{
 init_devices();
 while(1);
}
