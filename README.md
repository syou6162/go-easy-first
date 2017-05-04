# go-easy-first
[![CircleCI](https://circleci.com/gh/syou6162/go-easy-first.svg?style=shield)](https://circleci.com/gh/syou6162/go-easy-first)
[![Go Report Card](https://goreportcard.com/badge/github.com/syou6162/go-easy-first)](https://goreportcard.com/report/github.com/syou6162/go-easy-first)
[![Coverage Status](https://coveralls.io/repos/github/syou6162/go-easy-first/badge.svg?branch=coveralls)](https://coveralls.io/github/syou6162/go-easy-first?branch=coveralls)

go-easy-first - Dependency Parser with Easy-First Algorithm (An Efficient Algorithm for Easy-First Non-Directional Dependency Parsing, NAACL-2010, Yoav Goldberg and Michael Elhadad) written in Go.

# Build from source

```sh
% git clone https://github.com/syou6162/go-easy-first.git
% cd go-easy-first
% go build
```

# Usage
go-easy-first has `train` (training a parser phase) and `eval` (evaluating a trained parser phase) modes. To see the detail options, type `./go-easy-first --help`.

## Training a parser
To see the detail options, type `./go-easy-first train --help`.

```sh
% ./go-easy-first train --train-filename path/to/train.txt --dev-filename path/to/dev.txt --max-iter 10 --model-filename model.bin
0, 0.907, 0.893
1, 0.920, 0.901
2, 0.929, 0.904
3, 0.935, 0.906
4, 0.940, 0.907
5, 0.944, 0.907
6, 0.947, 0.908
7, 0.950, 0.908
8, 0.953, 0.908
9, 0.955, 0.908
```

## Evaluating a trained parser
To see the detail options, type `./go-easy-first eval --help`.

```
% ./go-easy-first eval --test-filename path/to/test.txt --model-filename model.bin
| SENTENCES | SECONDS | ACCURACY |
|-----------|---------|----------|
|      1346 |    4.60 |    0.888 |
```

# Raadmap
- [ ] Implement PP-Attachment features
- [ ] Beam search with max-violation perceptron
- [ ] Mini-batch update
- [ ] Embed weight parameters to a built binary file using go-bindata

# Author
Yasuhisa Yoshida
