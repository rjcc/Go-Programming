package main

import "fmt"

func PatternCount(pattern, text string) int {
	cnt := 0
	for i := range text[:len(text)-len(pattern)+1] {
		if pattern == text[i:i+len(pattern)] {
			cnt++
		}
	}
	return cnt
}

func MaxDict(dict map[string]int) int {
	init := false
	max := 0
	for _, num := range dict {
		if !init {
			init = true
			max = num
		} else {
			if num > max {
				max = num
			}
		}

	}
	return max
}

func FrequencyMap(text string, k int) map[string]int {
	res := make(map[string]int)
	for i := range text[:len(text)-k+1] {
		res[text[i:i+k]]++
	}
	return res
}

func FrequentWords(text string, k int) []string {
	a := FrequencyMap(text, k)
	res := []string{}
	max := 0
	for s, num := range a {
		if num > max {
			max = num
			res = []string{0: s}
		} else if num == max {
			res = append(res, s)
		}
		fmt.Print(num, res)
	}
	return res
}

func ReverseComplement(pattern string) string {
	complement := map[rune]rune{'A': 'T', 'C': 'G', 'G': 'C', 'T': 'A'}
	res := make([]rune, len(pattern))
	for i, v := range pattern {
		res[len(pattern)-1-i] = complement[v]
	}
	return string(res)
}

func PatternMatching(pattern, text string) []int {
	res := []int{}
	for i := range text[:len(text)-len(pattern)+1] {
		if pattern == text[i:i+len(pattern)] {
			res = append(res, i)
		}
	}
	return res
}

func ClumpFinding(genome string, k, L, t int) []string {
	left := 0
	set := make(map[string]bool)
	for left = range genome[:len(genome)-L+1] {
		cntmap := FrequencyMap(genome[left:left+L], k)
		for key, value := range cntmap {
			if value >= t {
				set[key] = true
			}
		}
	}
	res := make([]string, 0, len(set))
	for k := range set {
		res = append(res, k)
	}
	return res
}

func SkewArray(genome string) []int {
	skew := map[rune]int{'A': 0, 'C': -1, 'G': 1, 'T': 0}
	res := []int{0: 0}
	current := 0
	for _, c := range genome {
		current = current + skew[c]
		res = append(res, current)
	}
	return res
}

func MinimumSkew(genome string) []int {
	skew := map[rune]int{'A': 0, 'C': -1, 'G': 1, 'T': 0}
	current := 0
	min := 0
	res := []int{0: 1}
	for i, c := range genome {
		current = current + skew[c]
		if current < min {
			min = current
			res = []int{0: i + 1}
		} else if current == min {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	FrequentWords("ACGTTGCATGTCGCATGATGCATGAGAGCT", 4)
}
