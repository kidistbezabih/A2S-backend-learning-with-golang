package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func WordFreq(str string) map[string]int{
	words := strings.Fields(str)
	
	freq := make(map[string]int)

	for _, elem := range words{
		freq[strings.ToLower(elem)]++
	}

	return freq
}

func main() {
	fmt.Println("Please insert any string you want")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil{
		fmt.Println("There is an error", err)
		return
	}
	
	wordList := strings.TrimSpace(str)
	freq := WordFreq(wordList)

	fmt.Println(wordList, freq)

	fmt.Println("Words and their frequency \n -----------------------------------")

	for word, count := range freq{
		fmt.Printf("%s %d\n", word, count)
	}

}