
# spdc (Sparse DeCompresser)

spdc is a high-performance tool written in Go for decompressing gzip files while maintaining their sparse characteristics.
Often when large images or large sparse files are decompressed using standard gzip tools,
the sparsity of the files is lost, consuming significant disk space. By contrast,
spdc ensures that the sparsity is preserved, thereby saving disk space.

## Features
* Decompress gzip files while preserving their sparse properties to save disk space.
* High-speed decompression leveraging Go's built-in concurrency model.
* Supports only .gz (gzip) files

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