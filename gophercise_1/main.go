package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"henderjm/gophercises/quiz"
)

func main() {
	file, err := os.Open("problems.csv")
	checkErr(err)
	defer file.Close()

	q, err := quiz.NewQuiz(file)
	checkErr(err)

	c := make(chan quiz.QuizResult, 1)

	go q.StartQuiz(c)
	select {
	case res := <-c:
		checkErr(res.Error)
		fmt.Println(res.Result)
	case <-time.After(2 * time.Second):
		log.Fatalf("\nHEY!\nYou have run out of time, please try again")
	}

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
