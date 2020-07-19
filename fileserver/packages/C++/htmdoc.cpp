#include <stdlib.h>

#include <fstream>
#include <iostream>
#include <string>

#include "gumbo.h"

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

void replace_contents(GumboNode* node)
{
  if (node->type != GUMBO_NODE_ELEMENT) {
    return;
  }
  GumboAttribute* div;
  if (node->v.element.tag == GUMBO_TAG_DIV &&
      (div = gumbo_get_attribute(&node->v.element.attributes, "id"))) {
    std::cout << div->value << std::endl;
  }

}

int main(int argc, char** argv) {
  if (argc != 2) {
    std::cout << "Usage: THISPROGRAM <html filename>.\n";
    exit(EXIT_FAILURE);
  }
  const char* filename = argv[1];

  std::ifstream in(filename, std::ios::in | std::ios::binary);
  if (!in) {
    std::cout << "File " << filename << " not found!\n";
    exit(EXIT_FAILURE);
  }

  std::string contents;
  in.seekg(0, std::ios::end);
  contents.resize(in.tellg());
  in.seekg(0, std::ios::beg);
  in.read(&contents[0], contents.size());
  in.close();

  GumboOutput* output = gumbo_parse(contents.c_str());
  replace_contents(output->root);
  gumbo_destroy_output(&kGumboDefaultOptions, output);
}
