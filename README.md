# zmqcat

CLI tool for easily connecting to existing ZMQ streams

## Installation

This package requires Go and a single dependency, `libzmq-dev`. Install it however is appropriate for your system. For example, on Debian-based systems:

```shell
sudo apt-get install libzmq3-dev
```

Then, install `zmqcat` using one of the following methods.

### From source

```shell
go install github.com/nsbruce/zmqcat@latest
```

**Note**: If installing from source, the `--version` flag will not be available. To get version information, you must install from pre-built binaries.

### From pre-built binaries

Pre-built binaries are available on the [releases page](https://github.com/nsbruce/zmqcat/releases).
