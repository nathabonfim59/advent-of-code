#include "debug.h"
#include "ansi.h"
#include <stdlib.h>

void _dd_internal_with_fields(
	const char *struct_name,
	__attribute__((unused)) void *s,
	__attribute__((unused)) size_t size,
	field_pair_t *fields,
	size_t field_count
) {
	printf(ANSI_COLOR_CYAN ANSI_BOLD "%s " ANSI_COLOR_RESET, struct_name);
	printf(ANSI_COLOR_YELLOW "{\n" ANSI_COLOR_RESET);
	
	for (size_t i = 0; i < field_count; i++) {
		printf(ANSI_COLOR_GREEN "  %s: " ANSI_COLOR_RESET, fields[i].name);
		// Print the value from the union
		printf(ANSI_COLOR_MAGENTA "%d\n" ANSI_COLOR_RESET, fields[i].value.i);
	}
	
	printf(ANSI_COLOR_YELLOW "}\n" ANSI_COLOR_RESET);
	exit(1);
}
