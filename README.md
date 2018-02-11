# imagecomp

**imagecomp** is an image optimization or compression library written in
Golang that uses `pngquant` and `mozjpeg` under the hood.

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
imagecomp
```

## Author

Armin Primadi https://github.com/aprimadi
