#include "lib/version.h"

#include <iostream>

int main(int argc, char** argv) {
   std::cout << "v1.2.3." << BuildVersion() << '\n';
   return 0;
}
