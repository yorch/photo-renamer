# photo-renamer

[![Build Status](https://travis-ci.org/yorch/photo-renamer.svg?branch=master)](https://travis-ci.org/yorch/photo-renamer)
[![](https://images.microbadger.com/badges/image/yorch/photo-renamer.svg)](https://microbadger.com/images/yorch/photo-renamer)
[![](https://images.microbadger.com/badges/version/yorch/photo-renamer.svg)](https://microbadger.com/images/yorch/photo-renamer)

## Overview

This tool allows you to bulk rename image files (including RAW ones) using their EXIF info, to a format containing the timestamp. For photographers aficionados, this makes it easier to sort and organize your photos, instead of keeping their original names.

So for instance, if you took a picture on Mar 31st, 2020 at 10:32:03 AM, your photo could have a name similar to `DSC_1234.jpg` (if you use Nikon cameras, Cannon use a different prefix I think but also keeping a sequence number). With this tool, your picture would be renamed to `20200331103203.jpg`.

Imagine you have 500 photos in your SD card, from different days, with a single command, all of those files would have a neat naming convention that you can use to organize them in folders, easily identify them or just sort them out.

## Docker

This tool can be build and run with Docker.

### Directly from the Docker Hub image

If you just need to run this tool on any computer with Docker installed (MacOS and Linux, since I haven't used Windows for anything serious for at least 8 years, I wouldn't know if it runs fine there too), you can run:

```sh
docker run -it -v `pwd`:/media yorch/photo-renamer /media
```

This command will rename every image in the current directory using its EXIF creation date, like: `2020033110303.jpg`, for easier sorting and photo organization.

> ProTip: You can create an `alias` with the previous command in your `.bashrc` or `.zshrc`, and you would be able to call it real fast every time you need to rename your images / photos.

### Locally

To build it:

```sh
docker build . -t photo-renamer
```

And to run it:

```sh
docker run -it -v `pwd`/temp:/photos photo-renamer /photos
```

## Build

If you have Go installed, you can build the binary by running:

```sh
cd src/
go build -o renamer main.go
```

Or run:

```sh
./build-go.sh
```

You can also run the app directly by:

```sh
cd src/
go run main.go
```

## Credits

Script originally taken from [here](https://gist.github.com/eko/6b0caaefeaf82f2aa202804743040292),
enhanced to support images with the exact same timestamp (like when continuously shooting) and other improvements.

## License

MIT License

Copyright (c) 2020 Jorge Barnaby

See [LICENSE](LICENSE)
