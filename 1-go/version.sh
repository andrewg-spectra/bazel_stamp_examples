#!/bin/bash

INITIAL_COMMIT_HASH=`git rev-list --max-parents=0 HEAD`
echo BUILD_VERSION=`git log -1 --format=%at $INITIAL_COMMIT_HASH`