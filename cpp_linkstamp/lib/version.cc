#include "lib/version.h"

#include <string>

extern const char kBuildVersion[];

std::string BuildVersion() {
   return kBuildVersion;
}
