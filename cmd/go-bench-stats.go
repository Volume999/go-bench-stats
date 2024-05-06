package main

import (
	"go_bench_parser"
	"os"
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	return go_bench_stats.NewBenchStatsProcessor().Process(os.Stdin, os.Stdout, os.Stderr)
}
