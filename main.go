package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main(){

	csvFileName := flag.String("csv", "problems.csv",
		"A csv file in the format of a question/answer format")


	flag.Parse()

	_ = csvFileName

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file %s\n", *csvFileName))

	}

	read := csv.NewReader(file)

	lines, err := read.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file")
	}

	Problems := parseLines(lines)

	correct := 0

	for i, p := range Problems {
		fmt.Printf("Problem #%d: %s = \n", i + 1, p.Question )
		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == p.Answer {
			correct ++
		}
	}

	fmt.Printf("You scored %d our of %d.\n", correct, len(Problems))
}


func parseLines(lines [][]string) []Problems {

	ret := make([]Problems, len(lines))

	for i, line := range lines {
		ret[i] = Problems{
			Question: line[0],
			Answer: strings.TrimSpace(line[1]),
		}
	}

	return ret

}

type Problems struct {
	Question string
	Answer string
}

func exit(msg string){

	fmt.Println(msg)
	os.Exit(1)
}