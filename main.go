package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of question,answer")

	flag.Parse()

	// Open the mentioned file
	file, err := os.Open(*csvFileName)

	if err != nil {
		ErrExit(fmt.Sprintf("Unable to open file %s : %s\n", *csvFileName, err))

	}

	// Closes the file after the reading it
	defer file.Close()

	// Read the csv records data
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {

		ErrExit(fmt.Sprintf("Unable to parse csv file %s : %s\n", *csvFileName, err))
	}

	QuizGame(records)
}

// The quiz game is performed here
// after taking the csv data records
func QuizGame(records [][]string) {

	total := len(records)
	correct := 0
	solved := 0

	fmt.Println("Welcome To Quiz Game !")
	fmt.Println("Type 's' to stop the game.\n Any other value will be considered an answer\n")

	for i, values := range records {

		fmt.Printf("Question #%d : %s\n", i+1, values[0])
		fmt.Printf("Answer : ")
		ans := ReadInput()

		if ans == "s" || ans == "S" {
			fmt.Println("Game stopped!")
			break
		}

		if ans == values[1] {
			correct++
		}
		solved++
		fmt.Println()
	}

	if solved == total {
		fmt.Println("Bravo!! You attempted all the questions.\n")
	}

	fmt.Println("---------------------------------")
	fmt.Println("Total questions solved : ", solved)
	fmt.Println("Total correct answers  : ", correct)

	fmt.Println("---------------------------------")
}

// Reads the input from command line and returns as string
func ReadInput() string {

	var input string

	fmt.Scanf("%s", &input)

	return input

}

func ErrExit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
