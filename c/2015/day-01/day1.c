#include "stdlib.h"
#include <stddef.h>
#include <stdio.h>
#include <string.h>

const char FLOOR_UP = '(';
const char FLOOR_DOWN = ')';
const char EOL = '\n';

// HEADERS
int interpretDirections(char *line, size_t length);

int main(void)
{
	// Open a file
	FILE *fp;
	fp = fopen("./input.txt", "r");

	if (fp == NULL) {
		fprintf(stderr, "Fail to open file");
		return EXIT_FAILURE;
	}

	// Read the line
	char *line = NULL;
	size_t lineSize = 0;
	ssize_t length;

	// keep track of the current floor
	// interpret the floor directions
	int destFloor = 0;

	while (1) {
		length = getline(&line, &lineSize, fp);

		int isEOF = length == -1;
		if (isEOF) break;

		destFloor = interpretDirections(line, length);
		printf("Dest floor: %d\n", destFloor);
	}

	free(line);

	return EXIT_SUCCESS;
}

int interpretDirections(char *line, size_t length)
{
	int currentFloor = 0;

	for (int i = 0; i < length; i++) {
		char instruction = *&line[i];

		int isEOL = instruction == EOL;
		if (isEOL) break;

		if (instruction == FLOOR_UP) {
			currentFloor++;
		} else if (instruction == FLOOR_DOWN) {
			currentFloor--;
		}
	}

	return currentFloor;
}
