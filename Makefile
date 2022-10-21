NAME := gzip
PACKAGE_NAME := github.com/DoubleCircle-Salt/gzip

PLATFORM := linux
BUILD_DIR := build
GOBUILD = go build -tags "full" -trimpath -ldflags="-s -w" -o gzip.so -buildmode=c-shared gzip.go

normal: clean gzip

clean:
	rm -rf gzip.so
	rm -rf gzip.h

gzip:
	$(GOBUILD)

