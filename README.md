scount: a command-line streaming counter with rich progress report
==========

[![Build Status](https://github.com/oshikiri/scount/workflows/Go/badge.svg)](https://github.com/oshikiri/scount/actions) [![go report](https://goreportcard.com/badge/github.com/oshikiri/scount)](https://goreportcard.com/report/github.com/oshikiri/scount) [![release](https://img.shields.io/github/v/release/oshikiri/scount.svg)](https://github.com/oshikiri/scount/releases)


You can use it instead of `sort | uniq -c`.

```shell
cat ./demo/text8 | tr ' ' '\n' | scount -a | \
    jq '. | sort_by(.count) | reverse | .[0:3]'
```

![demo animation: approximate count using text8](demo/demo-text8-approximate-count.gif)


## Installation
### Using binary
See <https://github.com/oshikiri/scount/releases>.

### From source code
```sh
go get github.com/oshikiri/scount
go install github.com/oshikiri/scount
```

and then add `~/go/bin` to `$PATH`.


## Usage
```
$ scount -h
Usage of scount:
  -a    Use approximate counting algorithm (default is naive counting)
  -ae float
        Epsilon of lossy counting algorithm (default 1e-05)
  -as float
        Support of lossy counting algorithm (default 2e-05)
  -f int
        Flush counting progress every X [msec] (default 200)
  -n int
        Print the top N items (default 10)
  -q    Suppress a progress report
```


## Development
```sh
# Create ./scount binary
go build

# Run tests
go test
```
