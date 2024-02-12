package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Xatolik: Faylni ochib bo'lmadi:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var nums []int
	var words []string
	for scanner.Scan() {
		line := scanner.Text()

		for _, word := range strings.Fields(line) {
			if num, err := strconv.Atoi(word); err == nil {
				nums = append(nums, num)
			} else {
				words = append(words, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik: Faylni o'qib bo'lmadi:", err)
		return
	}

	numsFile, err := os.Create("numbers_output.txt")
	if err != nil {
		fmt.Println("Xatolik: Sonlar faylini yozib bo'lmadi:", err)
		return
	}
	defer numsFile.Close()

	var s int
	numsWriter := bufio.NewWriter(numsFile)
	for _, num := range nums {
		s += num
		fmt.Fprintln(numsWriter, num)
	}
	numsWriter.Flush()
	fmt.Fprint(numsFile, s)

	wordsFile, err := os.Create("words_output.txt")
	if err != nil {
		fmt.Println("Xatolik: So'zlar faylini yozib bo'lmadi:", err)
		return
	}
	defer wordsFile.Close()

	wordsWriter := bufio.NewWriter(wordsFile)
	for _, word := range words {
		if strings.Contains(word, "a") || strings.Contains(word, "o") {
			fmt.Fprintln(wordsWriter, word)
		}
	}
	wordsWriter.Flush()

	fmt.Println("Sonlar fayliga yozildi: numbers_output.txt")
	fmt.Println("So'zlar fayliga yozildi: words_output.txt")
}
