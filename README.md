scount: a commend-line streaming counter with rich progress report
==========

[![Build Status](https://github.com/oshikiri/scount/workflows/Go/badge.svg)](https://github.com/oshikiri/scount/actions) [![go report](https://goreportcard.com/badge/github.com/oshikiri/scount)](https://goreportcard.com/report/github.com/oshikiri/scount)


You can use it instead of `sort | uniq -c`.

```shell
cat ./demo/text8 | tr ' ' '\n' | scount -a | jq .the
```

![demo animation: approximate count using text8](demo/demo-text8-approximate-count.gif)


## Installation
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
  -e float
        Epsilon of lossy counting algorithm (default 1e-05)
  -f int
        Flush counting progress every X [msec] (default 200)
  -n int
        Print the top N items (default 10)
  -q    Turn on quiet mode
  -s float
        Support of lossy counting algorithm (default 2e-05)
```


## Build
```sh
go build
```


## Testing
```sh
go test
```

## TODO list
- UI
    - [x] ~~Format byte (1KB -> 1MB -> 1GB)~~
    - [x] ~~Add thousands separators~~
    - [ ] preserve previous command outputs (same behaviour as less command)
    - [ ] implement keybinds
        - [ ] <kbd>p</kbd> pause updating progresses
        - [ ] <kbd>q</kbd> quit
        - [ ] cursor
    - [x] ~~Add quiet option `-q`~~
    - [x] ~~Remove json_printer~~
- counting algorithm
    - [x] ~~Fix approximate counting algorithm~~
    - [ ] Add larger demo data
- enhancement
    - [ ] Add tuple counter (`-s` separator option)
    - barchart
        - [ ] character barchart
        - [ ] output progresses as JSON and plot barchart using d3.js
