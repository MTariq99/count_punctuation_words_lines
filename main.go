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
	var index int64 = 0
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
	temp := chunk
	fmt.Println("chunk size ", chunk+0)
	for i := 0; i < 5; i++ {
		if i == 0 {
			go CountingPunctuation(fileData[:chunk], ch1)
			res1 := <-ch1
			totalWords += res1.totalWords
			totalPunc += res1.totalPunc
			totalLines += res1.totalLines
		} else {
			index = chunk
			chunk += temp
			go CountingPunctuation(fileData[index:chunk], ch1)
			res1 := <-ch1
			totalWords += res1.totalWords
			totalPunc += res1.totalPunc
			totalLines += res1.totalLines
		}
	}
	fmt.Println("words----> ", totalWords, "\nPuncuations-----> ", totalPunc, "\nLines-----> ", totalLines+1, "\nBytes read------> ", fileBytes)
	fmt.Println("execution time ", time.Since(start))

}
