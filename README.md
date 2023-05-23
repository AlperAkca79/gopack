# What is gozip
gozip is Command-Line based file compressor/decompressor for Linux/Windows/macOS. 

# How to Use?
**First, you need to put *gozip* in the system path.**

### Usage

```
gozip -i inputFileName -f compressionFormat
```

### Arguments
- `-i`: Input file

- `-f`: Compression format (gzip, zip, rar etc...) **(defaults to gzip)**

- `-v`: Prints version of gozip in your system.

-  `-h` `--help` `-help`: Prints help message about **gozip**.

### Example
```
gozip -i README.md -f gzip
```

# How to Download?
### Via [Go (aka Golang)](https://github.com/golang/go)
```
go install github.com/AlperAkca79/gozip
```

### Via [Git](https://git-scm.com/)
```shell
git clone https://github.com/AlperAkca79/gozip.git
```

### Via [GitHub CLI](https://github.com/cli/cli)
```shell
gh repo clone AlperAkca79/gozip
```