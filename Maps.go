package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	wordsDist := make(map[string]int)
	
	for _, word := range words {
	//   if _, ok := wordsDist[word]; !ok{
	//     wordsDist[word] = 1
	//   } else {
	//     wordsDist[word] += 1
	//   }

	wordDist[word]++ // ゼロ値の利用
	}
	
	return wordsDist
}

func main() {
	wc.Test(WordCount)
}
