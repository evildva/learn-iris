#include <iostream>
#include <fstream>
#include <csignal>
#include <unistd.h>
#include <pthread.h>

using namespace std;

namespace space
{
	template <typename T> T const& Max(T const& a,T const& b)
	{
		return a>b?a:b;
	}
}

int sig=1;

void set(int i)
{
	sig=0;
	exit(i);
}

void* dowhile(void*)
{
	
	signal(SIGHUP,set);

	while(sig)
	{
		cout<<"sleep . . ."<<endl;
		sleep(1);
	}

	return (void*)0;
}

void* terminate(void*)
{
	sleep(10);
	raise(SIGHUP);

	return (void*)0;
}

void nonecode(void)
{
	cout<<"abc"<<endl;
	cout<<"abc"<<"\b\b\b"<<"def"<<"\b\b\b"<<"ABC"<<endl;
	//cout<<args<<"\n"<<argv[0]<<endl;

	fstream file;
	char data[100];
	file.open("./web_test_file.txt",ios::out|ios::in|ios::trunc/*|ios::ate|ios::app*/);
	file<<"test content"<<endl;
	file.seekg(4,ios::beg/*|ios::cur|ios::end*/);
	file>>data;
	cout<<data<<endl;

	try
	{
		throw bad_exception(); 
	}
	catch(bad_exception &e)
	{
		cout<<e.what()<<endl;
	}

	pthread_t tid1,tid2;

	int ret=pthread_create(&tid1,NULL,dowhile,NULL);
	if (ret != 0)
        {
           cout << "pthread_create error: error_code=" << ret << endl;
        }
    int ren=pthread_create(&tid2,NULL,terminate,NULL);

	if (ren != 0)
        {
           cout << "pthread_create error: error_code=" << ren << endl;
        }

    pthread_exit(NULL);
}

int main(int args,char** argv)
{
	
	class car
	{
	private:
		int num;
	public:
		void drive();
	};

	void car::drive()
	{

	}

	return 0;
}
