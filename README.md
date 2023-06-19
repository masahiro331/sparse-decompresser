
# spgzdc (Sparse GZip DeCompresser)

spgzdc is a high-performance tool written in Go for decompressing gzip files while maintaining their sparse characteristics. 
Often when large images or large sparse files are decompressed using standard gzip tools, 
the sparsity of the files is lost, consuming significant disk space. By contrast,
spgzdc ensures that the sparsity is preserved, thereby saving disk space.

## Features
* Decompress gzip files while preserving their sparse properties to save disk space.
* High-speed decompression leveraging Go's built-in concurrency model.
* Supports only .gz (gzip) files
 
## Usage

```
spgzdc [filename]
```
 
## Installation

You can install SPGZDC directly from our GitHub repository
or get github release.

```
go install github.com/yourusername/spgzdc@latest
```


## Prerequisites
Go 1.15 or higher

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.