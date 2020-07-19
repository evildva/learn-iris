#include <stdio.h>
#include <math.h>
#include <fcntl.h>
#include <unistd.h>

int main(){
printf("%f\n",sin(90));
char c;
while(read(STDIN_FILENO,&c,1)!=0){
write(STDOUT_FILENO,&c,1);
}
return 0;
}
