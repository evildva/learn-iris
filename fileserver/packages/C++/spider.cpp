#include <iostream>
#include <stdio.h>
#include <curl/curl.h>
#include "gumbo.h"

using namespace std;
 
static size_t downloadCallback(void *buffer, size_t sz, size_t nmemb, void *writer)
{
	string* psResponse = (string*) writer;
	size_t size = sz * nmemb;
	psResponse->append((char*) buffer, size);
 
	return sz * nmemb;
}

static void search_for_links(GumboNode* node) {
  if (node->type != GUMBO_NODE_ELEMENT) {
    return;
  }
  GumboAttribute* href;
  if (node->v.element.tag == GUMBO_TAG_A &&
      (href = gumbo_get_attribute(&node->v.element.attributes, "href"))) {
    std::cout << href->value << std::endl;
  }

  GumboVector* children = &node->v.element.children;
  for (unsigned int i = 0; i < children->length; ++i) {
    search_for_links(static_cast<GumboNode*>(children->data[i]));
  }
}
 //gcc http.cpp -o http -lcurl -lstdc++
int main()
{
	string strUrl = "";
	string strTmpStr;
	cin>>strUrl;
	CURL *curl = curl_easy_init();
	curl_easy_setopt(curl, CURLOPT_URL, strUrl.c_str());
	curl_easy_setopt(curl, CURLOPT_USERAGENT, "'User-Agent':'Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36'");
	curl_easy_setopt(curl, CURLOPT_HEADER, true);
	curl_easy_setopt(curl, CURLOPT_HEADEROPT, true);
	curl_easy_setopt(curl, CURLOPT_NOSIGNAL, 1L);
	curl_easy_setopt(curl, CURLOPT_TIMEOUT, 2);
	curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, downloadCallback); 
	curl_easy_setopt(curl, CURLOPT_WRITEDATA, &strTmpStr);
	CURLcode res = curl_easy_perform(curl);
	curl_easy_cleanup(curl);
	string strRsp;
	if (res != CURLE_OK)
	{
		strRsp = "error";
	}
	else
	{
		strRsp = strTmpStr;
	}

	cout<<strRsp.c_str()<<endl;

    GumboOutput* output = gumbo_parse(strRsp.c_str());
    search_for_links(output->root);
    gumbo_destroy_output(&kGumboDefaultOptions, output);
	
	return 0;
}


