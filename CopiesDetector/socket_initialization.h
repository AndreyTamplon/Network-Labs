#ifndef LAB1__SOCKET_INITIALIZATION_H_
#define LAB1__SOCKET_INITIALIZATION_H_
#include <stdio.h>
#include <sys/socket.h>
#include <netinet/ip.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include "errno.h"
#include "ip_info.h"
typedef struct NetworkInformation {
	char* interface_name;
	char* group_address;
	int port;
} NetworkInformation;

void init_listening_socket(int* listening_socket,
						   struct sockaddr_in* listen_address,
						   const NetworkInformation* network_information);
void init_send_socket(int* send_socket, struct sockaddr_in* group_address, const NetworkInformation* network_information);
#endif //LAB1__SOCKET_INITIALIZATION_H_
