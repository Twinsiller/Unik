#include "Led7_emulator.h"
#include "stm32f4xx_fsmc.h"
//
#define SIZE_X	40
#define SIZE_Y	40
//
void Draw_char(char data, int pos_x, int pos_y);
//
void Led7_em_str(char *str, int num)
{
	int i = 0;
	while(str[i] != 0)
	{
		Draw_char(str[i], i, num);
		i++;
		if(i > 7) break;
	}
}
//
void Draw_char(char data, int pos_x, int pos_y)
{
	TFT_Fill_Rectangle(pos_x *40, pos_y * 80, 40, 80, BLACK);
	switch (data)
	{
		case '0':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '1':
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '2':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			//TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '3':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '4':
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '5':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			//TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '6':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			//TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '7':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			//TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '8':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		case '9':
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 28, 10, WHITE);//1
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 28+1, 10, WHITE);//2
			TFT_Draw_Vertical_Line(pos_x *40 + 15+10+1, pos_y * 80 + 40 + 1, 10, WHITE);//3
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40 + 1 + 10 + 1, 10, WHITE);//4
			//TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 40 + 1, 10, WHITE);//5
			TFT_Draw_Vertical_Line(pos_x *40 + 15 - 1, pos_y * 80 + 28+1, 10, WHITE);//6
			TFT_Draw_Horizontal_Line(pos_x *40 + 15, pos_y * 80 + 40, 10, WHITE);//7
			break;
		default:
			break;
	}
}
