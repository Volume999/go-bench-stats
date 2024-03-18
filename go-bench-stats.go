package go_bench_stats

import "regexp"

var benchmarkPattern = regexp.MustCompile(
	`^Benchmark(?P<benchmarkName>[^-\s]+)(-(?P<cpus>\d+))?\s+(?P<iterations>\d+)(?P<metrics>.+)$`)

var benchmarkTypePattern = regexp.MustCompile(
	`(\/(?P<name>[^=]+)=(?P<val>[^\/]+))+?`)

var benchmarkMetricsPattern = regexp.MustCompile(
	`(\s+(?P<val>[\d.]+)\s(?P<unit>\S+))+?`)

type BenchStatsProcessor struct{}

func NewBenchStatsProcessor() *BenchStatsProcessor {
	return &BenchStatsProcessor{}
}

func (b *BenchStatsProcessor) Process() error {
	return nil
}
