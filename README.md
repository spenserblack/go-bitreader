# bitreader

[![Go Reference](https://pkg.go.dev/badge/github.com/spenserblack/go-bitreader.svg)](https://pkg.go.dev/github.com/spenserblack/go-bitreader)
[![CI](https://github.com/spenserblack/go-bitreader/actions/workflows/ci.yml/badge.svg)](https://github.com/spenserblack/go-bitreader/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/spenserblack/go-bitreader/branch/master/graph/badge.svg?token=156uzj6M1s)](https://codecov.io/gh/spenserblack/go-bitreader)
[![Go Report Card](https://goreportcard.com/badge/github.com/spenserblack/go-bitreader)](https://goreportcard.com/report/github.com/spenserblack/go-bitreader)

A `Reader` that reads *bits* instead of *bytes*.

⚠️ This module is deprecated in favor of [`go-bitio`](https://github.com/spenserblack/go-bitio). I foolishly made this module public
and ended up getting [its docs published on pkg.go.dev](https://pkg.go.dev/github.com/spenserblack/go-bitreader) before realizing
that it would be better to be an I/O module instead of just providing a reader.
