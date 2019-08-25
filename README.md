scount: a CLI streaming counter written in Go
==========

[![Build Status](https://travis-ci.org/oshikiri/scount.svg?branch=master)](https://travis-ci.org/oshikiri/scount)

![demo animation: approximate count using text8](demo/demo-text8-approximate-count.gif)


## Installation
```sh
go install github.com/oshikiri/scount
```

and then add `~/go/bin` to `$PATH`.


## Usage
```
$ scount -h
Usage of scount:
  -a    Use approximate counting algorithm (default is naive counting)
  -f int
        Flush counting progress every X [msec] (default 100)
  -n int
        Print the top N items (default 10)
  -t float
        theta of KSP algorithm (default 1e-05)
```


## Testing
```sh
# unit testing
go test

# testing command line interface
bash ./test/test_scount.sh
```

## TODO
- [ ] Setup UI: implement keybinds
- [ ] Add tuple counter (`-s` separator option)
