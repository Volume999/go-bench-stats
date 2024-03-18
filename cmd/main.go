package main

import (
	"go_bench_parser"
)

func main() {
	if err := start(); err != nil {
		panic(err)
	}
}

func start() error {
	return go_bench_stats.NewBenchStatsProcessor().Process()
}
