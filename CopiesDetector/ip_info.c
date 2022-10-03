#include "ip_info.h"


int get_domain(const char* ip_address) {
	char buffer[20];
	if (inet_pton(AF_INET, ip_address, buffer) != 1) {
		if (inet_pton(AF_INET6, ip_address, buffer) != 1) {
			fprintf(stderr, "Neither IPv4 nor IPv6 address\n");
			return -1;
		}
		else {
			return AF_INET6;
		}
	}
	else {
		return AF_INET;
	}
}

char* get_host_ip_address(const char* net_interface) {
	struct ifreq ifr;
	int socket_fd = socket(AF_INET, SOCK_DGRAM, 0);
	if (socket_fd < 0) {
		perror("Failed to create socket");
		exit(EXIT_FAILURE);
	}
	ifr.ifr_addr.sa_family = AF_INET;
	strncpy(ifr.ifr_name, net_interface, IFNAMSIZ - 1);
	ioctl(socket_fd, SIOCGIFADDR, &ifr);
	close(socket_fd);
	return inet_ntoa(((struct sockaddr_in*)&ifr.ifr_addr)->sin_addr);
}