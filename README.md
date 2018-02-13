# imagecomp

**imagecomp** is an image optimization or compression library written in
Golang that uses `pngquant` and `mozjpeg` under the hood. It is useful for
optimizing images to conform to the Google PageSpeed Insights suggestion
(https://developers.google.com/speed/pagespeed/insights/)

## Dependencies

**imagecomp** requires `pngquant` and `mozjpeg` to be installed.

### On Mac OS X

To install the dependencies using brew do:

```
brew install pngquant
brew install mozjpeg
```

### On Ubuntu

To install `pngquant`:

```
sudo apt-get install pngquant
```

To install `mozjpeg`:

```
git clone https://github.com/mozilla/mozjpeg.git
cd mozjpeg
autoreconf -fiv
./configure
make
sudo make install
sudo ln -s /opt/mozjpeg/bin/cjpeg /usr/bin/cjpeg
sudo ln -s /opt/mozjpeg/bin/jpegtran /usr/bin/jpegtran
```

## Usage

To install, do:

```
go get github.com/aprimadi/imagecomp
go install github.com/aprimadi/imagecomp
```

To optimize images on a given directory do:

```
imagecomp .
```

Multiple directories can also be specified:

```
imagecomp app/assets public/assets
```

## Advanced Options

**imagecomp** accepts two options `-include` and `-exclude` which supports
wildcard character `*`. Use this two options to include and exclude path or
images being optimized. All paths are included by *default*.

Example:

```
imagecomp . -exclude "public/assets/*"
```

Not the use of quote `"`, this is mandatory otherwise the command line (i.e.
bash) will replace the arguments with real path.

Also, note that order matters. For example:

```
imagecomp . -exclude "*" -include "*.jpg"
```

Is not the same as:

```
imagecomp . -include "*.jpg" -exclude "*"
```

Which will excludes all images from being compressed, while the former only
process `.jpg` images.

## Author

Armin Primadi https://github.com/aprimadi

This library were developed by the [engineering team](https://www.dexcode.com/people) at [Dexcode](https://www.dexcode.com).
