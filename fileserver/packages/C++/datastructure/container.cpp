#include <iostream>
#include <time.h>
#include <stdio.h>
#include <stdlib.h>
#include <string>
#include <array>
#include <vector>
#include <set>
#include <map>
#include <algorithm>

using namespace std;

template<class T>
void test(T t)
{
    std::cout << std::tuple_size<T>::value << '\n';
    typename std::tuple_element<0,T>::type s=5;
    std::cout << s <<endl;
}

template<class BidirIt, class OutputIt>
OutputIt reverse_copy(BidirIt first, BidirIt last, OutputIt d_first)
{
    while (first != last) {
        *(d_first++) = *(--last);
    }
    return d_first;
}

struct TreeNode {
	int val;
 	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<int> Preorder ;
    map<int,int> dic;
    
    TreeNode* build(int pre_root ,int in_left ,int in_right){
        //如果左边界大于右边界说明到过了叶子
        if(in_left > in_right){
            return NULL;
        }
        //pre_root 是先序里面的索引 ！！
        TreeNode* root = new TreeNode(Preorder[pre_root]);
        //获取先序中的节点在中序中的节点， 即index 左边就是这节点的左子树，index右边就是节点的右子树
        int index = dic[Preorder[pre_root]];
        //当前节点左树即为先序索引+1 （没了话会在下一次迭代返回NULL）
        root->left = build(pre_root+1,in_left,index-1);
        //当前节点右树即为 根结点在前序中的索引+左树所有节点数（即节点在中序中的索引）-左边界+1 ，下一次的左边界为根在中序的索引+1  
        root->right = build(pre_root+index-in_left+1,index+1 ,in_right);
        return root;
    }

    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        //赋值至外部变量
        Preorder = preorder;
        //使用map映射inorder的值和索引，提高找到索引效率
        for(int i=0;i<inorder.size();i++){
            dic[inorder[i]] = i;
        }
        return build(0,0,preorder.size()-1);
    }
};

int main(int argc, char const *argv[])
{
	
	struct timespec time_start={0, 0},time_end={0, 0};
    clock_gettime(CLOCK_REALTIME, &time_start);
    printf("start time %llus,%llu ns\n", time_start.tv_sec, time_start.tv_nsec);

	array<int,4> b={1,2,3,4};
	test(b);
	cout<< b[1]<<endl;
	cout<< b.at(2)<<endl;
	cout<< b.size()<<endl;
	cout<< b.max_size()<<endl;

	std::vector<int> preorder = {3,9,8,5,4,10,20,15,7};
    std::vector<int> inorder = {4,5,8,10,9,3,15,20,7};

    //Solution s;
    //s.buildTree(preorder,inorder);

    set<int> c={1,3,4,2,5,7,9};
    for(auto a :c){
    	cout<<a<<" ";
    }
    cout<<endl;
    for_each(c.begin(),c.end(),[](const auto &n){ cout<<n<<" "; });
    cout<<endl;

    map<string,string> m={{"a","a"},{"b","b"},{"c","c"},{"d","d"}};
    for_each(m.begin(),m.end(),[](const auto &n){ cout<<n.first<<" "<<n.second<<" "; });
    cout<<endl;

    int *array=(int *)malloc(5 * sizeof(int)),*p;
    p=array;
    for(int i=0;i<5;i++){
    	//array[i]=i;//ok
    	*(array+i)=i;
    	cout<<"array "<<&array[i]<<" i "<<i<<" "<<array++<<endl;
    }
    for(int i=0;i<5;i++){
    	cout<<array[i]<<" ";
    }
    cout<<endl;
    free(p);

    int *arr=new int[5];
    for(int i=0;i<5;i++){
    	arr[i]=i;
    	cout<<"array "<<&array[i]<<" i "<<i<<" "<<array++<<endl;
    }
    for(int i=0;i<5;i++){
    	cout<<arr[i]<<" ";
    }
    delete arr;

    cout<<sizeof(unsigned int)<<endl;
	
    clock_gettime(CLOCK_REALTIME, &time_end);
    printf("endtime %llus,%llu ns\n", time_end.tv_sec, time_end.tv_nsec);
    printf("duration:%llus %lluns\n", time_end.tv_sec-time_start.tv_sec, time_end.tv_nsec-time_start.tv_nsec);
	return 0;
}