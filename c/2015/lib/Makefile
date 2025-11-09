CC = gcc
CFLAGS = -Wall -Wextra -std=c99 -D_GNU_SOURCE
LIB_DIR = ../lib
LIB_SOURCES = $(LIB_DIR)/file.c $(LIB_DIR)/debug.c $(LIB_DIR)/text.c
LIB_OBJECTS = $(LIB_SOURCES:.c=.o)

.PHONY: clean-lib

clean-lib:
	rm -f $(LIB_OBJECTS)

$(LIB_DIR)/%.o: $(LIB_DIR)/%.c
	$(CC) $(CFLAGS) -c $< -o $@
