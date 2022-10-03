#ifndef LAB1__IP_INFO_H_
#define LAB1__IP_INFO_H_
#include <sys/socket.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include <sys/ioctl.h>
#include <net/if.h>
#include <stdio.h>
int get_domain(const char* ip_address);
char* get_host_ip_address(const char* interface_name);
#endif //LAB1__IP_INFO_H_
