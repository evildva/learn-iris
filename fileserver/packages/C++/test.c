/*
#include <stdio.h>

int main(int args,char* argv[]){

unsigned int a=10,b=20;
int c=5,d=8;

printf("a-b=%u\n",a-b);
printf("a-b=%d\n",a-b);
printf("d-a=%u\n",d-a);
printf("d-a=%d\n",d-a);

return 0;
}
*/
/*
#include <iostream>

using namespace std;
namespace sa{
void ab(){
cout<<"sa a"<<endl;
}

void ab(string s){
cout<<s<<endl;
}
}

using namespace sa;

void ab(){
cout<<"a"<<endl;
}

int main(int args,char* argv[]){
unsigned int a=10,b=20;
int c=-5,d=15;
cout<<"a-b="<<a-b<<endl;
cout<<"c-b="<<c-b<<endl;
cout<<"a-c="<<a-c<<endl;
cout<<"d-a="<<d-a<<endl;

::ab();

int i,&ri=i;

i=5;
ri=10;
cout<<ri<<endl;

auto m=3.33;
cout<<decltype(m)<<endl;
return 0;
}

string s;
cin>>s;
for(int i=s.size()-1;i>=0;i--){
cout<<s[i];
}
*/
/*
#include <stdio.h>

void * memset(void* buffer, int c, int count) 
{ 
     char * buffer_p=(char*)buffer; 
     //assert(buffer != NULL); 
     while(count-->0) 
         *buffer_p++=(char)c; 
     return buffer; 
}

int main(int args,char* argv[]){
char a[10]={0},b[5]={1},*c;
c=memset(a,3,5);
int i=0;
for(i=0;i<5;i++){
printf("%d\n",c[i]);
}
return 0;
}
*/

#include <iostream>

using namespace std;

int add(int b){
static int a=10;
a+=b;
return a;
}
int main(){
int i = 1;
int j = i++;
cout<<i<<endl;
if((i>j++) && (i++ == j)){
cout<<i<<endl;
i+=j;
}
cout<<i<<endl;
return 0;
}
