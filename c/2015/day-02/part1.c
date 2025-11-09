#include "../lib/file.h"
#include "../lib/debug.h"

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

	dd_struct(dm, length);

	return 0;
}

struct Dimentions parseLine(char *line) {
	struct Dimentions dimentions = {
		.height = 0,
		.length = 1,
		.width = 0
	};

	return dimentions;
}
