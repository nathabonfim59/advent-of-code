#include "../lib/file.h"
#include "../lib/debug.h"
#include "../lib/text.h"
#include <stdio.h>
#include <stdlib.h>

struct Dimentions {
	int length;
	int width;
	int height;
};

struct Dimentions parseLine(char *line);

int main() {
	struct File file = openFile("sample.txt");
	char *line = getLine(&file);
	struct Dimentions dm = parseLine(line);

	dd_struct(dm, length, width, height);

	return 0;
}

struct Dimentions parseLine(char *line) {
	char **parts = NULL;
	int count = 0;

	split(line, 'x', &parts, &count);

	if (count != 3) {
		printf("Invalid line format: %s\n", line);
	}

	struct Dimentions dimentions = {
		.length = atoi(parts[0]),
		.width = atoi(parts[1]),
		.height = atoi(parts[2]),
	};

	return dimentions;
}
