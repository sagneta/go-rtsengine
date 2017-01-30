#!/bin/sh
# Set your current working directory to the GOPATH
# thus allowing development here not to interfer with development elsewhere.
export GOPATH=`pwd`
export PATH=:$GOPATH/bin:$PATH

