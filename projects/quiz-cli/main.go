package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'ques,ans' ")
	limit := flag.Int("l", 5, "limit for the quiz")
	flag.Parse()

	var username, password string

	initial(&username, &password)

	login(&username, &password)

	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open the csv file %s", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("failed to parse the provided csv file")
	}

	problems := parseProblems(lines)

	// RANDOM SHUFFLING QUESTIONS
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	count := 0

	for i, pro := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, pro.q)
		answerCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanf("%s", &ans)
			answerCh <- ans
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n total correct ans:= %d , total problems:= %d \n", count, len(problems))
			leaderboard(username, count)
			return
		case answer := <-answerCh:
			if answer == pro.a {
				fmt.Println("correct ans")
				count++
			}
		}
	}
	fmt.Printf("\n total correct ans:= %d , total problems:= %d \n", count, len(problems))
	leaderboard(username, count)
}

type Problem struct {
	q string
	a string
}

type Player struct {
	Username string
	Score    int
}

func parseProblems(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for i, line := range lines {
		ret[i] = Problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func login(username, password *string) {
	for *username == "" {
		fmt.Println("Please enter a username:")
		fmt.Scanf("%s", username)
	}
	for *password == "" {
		fmt.Println("please enter a password:")
		fmt.Scanf("%s", password)
	}
	file, err := os.Open("./creds.text")
	defer file.Close()
	if err != nil {
		exit("internal server error")
	} else {
		isValid := checkCreds(*username, *password, file)

		if !isValid {
			fmt.Println("invalid input, Pleas try again :)")
			*username, *password = "", ""
			initial(username, password)
		}
	}
}

func leaderboard(username string, score int) {
	fileName := "score.txt"

	// Open the file in append mode, create it if it doesn't exist
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Write the new score to the file
	_, err = file.WriteString(fmt.Sprintf("%s,%d\n", username, score))
	if err != nil {
		log.Fatalf("failed to write to file: %s", err)
	}

	// Now, read all the scores from the file
	players := readScores(fileName)

	// Sort the players by score in descending order
	sort.Slice(players, func(i, j int) bool {
		return players[i].Score > players[j].Score
	})

	// Display the scores in a tabular format
	displayScores(players)
}

// Function to read all scores from the file and return a slice of Player structs
func readScores(fileName string) []Player {
	// Open the file for reading
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	var players []Player

	// Scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) == 2 {
			username := parts[0]
			score, err := strconv.Atoi(parts[1])
			if err == nil {
				players = append(players, Player{Username: username, Score: score})
			}
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}

	return players
}

func displayScores(players []Player) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.AlignRight)

	// Print the headers
	fmt.Fprintln(w, "Username\tScore\t")

	// Print each player's score
	for _, player := range players {
		fmt.Fprintf(w, "%s\t%d\t\n", player.Username, player.Score)
	}

	// Flush the writer to display the data
	w.Flush()
}

func checkCreds(username, password string, file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}
		fileUsername := parts[0]
		filePassword := parts[1]

		if fileUsername == username {
			if filePassword == password {
				return true
			}
		}
	}
	return false
}

func register(username, password *string) {
	// open the file and check if its already exist then give the error
	for *username == "" {
		fmt.Println("Please enter a username:")
		fmt.Scanf("%s", username)
	}

	// Open the file for reading and writing (O_RDWR), create if doesn't exist
	file, err := os.OpenFile("./creds.text", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		exit("internal server error")
	}
	defer file.Close()

	// Check if username exists
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			// Skip malformed lines
			continue
		}
		fileUsername := parts[0]

		// Check if the username already exists
		if fileUsername == *username {
			fmt.Println("Username already exists, please try again.")
			*username = ""
			*password = ""
			initial(username, password) // Assuming this prompts for a new username/password
			return
		}
	}

	// If username doesn't exist, proceed to ask for password
	for *password == "" {
		fmt.Println("Please enter a password:")
		fmt.Scanf("%s", password)
	}

	// Append the new username and password to the file
	data := fmt.Sprintf("%s,%s\n", *username, *password) // Add a new line
	_, err = file.WriteString(data)
	if err != nil {
		exit("Unable to register.")
	}

	fmt.Println("Registration successful!")
}

func initial(username, password *string) {
	var opt string
	for opt != "1" && opt != "2" {
		fmt.Println("Please choose:")
		fmt.Println("1 #Register")
		fmt.Println("2 #Login")
		fmt.Scanf("%s", &opt)
	}
	if opt == "1" {
		register(username, password)
	} else {
		login(username, password)
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
