#include "lib/version.h"
#include "lib/version.gen.h"

#include <string>

extern const char kBuildVersion[];

std::string BuildVersion() {
   return kBuildVersion;
}
