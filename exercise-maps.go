// https://tour.golang.org/moretypes/23
// Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.
package main

import (
    "strings"
    "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
    wordCount := make(map[string]int)

    words := strings.Fields(s)
    for _, word := range words {
        wordCount[word] += 1
    }

    return wordCount
}

func main() {
    wc.Test(WordCount)
}

