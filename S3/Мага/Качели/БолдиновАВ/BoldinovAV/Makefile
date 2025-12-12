MCU=atmega32
CC=avr-gcc
OBJDIR=obj
OUTFILE=main.hex

SRC=main.c utils.c
OBJ=$(patsubst %.c,$(OBJDIR)/%.o,$(SRC))

all: $(OUTFILE)

$(OUTFILE): $(OBJ)
    $(CC) -mmcu=$(MCU) -Os -Wl,-Map=$*.map -o $@ $^ -lm

$(OBJDIR)/%.o : %.c
    mkdir -p $(dir $@)
    $(CC) -c -Wall -O2 -std=c99 -mmcu=$(MCU) -I./inc $< -o $@

clean:
    rm -f $(OBJDIR)/*.o *.hex *.map

.PHONY: all clean
