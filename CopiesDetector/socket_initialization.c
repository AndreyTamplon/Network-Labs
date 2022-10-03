#include "socket_initialization.h"

void init_listening_socket(int* listening_socket,
						   struct sockaddr_in* listen_address,
						   const NetworkInformation* network_information) {
	*listening_socket = socket(AF_INET, SOCK_DGRAM, 0);
	if (*listening_socket < 0) {
		perror("receiving socket creation failed");
		exit(EXIT_FAILURE);
	}
	const int optval = 1;
	if (setsockopt(*listening_socket, SOL_SOCKET, SO_REUSEADDR, &optval, sizeof(int)) < 0) {
		perror("setsockopt SO_REUSEADDR failed");
		exit(EXIT_FAILURE);
	}
	memset((char*)listen_address, 0, sizeof(*listen_address));
	listen_address->sin_family = (sa_family_t)get_domain(network_information->group_address);
	listen_address->sin_port = htons((uint16_t)network_information->port);
	listen_address->sin_addr.s_addr = INADDR_ANY;
	if (bind(*listening_socket, (struct sockaddr*)listen_address, sizeof(*listen_address)) < 0) {
		perror("receiving_socket bind failed");
		close(*listening_socket);
		exit(EXIT_FAILURE);
	}
	char const* local_address = get_host_ip_address(network_information->interface_name);
	struct ip_mreq mreq;
	mreq.imr_multiaddr.s_addr = inet_addr(network_information->group_address);
	mreq.imr_interface.s_addr = inet_addr(local_address);
	if (setsockopt(*listening_socket, IPPROTO_IP, IP_ADD_MEMBERSHIP, &mreq, sizeof(mreq)) < 0) {
		perror("setsockopt IP_ADD_MEMBERSHIP failed");
		printf("%s", strerror(errno));
		close(*listening_socket);
		exit(EXIT_FAILURE);
	}
}

void init_send_socket(int* send_socket,
					  struct sockaddr_in* group_address,
					  const NetworkInformation* network_information) {
	int domain = get_domain(network_information->group_address);
	*send_socket = socket(get_domain(network_information->group_address), SOCK_DGRAM, 0);
	if (*send_socket < 0) {
		perror("sending socket creation failed");
		exit(EXIT_FAILURE);
	}
	memset((char*)group_address, 0, sizeof(*group_address));
	group_address->sin_family = (sa_family_t)domain;
	group_address->sin_port = htons((uint16_t)network_information->port);
	group_address->sin_addr.s_addr = inet_addr(network_information->group_address);
	char loopch = 0;
	if (setsockopt(*send_socket, IPPROTO_IP, IP_MULTICAST_LOOP,
				   &loopch, sizeof(loopch)) < 0) {
		perror("setting IP_MULTICAST_LOOP failed:");
		close(*send_socket);
		exit(EXIT_FAILURE);
	}
	char const* local_address = get_host_ip_address(network_information->interface_name);
	struct in_addr localInterface;
	localInterface.s_addr = inet_addr(local_address);
	if (setsockopt(*send_socket, IPPROTO_IP, IP_MULTICAST_IF,
				   (char*)&localInterface, sizeof(localInterface)) < 0) {
		perror("setting local interface failed:");
		close(*send_socket);
		exit(EXIT_FAILURE);
	}
}
