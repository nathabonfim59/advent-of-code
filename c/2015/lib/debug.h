#include <stddef.h>
#include <stdio.h>

// Field pair structure for generic struct dumping
typedef struct {
    const char *name;
    // We'll use a union to handle different types
    union {
        int i;
        unsigned int u;
        float f;
        double d;
        char c;
        char *s;
    } value;
} field_pair_t;

// Generic field printer using _Generic
#define PRINT_FIELD(value) _Generic((value), \
    int: printf(ANSI_COLOR_MAGENTA "%d" ANSI_COLOR_RESET, value), \
    unsigned int: printf(ANSI_COLOR_MAGENTA "%u" ANSI_COLOR_RESET, value), \
    short: printf(ANSI_COLOR_MAGENTA "%hd" ANSI_COLOR_RESET, value), \
    unsigned short: printf(ANSI_COLOR_MAGENTA "%hu" ANSI_COLOR_RESET, value), \
    long: printf(ANSI_COLOR_MAGENTA "%ld" ANSI_COLOR_RESET, value), \
    unsigned long: printf(ANSI_COLOR_MAGENTA "%lu" ANSI_COLOR_RESET, value), \
    long long: printf(ANSI_COLOR_MAGENTA "%lld" ANSI_COLOR_RESET, value), \
    unsigned long long: printf(ANSI_COLOR_MAGENTA "%llu" ANSI_COLOR_RESET, value), \
    float: printf(ANSI_COLOR_MAGENTA "%f" ANSI_COLOR_RESET, value), \
    double: printf(ANSI_COLOR_MAGENTA "%f" ANSI_COLOR_RESET, value), \
    char: printf(ANSI_COLOR_MAGENTA "'%c'" ANSI_COLOR_RESET, value), \
    char*: printf(ANSI_COLOR_MAGENTA "\"%s\"" ANSI_COLOR_RESET, value), \
    default: printf(ANSI_COLOR_MAGENTA "??" ANSI_COLOR_RESET) \
)

// Macro to create field name-value pairs
#define FIELD_PAIR(struct_var, field) {#field, {.i = struct_var.field}}

// Helper macros for variadic field mapping
#define _GET_NTH_ARG(_1, _2, _3, _4, _5, _6, _7, _8, _9, _10, _11, _12, _13, _14, _15, _16, N, ...) N
#define _COUNT_ARGS(...) _GET_NTH_ARG(__VA_ARGS__, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1)

#define _MAP_FIELD_PAIR_1(var, f1) FIELD_PAIR(var, f1)
#define _MAP_FIELD_PAIR_2(var, f1, f2) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2)
#define _MAP_FIELD_PAIR_3(var, f1, f2, f3) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3)
#define _MAP_FIELD_PAIR_4(var, f1, f2, f3, f4) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3), FIELD_PAIR(var, f4)
#define _MAP_FIELD_PAIR_5(var, f1, f2, f3, f4, f5) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3), FIELD_PAIR(var, f4), FIELD_PAIR(var, f5)
#define _MAP_FIELD_PAIR_6(var, f1, f2, f3, f4, f5, f6) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3), FIELD_PAIR(var, f4), FIELD_PAIR(var, f5), FIELD_PAIR(var, f6)
#define _MAP_FIELD_PAIR_7(var, f1, f2, f3, f4, f5, f6, f7) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3), FIELD_PAIR(var, f4), FIELD_PAIR(var, f5), FIELD_PAIR(var, f6), FIELD_PAIR(var, f7)
#define _MAP_FIELD_PAIR_8(var, f1, f2, f3, f4, f5, f6, f7, f8) FIELD_PAIR(var, f1), FIELD_PAIR(var, f2), FIELD_PAIR(var, f3), FIELD_PAIR(var, f4), FIELD_PAIR(var, f5), FIELD_PAIR(var, f6), FIELD_PAIR(var, f7), FIELD_PAIR(var, f8)

// Need extra indirection for macro concatenation to work
#define _MAP_FIELD_PAIRS_DISPATCH(N, var, ...) _MAP_FIELD_PAIR_##N(var, __VA_ARGS__)
#define _MAP_FIELD_PAIRS_IMPL2(N, var, ...) _MAP_FIELD_PAIRS_DISPATCH(N, var, __VA_ARGS__)
#define _MAP_FIELD_PAIRS(var, ...) _MAP_FIELD_PAIRS_IMPL2(_COUNT_ARGS(__VA_ARGS__), var, __VA_ARGS__)

// Convenience macro: dd_struct(variable, field1, field2, ...)
#define dd_struct(var, ...) dd(var, _MAP_FIELD_PAIRS(var, __VA_ARGS__))

// Generic dd macro - pass struct and fields as pairs
#define dd(struct_var, ...) _dd_internal_with_fields(#struct_var, &(struct_var), sizeof(struct_var), (field_pair_t[]){__VA_ARGS__}, sizeof((field_pair_t[]){__VA_ARGS__})/sizeof(field_pair_t))

// You can create convenience macros for your struct types like this:
// #define DD_POINT(struct_var) dd(struct_var, FIELD_PAIR(struct_var, x), FIELD_PAIR(struct_var, y))
// #define DD_PERSON(struct_var) dd(struct_var, FIELD_PAIR(struct_var, name), FIELD_PAIR(struct_var, age), FIELD_PAIR(struct_var, height))

void _dd_internal_with_fields(const char *struct_name, void *s, size_t size, field_pair_t *fields, size_t field_count);
