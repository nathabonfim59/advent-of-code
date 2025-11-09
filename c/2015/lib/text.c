#include "text.h"
#include <stdlib.h>
#include <string.h>

void split(const char *str, char delimiter, char ***tokens, int *count) {
	char **result = NULL;
	int tokensCount = 0;
	const char *start = str;
	const char *end = str;

	while (*end != '\0') {
		if (*end == delimiter) {
			size_t length = end - start;
			char *token = (char *)malloc(length + 1);
			strncpy(token, start, length);
			token[length] = '\0';

			result = (char **)realloc(result, sizeof(char *) * (tokensCount + 1));
			result[tokensCount++] = token;

			start = end + 1;
		}
		end++;
	}

	// Add the last token
	if (start != end) {
		size_t length = end - start;
		char *token = (char *)malloc(length + 1);
		strncpy(token, start, length);
		token[length] = '\0';

		result = (char **)realloc(result, sizeof(char *) * (tokensCount + 1));
		result[tokensCount++] = token;
	}

	*tokens = result;
	*count = tokensCount;
}
