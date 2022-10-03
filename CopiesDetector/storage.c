#include "storage.h"

#include "malloc.h"
#include "string.h"
#include <stdbool.h>

bool init_array(Array* a, int initialSize) {
	a->members = (Member*)malloc(initialSize * sizeof(Member));
	if (a->members == NULL) {
		return false;
	}
	a->used = 0;
	a->size = initialSize;
	return true;
}

bool append_array(Array* a, const Member* element) {
	if (a->used == a->size) {
		a->size *= 2;
		a->members = (Member*)realloc(a->members, a->size * sizeof(Member));
		if (a->members) {
			a->members[a->used++] = *element;
			return true;
		}
		else {
			return false;
		}
	}
	else {
		a->members[a->used++] = *element;
		return true;
	}
}

int find_element(const Array* a, const Member* element) {
	for (int i = 0; i < a->used; i++) {
		if(strcmp(a->members[i].address, element->address) == 0) {
			return i;
		}
	}
	return -1;
}

bool set_element_last_message_time(Array* a, const Member* element) {
	for (int i = 0; i < a->used; i++) {
		if(strcmp(a->members[i].address, element->address) == 0) {
			a->members[i].last_message_time = element->last_message_time;
			return true;
		}
	}
	return false;
}

void delete_element(Array* a, const Member* element) {
	for (int i = 0; i < a->used; i++) {
		if (strcmp(a->members[i].address, element->address) == 0) {
			for (int j = i; j < a->used - 1; j++) {
				a->members[j] = a->members[j + 1];
			}
			a->used--;
			break;
		}
	}
}

void delete_element_by_id(Array* a, int id) {
	for (int i = id; i < a->used - 1; i++) {
		a->members[i] = a->members[i + 1];
	}
	a->used--;
}

void print_array(Array* a) {
	for (int i = 0; i < a->used; i++) {
		printf("%d %s", i, a->members[i].address);
	}
}

void free_array(Array* a) {
	free(a->members);
	a->members = NULL;
	a->used = a->size = 0;
}
