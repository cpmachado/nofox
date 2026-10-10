# nofox

[![build](https://github.com/cpmachado/nofox/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/cpmachado/nofox/actions/workflows/go.yml)

So, nofox is a bf interpreter for now. The lib however already supports
different mappings for different mappings.

## Status

- Minimal implementation.

## Usage:

```shell
./nofox -f samples/hello.bf
Hello World!
```

```shell
$ ./nofox -h
Usage of ./nofox:
  -f string
    	bf script to load (default "stdin")
  -l int
    	tapesize, must be a positive integer (default 30000)
  -v	display version
```

## Reference

- Esolang page: <https://esolangs.org/wiki/Brainfuck>
