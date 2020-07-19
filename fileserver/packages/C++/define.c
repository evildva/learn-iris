#include <stdio.h>
#include "define.h"

int main()
{
	#ifdef A
		printf("A\n");
	#endif

	#ifdef B
		printf("B\n");
	#endif
	
	return 0;
}