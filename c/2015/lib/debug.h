#include <stddef.h>
#include <stdio.h>

void print_struct(void *s, size_t size, const char **field_names, size_t field_count);
void _dd_internal(const char *struct_name, void *s, size_t size, const char **field_names, size_t field_count);

#define dd(struct_var) _dd_internal(#struct_var, &(struct_var), sizeof(struct_var), (const char*[]){"length", "width", "height"}, 3)
