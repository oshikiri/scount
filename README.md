scount: a commend-line streaming counter with rich progress report
==========

You can use it instead of `sort | uniq -c`.

[![Build Status](https://github.com/oshikiri/scount/workflows/Go/badge.svg)](https://github.com/oshikiri/scount/actions) [![go report](https://goreportcard.com/badge/github.com/oshikiri/scount)](https://goreportcard.com/report/github.com/oshikiri/scount)

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
  -f int
        Flush counting progress every X [msec] (default 100)
  -n int
        Print the top N items (default 10)
  -q    Quiet mode
  -t float
        theta of KSP algorithm (default 1e-05)
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
    - [ ] Add thousands separators
    - [ ] preserve previous command outputs (same behaviour as less command)
    - [ ] implement keybinds
        - [ ] <kbd>p</kbd> pause updating progresses
        - [ ] <kbd>q</kbd> quit
        - [ ] cursor
    - [x] ~~Add quiet option `-q`~~
    - [x] ~~Remove json_printer~~
- counting algorithm
    - [ ] Fix approximate counting algorithm
    - [ ] Add larger demo data
- enhancement
    - [ ] Add tuple counter (`-s` separator option)
    - barchart
        - [ ] character barchart
        - [ ] output progresses as JSON and plot barchart using d3.js
