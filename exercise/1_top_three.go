package exercise

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type TopThreeWords struct {
	First  string
	Second string
	Third  string
}

func ComputeWordCount(f io.Reader) (map[string]int, error) {
	scanner := bufio.NewScanner(f)
	wc := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		for _, token := range tokens {
			wc[token] = wc[token] + 1
		}
	}
	return wc, nil
}

type WordAndCount struct {
	Word  string
	Count int
}

func flatten(wc map[string]int) []WordAndCount {
	var wac []WordAndCount
	for k, v := range wc {
		wac = append(wac, WordAndCount{k, v})
	}
	return wac
}

func computeTopThree(wac []WordAndCount) (TopThreeWords, error) {
	if len(wac) < 3 {
		return TopThreeWords{}, fmt.Errorf("not enough words")

	}
	slices.SortFunc(wac, func(a, b WordAndCount) int {
		return b.Count - a.Count
	})
	return TopThreeWords{
		First:  wac[0].Word,
		Second: wac[1].Word,
		Third:  wac[2].Word,
	}, nil

}

func CalculateTopThree(fname string) (TopThreeWords, error) {

	f, err := os.Open(fname)
	if err != nil {
		return TopThreeWords{}, err
	}
	defer f.Close()

	wc, err := ComputeWordCount(f)
	if err != nil {
		return TopThreeWords{}, err
	}
	wac := flatten(wc)
	return computeTopThree(wac)
}
