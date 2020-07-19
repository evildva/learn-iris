#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>
#include <stdio.h>

int main()
{
	int sock=socket(AF_INET,SOCK_STREAM,0);
	if(sock==-1)
	{
		printf("%s\n", "socket fail\n");
		return -1;
	}

	struct sockaddr_in addr={
		0
	};
	addr.sin_family=AF_INET;
	addr.sin_port=htons(8000);
	//addr.sin_addr.s_addr=inet_addr("192.168.43.2");

	if(bind(sock,(struct sockaddr*)&addr,sizeof(struct sockaddr))==-1)
	{
		printf("%s\n", "bind fail\n");
		return -1;
	}

	if(listen(sock,10)==-1)
	{
		printf("%s\n", "listen fail\n");
		return -1;
	}

	struct sockaddr_in client={
		0
	};
	int addrlen=sizeof(struct sockaddr);
	int cs=accept(sock,(struct sockaddr*)&client,&addrlen);

	if(cs==-1)
	{
		printf("%s\n", "accept fail\n");
		return -1;
	}
	else
	{
		char buff[2048];
		if(recv(cs,buff,2047,0)==-1)
			printf("recv error\n");
		else
			
		printf("request header: \n%s",buff);

		char buf[1024],content[]="<html>\n\r<head><title>Server</title></head>\n\r<body>\n\rDave O'Hallaron\n\r</body>\n\r</html>\n\r";
		sprintf(buf, "HTTP/1.0 200 OK\n");
	    sprintf(buf, "%sServer: Tiny Web Server\r\n", buf);
	    sprintf(buf, "%sConnection: close\r\n", buf);
	    sprintf(buf, "%sContent-length: %d\r\n", buf, strlen(content));
	    sprintf(buf, "%sContent-type: %s\r\n\r\n", buf, "text/html");
	    write(cs, buf, strlen(buf));
	    write(cs, content, strlen(content));
	    printf("Response headers:\n");
	    printf("%s", buf);

	    printf("%s", content);

	    printf("%s  %d\n", inet_ntoa(client.sin_addr),ntohs(client.sin_port));
	}
}
