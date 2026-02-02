func main() {
	menu := [][]string{{"Show", "Add"}, {"Delete", "Exit"}}
	x, y := 0, 0
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J") // clear screen
		for i := 0; i < len(menu); i++ {
			for j := 0; j < len(menu[i]); j++ {
				if i == y && j == x {
					fmt.Printf("→%s\t", menu[i][j])
				} else {
					fmt.Printf("  %s\t", menu[i][j])
				}
			}
			fmt.Println()
		}
		fmt.Print("\nUse W/A/S/D to move, Enter to select: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		switch strings.ToLower(input) {
		case "w":
			if y > 0 {
				y--
			}
		case "s":
			if y < len(menu)-1 {
				y++
			}
		case "a":
			if x > 0 {
				x--
			}
		case "d":
			if x < len(menu[y])-1 {
				x++
			}
		case "":
			switch menu[y][x] {
			case "Show":
				fmt.Println("Showing notes...")
			case "Add":
				fmt.Println("Adding note...")
			case "Delete":
				fmt.Println("Deleting note...")
			case "Exit":
				fmt.Println("Goodbye!")
				return
			}
			fmt.Println("Press Enter to continue...")
			scanner.Scan()
		}
	}
}