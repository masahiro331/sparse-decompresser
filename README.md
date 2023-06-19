
# spdc (Sparse DeCompresser)

spdc is a high-performance tool written in Go for decompressing gzip files while maintaining their sparse characteristics.
Often when large images or large sparse files are decompressed using standard gzip tools,
the sparsity of the files is lost, consuming significant disk space. By contrast,
spdc ensures that the sparsity is preserved, thereby saving disk space.

## Features
* Decompress gzip files while preserving their sparse properties to save disk space.
* High-speed decompression leveraging Go's built-in concurrency model.
* Supports only .gz (gzip) files

## why output sparse file?

In the case of a large image file, the file size is large, but the actual data is small.
If you decompress the image file with gzip, the file size will be the same as the original file size.

```
$ ls -lah ubuntu-2204.img.gz
-rw-r--r--  1 masahiro331  staff   760M  6 19 23:53 ubuntu-2204.img.gz

$ file ubuntu-2204.img.gz
ubuntu-2204.img.gz: gzip compressed data, was "ubuntu-2204.img", last modified: Sun Jun 18 12:17:19 2023, from Unix, original size modulo 2^32 0
```

Use spdc to decompress the image file.
Actual data size is 1.8G.

```
$ time ./spdc ubuntu-2204.img.gz
real    1m8.568s
user    0m56.180s
sys     0m50.441s

$ ls -lah ubuntu-2204.img
-rw-r--r--  1 masahiro331  staff    40G  6 19 23:56 ubuntu-2204.img

$ du -h ubuntu-2204.img
1.8G    ubuntu-2204.img
```

Use gzip to decompress the image file.
Actual data size is 40G.

```
$ time gzip -d ubuntu-2204.img.gz
real    1m39.572s
user    1m5.913s
sys     0m7.954s

$ ls -lah ubuntu-2204.img
-rw-r--r--  1 masahiro331  staff    40G  6 19 23:53 ubuntu-2204.img

$ du -h ubuntu-2204.img
 40G    ubuntu-2204.img
```

## Usage

```
spdc [filename]
```

## Installation

You can install spdc directly from our GitHub repository
or get github release.

```
go install github.com/yourusername/spdc@latest
```


## Prerequisites
Go 1.19 or higher

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.