/** 
* plits a string into tokens based on a specified delimiter.
*
* @param str The input string to be split.
* @param delimiter The character used to split the string.
* @param tokens A pointer to an array of strings where the tokens will be stored.
*              The caller is responsible for freeing the allocated memory.
* @param count A pointer to an integer where the number of tokens will be stored.
*/
void split(const char* str, char delimiter, char*** tokens, int* count);
