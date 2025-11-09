#include "debug.h"
#include "ansi.h"
#include <stdlib.h>

void print_struct(
	void *s,
	__attribute__((unused)) size_t size,
	const char **field_names,
	size_t field_count
) {
	unsigned char *ptr = (unsigned char *)s;
	printf(ANSI_COLOR_YELLOW "{\n" ANSI_COLOR_RESET);
	for (size_t i = 0; i < field_count; i++) {
		printf(ANSI_COLOR_GREEN "  %s: " ANSI_COLOR_RESET, field_names[i]);
		// Print as int for simplicity
		int value = *(int *)(ptr + i * sizeof(int));
		printf(ANSI_COLOR_MAGENTA "%d\n" ANSI_COLOR_RESET, value);
	}
	printf(ANSI_COLOR_YELLOW "}\n" ANSI_COLOR_RESET);
}

void _dd_internal(
	const char *struct_name,
	void *s,
	__attribute__((unused)) size_t size,
	const char **field_names,
	size_t field_count
) {
	printf(ANSI_COLOR_CYAN ANSI_BOLD "%s " ANSI_COLOR_RESET, struct_name);
	print_struct(s, size, field_names, field_count);
	exit(1);
}
