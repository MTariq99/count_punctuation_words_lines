package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type result struct {
	totalWords int
	totalPunc  int
	totalLines int
}

func CountingPunctuation(fileData []byte, channel chan result) {
	tw, tp, tl := 0, 0, 0
	for i := 0; i < len(fileData); i++ {
		if fileData[i] == ' ' || fileData[i] == '\n' {
			tw++
		} else if fileData[i] == ',' || fileData[i] == '.' {
			tp++
		}
		if fileData[i] == '\n' {
			tl++
		}
	}
	res := result{
		totalWords: tw,
		totalPunc:  tp,
		totalLines: tl,
	}

	channel <- res

}

func main() {
	totalWords, totalPunc, totalLines := 0, 0, 0
	file, err := os.Open("file.txt")
	fileSize, _ := os.Stat("file.txt")
	fileData := make([]byte, fileSize.Size())
	if err != nil {
		fmt.Println("file reading error ", err)
	}
	if err == io.EOF {
		return
	}
	defer file.Close()
	fileBytes, _ := file.Read(fileData)

	ch1 := make(chan result)
	start := time.Now()
	chunk := fileSize.Size() / 5
	for i := 0; i < 5; i++ {
		go CountingPunctuation(fileData[:chunk], ch1)
		res1 := <-ch1
		totalWords += res1.totalWords
		totalPunc += res1.totalPunc
		totalLines += res1.totalLines
	}
	fmt.Println("words----> ", totalWords, "\nPuncuations-----> ", totalPunc, "\nLines-----> ", totalLines, "\nBytes read------> ", fileBytes)
	fmt.Println("execution time ", time.Since(start))

	// go CountingPunctuation(fileData[:13360688], ch1)
	// go CountingPunctuation(fileData[13360688:26721376], ch1)
	// go CountingPunctuation(fileData[26721376:40082064], ch1)
	// go CountingPunctuation(fileData[40082064:53442752], ch1)
	// go CountingPunctuation(fileData[53442752:], ch1)

	// res1 := <-ch1
	// res2 := <-ch1
	// res3 := <-ch1
	// res4 := <-ch1
	// res5 := <-ch1
	// fmt.Println("\ntotal words ", res1.totalWords+res2.totalWords+res3.totalWords+res4.totalWords+res5.totalWords+1)
	// fmt.Println("\ntotal punctuations ", res1.totalPunc+res2.totalPunc+res3.totalPunc+res4.totalPunc+res5.totalPunc)
	// fmt.Println("\ntotal lines ", res1.totalLines+res2.totalLines+res3.totalLines+res4.totalLines+res5.totalLines)
	// fmt.Println("\ntotal bytes used ", fileBytes)
	// fmt.Println("execution time ", time.Since(start))

}
