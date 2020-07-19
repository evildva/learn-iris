#include <iostream>
#include <string>
#include <stdio.h>

extern "C"
{
#include "epollserver.h"
}

using namespace std;   

int main(int argc, char* argv[])  
{
	if(argc < 2)
    {
        printf("usage: ./server port\n");
        exit(1);
    }

	epoll_run(stoi(argv[1],0,10));

    return 0;  
}
