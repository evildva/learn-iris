#include <iostream>
#include <vector>
#include <fstream>
#include <sstream>
#include <string>
#include <regex>
#include <filesystem>

using namespace std;
using namespace std::filesystem;

int main(int nums,char** args)
{
	/*
	vector<int> v;
	for(int i=0;i<=9;++i)
	{
		v.push_back(i);
	}

	vector<int>::iterator it;
	for(it=v.begin();it!=v.end();++it)
	{
		cout<<*it<<endl;
	}

	cout<<*(v.end())<<endl;
	cout<<*(v.rend())<<endl;
	*/
	path pat;
	fstream fs;
	string line;
	string searchpath="";

	cin>>searchpath;

	if(searchpath!="")
	{
		path pa(searchpath);
		if(exists(pa))
		{
			pat=pa;
			cout<<string(pa)<<endl;
		}
		else
		{
			cout<<"路径不存在，需要绝对路径或　．路径"<<endl;
			return -1;
		}
	}

	vector<directory_entry> spath;

	for(auto& file:recursive_directory_iterator(pat))
	{
		if(!file.is_directory())
			spath.push_back(file);
		else
			cout<<"dir: "+string(file.path())<<endl;
	}

	for(auto f:spath)
	{
		/*
		cout<<"file: "+string(f.path())<<endl;
		cout<<f.path().extension()<<endl;
		*/
		path pa=f.path();
		path p=pa.extension();
		if(p==".c"||p==".cpp"||p==".h")
		{
			//cout<<"file: "+string(f.path())<<endl;
			fs.open(pa,ios::in);
			fs>>line;
			string s="#include";
			if(find(line,s))
				cout<<pa<<" "<<line<<endl;
		}
	}

	return 0;
}