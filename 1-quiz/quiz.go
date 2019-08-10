/*
Input on file name to parse as csv

Then print out how many total problems and how many right answers 

And add a timer via channel 
Solution https://github.com/gophercises/quiz/blob/solution-p1/main.go

Parse CSV Golang https://www.thepolyglotdeveloper.com/2017/03/parse-csv-data-go-programming-language/

Timing out using channel https://blog.golang.org/go-concurrency-patterns-timing-out-and
https://telliott.io/2016/09/29/three-ish-ways-to-implement-timeouts-in-go.html
Can also use the Timer in time package 

Note there are alternative ways to parse csv and add timer 

go build . && ./1-quiz --limit 5 
./1-quiz --help 
./1-quiz -csv="abc.csv"
*/


package main 

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

// desclare flags 
var fileName = flag.String("filename", "problems.csv", "a csv in the form of `question, answer`. Default problems.csv")
var limit = flag.Int("limit",10, "timer. Default 10 seconds")

var countProblems int = 0
var countRight int = 0 


func main() {
	// parse the flags 
	flag.Parse()

	// csv filename is a pointer to a string (that's how flag works )
	csvFile, err := os.Open(*fileName)
	if err != nil {
		// similar to printf and exit
		log.Fatalf("Failed to open the CSV file: %s\n", *fileName)
	}
	// create a buffer reader (not necessary though since quiz won't block memory)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	
	// ch := make(chan bool, 1)
	timeout := make(chan bool, 1)	

	go func() {
		time.Sleep(time.Duration(*limit) * time.Second)
		timeout <- true
	}()

	for {
		line, err := reader.Read() 
		if err == io.EOF {
			fmt.Printf("problems %d and right answers %d \n", countProblems, countRight)
			break
		} else if err !=nil {
			log.Fatal(err)
		}

		fmt.Println(line[0]) 
		countProblems +=1

		answerCh := make(chan string)
		go func() {
			var answer string
			// scanf get rids of spaces and capture new line
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		
		select {
			case <-timeout:
				fmt.Printf("problems %d and right answers %d \n", countProblems, countRight)
				fmt.Println("Timed out")
				return 
			case answer := <-answerCh:
				if answer == line[1] {
					countRight +=1
				}
		}
	
	}
		

}

