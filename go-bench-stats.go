package go_bench_stats

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var benchmarkPattern = regexp.MustCompile(
	`^Benchmark(?P<benchmarkName>[^/\s]+)(?P<benchmarkDimensions>[^-\s]+)(?:-(?P<cpus>\d+))?\s+(?P<iterations>\d+)(?P<metrics>.+)$`)

var benchmarkDimensionsPattern = regexp.MustCompile(
	`(\/(?P<name>[^=]+)=(?P<val>[^\/]+))+?`)

var benchmarkMetricsPattern = regexp.MustCompile(
	`(\s+(?P<val>[\d.]+)\s(?P<unit>\S+))+?`)

var headerCreated = false

type BenchStatsProcessor struct{}

func NewBenchStatsProcessor() *BenchStatsProcessor {
	return &BenchStatsProcessor{}
}

func (b *BenchStatsProcessor) Process(in io.Reader, out, errOut io.Writer) error {
	s := bufio.NewScanner(in)
	for s.Scan() {
		line := s.Text()
		if !benchmarkPattern.MatchString(line) {
			continue
		}
		match := findNamedMatches(benchmarkPattern, line)
		//dimensions := findNamedMatches(benchmarkDimensionsPattern, match["benchmarkName"])
		//dimensions := benchmarkDimensionsPattern.FindAllStringSubmatch(match["benchmarkName"], -1)
		dimensions := findNamedMatchesAll(benchmarkDimensionsPattern, match["benchmarkDimensions"])
		metrics := findNamedMatchesAll(benchmarkMetricsPattern, match["metrics"])

		if !headerCreated {
			createHeader(dimensions, metrics, out)
			headerCreated = true
		}
		data := []string{match["benchmarkName"]}
		for _, dimension := range dimensions {
			data = append(data, dimension["val"])
		}
		data = append(data, match["cpus"])
		data = append(data, match["iterations"])
		for _, metric := range metrics {
			data = append(data, metric["val"])
		}
		fmt.Fprintln(out, strings.Join(data, ","))
	}
	return nil
}

func createHeader(dimensions []map[string]string, metrics []map[string]string, out io.Writer) {
	header := []string{"name"}

	for _, dimension := range dimensions {
		header = append(header, dimension["name"])
	}

	header = append(header, "cpus")
	header = append(header, "iterations")

	for i, metric := range metrics {
		header = append(header, fmt.Sprintf("metric%d (%s)", i, metric["unit"]))
	}

	fmt.Fprintln(out, strings.Join(header, ","))
}

func findNamedMatches(pattern *regexp.Regexp, line string) map[string]string {
	matches := pattern.FindStringSubmatch(line)
	res := map[string]string{}
	for i, name := range matches {
		res[pattern.SubexpNames()[i]] = name
	}
	return res
}

func findNamedMatchesAll(pattern *regexp.Regexp, line string) []map[string]string {
	matches := pattern.FindAllStringSubmatch(line, -1)
	res := make([]map[string]string, len(matches))
	for i, match := range matches {
		res[i] = map[string]string{}
		for j, name := range match {
			res[i][pattern.SubexpNames()[j]] = name
		}
	}
	return res
}
