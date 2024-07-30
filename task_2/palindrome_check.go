package main

import (
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
	"unicode"

)

func isPalindrome(str string) bool{
	l := len(str)
	for i := 0; i < int(math.Floor(float64(l/2))); i++ {
		if str[i] != str[l-1-i]{
			return false
		}
	}
	return true
}

func main(){
	fmt.Println("Please insert any string you want")
	// var word string

	// fmt.Scanln(word)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil{
		return
	}

	str = strings.ToLower(str)
	var cleaned strings.Builder

	for _, char := range str{
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			cleaned.WriteRune(char)
		}
	}

	if isPalindrome(cleaned.String()){
		fmt.Println("Yes, is panlindrome")
	}else{
		fmt.Println("No, the word is not palindrome")
	}
}