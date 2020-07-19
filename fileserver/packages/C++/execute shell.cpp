#include <iostream>
#include <unistd.h>
#include <sys/types.h>
#include <cstdlib>

using namespace std;

int main()
{
	FILE *pp = popen("cd . && ls -l", "r"); // build pipe
	if (!pp)
		return 1;

	// collect cmd execute result
	char tmp[1024];
	while (fgets(tmp, sizeof(tmp), pp) != NULL)
		std::cout << tmp << std::endl; // can join each line as string
	pclose(pp);

	system("ps -ef| grep gnome");

	return 0;
}

