# photo-renamer

[![Build Status](https://travis-ci.org/yorch/photo-renamer.svg?branch=master)](https://travis-ci.org/yorch/photo-renamer)
[![](https://images.microbadger.com/badges/image/yorch/photo-renamer.svg)](https://microbadger.com/images/yorch/photo-renamer)
[![](https://images.microbadger.com/badges/version/yorch/photo-renamer.svg)](https://microbadger.com/images/yorch/photo-renamer)

Script originally taken from [here](https://gist.github.com/eko/6b0caaefeaf82f2aa202804743040292) and
enhanced to support images with the exact same datestamp (like when continuously shooting).

## Build

To build the binary, we can run:

```sh
cd src/
go build -o renamer main.go
```

Or run:

```sh
./build-go.sh
```

## Docker

This tool can be build and run with Docker.

To build it:

```sh
docker build . -t photo-renamer
```

And to run it:

```sh
docker run -it -v `pwd`/temp:/photos photo-renamer /photos
```
