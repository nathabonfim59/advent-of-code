#include <stdio.h>

struct File {
	char *filepath;
	FILE *handle;
};

struct File openFile(char *filepath);

char* getLine(struct File *file);
