package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func getWords(numWords int) []string {
	wordsURL := fmt.Sprintf("https://random-word-api.herokuapp.com/word?number=%d", numWords)
	resp, err := http.Get(wordsURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	words := make([]string, 0)
	err = json.NewDecoder(resp.Body).Decode(&words)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return words
}

func getNums(numNums int) []string {
	nums := make([]string, numNums)
	for i := 0; i < numNums; i++ {
		num := rand.Intn(100)
		nums[i] = fmt.Sprintf("%d", num)
	}
	return nums
}

func getLetter() string {
	letterNum := 65 + rand.Intn(26)
	return string(rune(letterNum))
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	words := getWords(2)
	nums := getNums(2)
	letter := getLetter()
	fmt.Println(words[0] + nums[0] + letter + nums[1] + words[1])
}
