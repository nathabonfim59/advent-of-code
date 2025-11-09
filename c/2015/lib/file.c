#define _GNU_SOURCE
#include "file.h"

struct File openFile(char *filepath)
{
	struct File file;
	file.filepath = filepath;
	file.handle = fopen(filepath, "r");

	if (file.handle == NULL) {
		fprintf(stderr, "Fail to open file: %s\n", filepath);
	}

	return file;
}
	
char* getLine(struct File *file)
{
	char *line = NULL;
	size_t lineSize = 0;

	ssize_t length = getline(&line, &lineSize, file->handle);

	return line;
}
