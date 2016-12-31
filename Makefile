# Makefile for EDL using Go
#

# get the tag, hash, and repo state of the project
XFLAG="-X main.Version=$(shell git describe --tags) -X main.Build=$(shell git describe --always --dirty)"

install:
# build the go code and apply a custom variable with the extracted version 
# data from GIT		
	go install -ldflags=$(XFLAG)