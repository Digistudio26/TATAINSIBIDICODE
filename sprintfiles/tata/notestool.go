package main

import(
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func printHelp() {
	fmt.Println("Usage: ./notestool [TAG]")
	fmt.Println()
	fmt.Println("Create or manage a collection of notes.")
	fmt.Println("Example: ./notestool coding_ideas")
}


func main() {
	if len(os.Args) != 2 {
		fmt.Println("\033[31mError: wrong number of arguments\033[0m\n")
		printHelp()
		os.Exit(1)
	}

	tag := os.Args[1]

	if strings.ToLower(tag) == "help" {
		printHelp()
		os.Exit(0)
	}

	notes := loadOrCreateCollection(tag)
	runMenuLoop(tag, notes)
}


func loadOrCreateCollection(filename string) []string {
	filePath := filename + ".txt"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.WriteFile(filePath, []byte{}, 0644)
		return []string{}
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(content), "\n")
	var notes []string

	for _, line := range lines {

		trimmed := strings.TrimSpace(line);
		if  trimmed != "" {
			notes = append(notes, trimmed)
		}
	}

	return notes
}



func showWelcome(tag string) {
	fmt.Println("\033[1;35mWelcome to the notes tool!\033[0m")
	fmt.Println()
}

func displayMenu() {
	fmt.Println("Select operation:")
	fmt.Println("1. Show notes.")
	fmt.Println("2. Add a note.")
	fmt.Println("3. Delete a note.")
	fmt.Println("4. Exit.")
	fmt.Println()
}

func displayNotes(notes []string) {
	fmt.Println("Notes:")
	if len(notes) == 0 {
		fmt.Println("\033[33m (empty note) \033[0m")
		return
	}

	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
	fmt.Println()
}


func clearScreen() {
	fmt.Print("\033[H\033[2J")
}


func runMenuLoop(tag string, notes []string) {

	clearScreen()
	showWelcome(tag)
	scanner := bufio.NewScanner(os.Stdin)

	for {

		displayMenu()

		fmt.Print("Choose 1 - 4: ")
		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())
		clearScreen()

		switch choice {
		case "1":
			clearScreen()
			displayNotes(notes)
		case "2":
			notes = addNote(notes, tag)
		case "3":
			notes = deleteNote(notes, tag)
		case "4":
			fmt.Println("\033[1;35mGoodbye!\033[0m")
			return
		default:
			fmt.Println("\033[31mInvalid. Please select 1-4.\033[0m")
		}
	}
}


func saveNotes(notes []string, tag string) error {
	filePath := tag + ".txt"
	content := strings.Join(notes, "\n")

	if len(notes) > 0 {
		content += "\n"
	}

	return os.WriteFile(filePath, []byte(content), 0644)
}

func addNote(notes []string, tag string) []string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nEnter the note you want to add: ")
	scanner.Scan()
	text := strings.TrimSpace(scanner.Text())

	if text == "" {
		fmt.Println("\n\033[31mYou did not add anything\033[0m\n")
		return notes
	}

	timestamp := time.Now().Format("Mon Jan _2 15:04:05 2006")
	noteWithTime := fmt.Sprintf("%s - \033[32m[%s]\033[0m", text, timestamp)

	notes = append(notes, noteWithTime)

	//notes = append(notes, text)

	err := saveNotes(notes, tag)
	if  err != nil {
		fmt.Printf("\n\033[31mError saving: %v\033[0m\n\n", err)
	}

	fmt.Printf("\033[32mNote added! (Total: %d)\033[0m\n", len(notes))
	return notes
}


func deleteNote(notes []string, tag string) []string {
	if len(notes) == 0 {
		fmt.Println("\n\033[31mNo notes to delete.\033[0m")
	return notes
	}

	displayNotes(notes)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the number to remove or 0 to cancel: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	num, err := strconv.Atoi(input)
	if err != nil || num < 0 || num > len(notes) {
		fmt.Println("\n\033[31mInvalid number.\033[0m\n")
		return notes
	}

	if num == 0 {
		fmt.Println("\n\033[31mCancelled.\033[0m")
		return notes
	}

	index := num - 1
	deletedNote := notes[index]
	notes = append(notes[:index], notes[index+1:]...)

	
	if  err := saveNotes(notes, tag); err != nil {
		fmt.Printf("\n\033[31mError saving: %v\033[0m\n", err)
	}

	fmt.Printf("\nDeleted: %03d - %s\n", num, deletedNote)
	fmt.Printf("Remaining: %d notes\n", len(notes))
	return notes
}