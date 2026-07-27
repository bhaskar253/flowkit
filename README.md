# FlowKit

[![FlowKit CI](https://github.com/bhaskar253/flowkit/actions/workflows/flowkit-ci.yml/badge.svg)](https://github.com/bhaskar253/flowkit/actions/workflows/flowkit-ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/bhaskar253/flowkit.svg)](https://pkg.go.dev/github.com/bhaskar253/flowkit)
[![Go Version](https://img.shields.io/github/go-mod/go-version/bhaskar253/flowkit)](https://go.dev/)
[![License](https://img.shields.io/github/license/bhaskar253/flowkit)](LICENSE)
[![GitHub Stars](https://img.shields.io/github/stars/bhaskar253/flowkit)](https://github.com/bhaskar253/flowkit/stargazers)

A lightweight generic workflow and pipeline toolkit for Go.

## Features

- Generic pipelines
- Producer / Transformer / Consumer model
- Stream processing
- Functional operators
- Iterator support
- Zero external dependencies
- Built with Go generics


## Installation

    go get github.com/bhaskar253/flowkit


## Quick Example

    package main

    import (
        "fmt"

        "github.com/bhaskar253/flowkit"
    )

    func main() {
        count := 1
        pipeline := flowkit.New(

            func() (int, bool) {
                if count > 5 {
                    return 0, false
                }
                value := count
                count++
                return value, true
            },

            func(value int) int {
                return value * value
            },

            func(value int) {
                fmt.Println(value)
            },
        )
        pipeline.Run()
    }


Output:

    1
    4
    9
    16
    25


## Development

Format:

    go fmt ./...


Test:

    go test ./...


Validate:

    go vet ./...


## License

Apache License 2.0
