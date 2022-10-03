cmake_minimum_required(VERSION 3.16)
project(lab1)

set(CMAKE_CXX_STANDARD 17)

add_executable(lab1 main.c storage.c storage.h ip_info.c ip_info.h socket_initialization.c socket_initialization.h)
