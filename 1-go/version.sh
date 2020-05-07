#!/usr/bin/env bash

INITIAL_COMMIT_HASH=`git rev-list --max-parents=0 HEAD`
INITIAL_COMMIT_TIMESTAMP=`git log -1 --format=%at $INITIAL_COMMIT_HASH`
CURRENT_TIMESTAMP=`date +%s`
echo BUILD_VERSION $(($CURRENT_TIMESTAMP-$INITIAL_COMMIT_TIMESTAMP))
