package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type question struct {
	question string
	answer   string
}

func main() {

	// 读取 csv 中的 问题和答案
	var questions []question

	content, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		log.Fatalln("Read csv file error")
	}

	r := csv.NewReader(strings.NewReader(string(content)))

	for {
		line, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Read csv file error")
		}
		qa := question{
			question: line[0],
			answer:   line[1],
		}
		questions = append(questions, qa)
	}

	// cli 中对 clinet 进行提问
	score := 0
	for _, qa := range questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(qa.question)
		for {
			text, err := reader.ReadString('\n')
			if err != nil {
				log.Fatalln("Read console input error")
			}
			if strings.EqualFold(qa.answer, strings.Trim(text, "\n")) {
				score++
				fmt.Println("Answer is right")
				break
			}
			fmt.Println("Wrong answer should be ", qa.answer)
			break
		}
	}
	fmt.Printf("total question count : %d,total score is %d \n", len(questions), score)

}
