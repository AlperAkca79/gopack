# What is gozip
gozip is Command-Line based file compressor/decompressor for Linux/Windows/macOS.

# How to Use?
```shell
gozip -i inputFileName -f compressionFormat
```

### Arguments
- `-i`: Input file

- `-f`: Compression format (gzip, zip, rar etc...)

-  `-h` `--help` `-help`: Prints help message about **gozip**.

### Example
```shell
gozip -i README.md -f gzip
```