package quiz

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/tcnksm/go-input"
)

type QuizResult struct {
	Result string
	Error  error
}

type pair struct {
	question string
	answer   string
}

type Quiz struct {
	QuestionAnswer []pair
	correct        int
	wrong          int
}

func NewQuiz(file *os.File) (Quiz, error) {
	reader := bufio.NewReader(file)
	r := csv.NewReader(reader)

	var quiz Quiz

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Quiz{}, err
		}
		quiz.QuestionAnswer = append(quiz.QuestionAnswer, pair{
			question: record[0],
			answer:   record[1],
		})
	}
	return quiz, nil
}

func (q *Quiz) StartQuiz(c chan QuizResult) {
	res := new(QuizResult)
	defer close(c)
	for _, qa := range q.QuestionAnswer {
		ui := &input.UI{}

		answer, err := ui.Ask(qa.question, &input.Options{
			Required: true,
			Loop:     true,
		})
		if err != nil {
			res.Result = ""
			res.Error = err
			c <- *res
		}
		if qa.answer == answer {
			q.correct++
		} else {
			q.wrong++
		}
	}
	fmt.Println("*** FINISHED QUIZ ***")

	res.Result = fmt.Sprintf("You got %d correct and %d wrong", q.correct, q.wrong)
	res.Error = nil
	c <- *res
}
