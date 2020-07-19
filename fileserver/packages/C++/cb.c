#include <stdio.h>

typedef void (*callback)(char* s);

void print(char* s)
{
	printf("%s\n",s);
}

void doo(char* s,void (*cb)(char* s))
{
	printf("%s\n",s);
	cb(s);
}

int main()
{
	doo("abc",print);
	return 0;
}

