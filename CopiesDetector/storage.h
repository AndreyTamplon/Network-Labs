#ifndef LAB1__STORAGE_H_
#define LAB1__STORAGE_H_
#include <stdbool.h>
#include <arpa/inet.h>
typedef struct Member{
	char* address;
	long long last_message_time;
} Member;

typedef struct Array{
	Member* members;
	int used;
	int size;
} Array;

bool init_array(Array *a, int initialSize);
bool append_array(Array* a, const Member* element);
void delete_element(Array* a, const Member* element);
void free_array(Array* a);
void print_array(Array* a);
int find_element(const Array* a, const Member* element);
bool set_element_last_message_time(Array* a, const Member* element);
void delete_element_by_id(Array* a, int id);

#endif //LAB1__STORAGE_H_
