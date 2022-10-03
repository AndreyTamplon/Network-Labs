#include <stdio.h>
#include <sys/socket.h>
#include <netinet/ip.h>
#include <unistd.h>
#include <string.h>
#include <arpa/inet.h>
#include <stdlib.h>
#include "errno.h"
#include "time.h"
#include "storage.h"
#include "string.h"
#include <signal.h>
#include "socket_initialization.h"

int send_socket;
int listening_socket;
struct sockaddr_in group_address;
struct sockaddr_in listen_address;

void CopiesDetector() {
	/*
	 * It's important that hello and bye message have the same length, otherwise we'd have to clear buffer on each iteration
	 */
	char const* message = "hey";
	Array copies;
	if (init_array(&copies, 1) == false) {
		perror("init_array failed");
		close(send_socket);
		close(listening_socket);
		return;
	}
	while (1) {
		if (sendto(send_socket, message, strlen(message), 0,
				   (struct sockaddr*)&group_address, sizeof(group_address)) < 0) {
			perror("sendto failed");
			close(send_socket);
			close(listening_socket);
			free_array(&copies);
			return;
		}
		char buffer[1024];
		int len = sizeof(listen_address);
		ssize_t bytes_read = recvfrom(listening_socket, buffer, 1024, 0, (struct sockaddr*)&listen_address, &len);
		if (bytes_read < 0) {
			perror("receive from failed");
			close(send_socket);
			close(listening_socket);
			free_array(&copies);
			return;
		}
		Member member;
		member.address = (char*)malloc(sizeof(char) * 16);
		strcpy(member.address, inet_ntoa(listen_address.sin_addr));
		member.last_message_time = time(NULL);
		int index = find_element(&copies, &member);
		if (strcmp(buffer, "bye") == 0) {
			delete_element_by_id(&copies, index);
			printf("Member %s left the group\n", member.address);
		}
		else {
			if (index == -1) {
				append_array(&copies, &member);
				printf("New member joined %s\n", member.address);
			}
			else {
				copies.members[index].last_message_time = member.last_message_time;
			}
		}
		for (int i = 0; i < copies.used; i++) {
			if (copies.members[i].last_message_time - time(NULL) > 5) {
				delete_element(&copies, &copies.members[i]);
				printf("Member %s left the group\n", copies.members[i].address);
			}
		}
	}
}

void print_error() {
	const char* error_message = strerror(errno);
	write(STDERR_FILENO, error_message, strlen(error_message));
}

void signal_catcher(int sig) {
	if (sig == SIGINT) {
		char const* bye_message = "bye";
		if (sendto(send_socket, bye_message, strlen(bye_message), 0,
				   (struct sockaddr*)&group_address, sizeof(group_address)) < 0) {
			print_error();
			close(send_socket);
			close(listening_socket);
			exit(EXIT_FAILURE);
		}
		close(send_socket);
		close(listening_socket);
		exit(EXIT_SUCCESS);
	}
}

int main(int argc, char** argv) {
	if (argc < 3) {
		printf("Usage: %s <IP address> <network interface name>", argv[0]);
		return 1;
	}
	struct sigaction act;
	act.sa_handler = signal_catcher;
	sigaddset(&act.sa_mask, SIGINT);
	if (sigaction(SIGINT, &act, NULL) == -1) {
		perror("Failed to change signal SIGINT handler function");
		exit(EXIT_FAILURE);
	}
	NetworkInformation network_information = {argv[2], argv[1], 5555};
	init_listening_socket(&listening_socket, &listen_address, &network_information);
	init_send_socket(&send_socket, &group_address, &network_information);
	CopiesDetector();
	return 0;
}


