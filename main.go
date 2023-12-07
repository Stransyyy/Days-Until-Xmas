package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// Prints the Ascii art. strings.REplaceAll() replaces characters, in this case I'm using it because the ascii art I selected has back ticks symbols and that was messing up the multi line string
func bob() string {
	s := `
	_____________,--,
	| | | | | | |/ .-.\   HANG IN THERE
	|_|_|_|_|_|_/ /   $.      SANTA
	 |_|__|__|_; |      \
	 |___|__|_/| |     .'$}
	 |_|__|__/ | |   .'.'$\
	 |__|__|/  ; ;  / /    \.-"-.
	 ||__|_;   \ \  ||    /$___. \
	 |_|___/\  /;.$,\\   {_'___.;{}
	 |__|_/ $;$__|$-.;|  |C$ e e$\
	 |___$L  \__|__|__|  | $'-o-' }
	 ||___|\__)___|__||__|\   ^  /$\
	 |__|__|__|__|__|_{___}'.__.$\_.'}
	 ||___|__|__|__|__;\_)-'$\   {_.-;
	 |__|__|__|__|__|/$ ($\__/     '-'
	 |_|___|__|__/$      |
-------|__|___|__/$         \-------------------
-.__.-.|___|___;$            |.__.-.__.-.__.-.__
|     |     ||             |  |     |     |
-' '---' '---' \             /-' '---' '---' '--
   |     |    '.        .' |     |     |     |
'---' '---' '---' $-===-'$--' '---' '---' '---'
|     |     |     |     |     |     |     |
-' '---' '---' '---' '---' '---' '---' '---' '--
   |     |     |     |     |     |     |     |
'---' '---' '---' '---' '---' '---' '---' '---'
`

	return strings.ReplaceAll(s, `$`, "`")
}

func main() {

	fmt.Println(bob())

	hoy := time.Now().Unix()

	fecha := time.Unix(hoy, 0).Format("2006-01-02 15:04:05")

	figure.NewColorFigure("Hello -ho -ho -ho", "", "red", true).Print()

	fmt.Printf("\n \nToday is %v", fecha)

	fmt.Print("\n\n")

	preguntas()

}

func preguntas() {

	if yesOrNoQuestion("Would you like to know how many days left until Christmas?\n\nAnswer with a (yes/no) please.\n\n") {

		ifYesContinue("\nGreat. Would you like how many days until XMAS in days/minutes/seconds?")

	} else {
		fmt.Print("\nExiting the program. Goodbye! Santa will not miss you.\n\n")
	}
}

func yesOrNoQuestion(question string) bool {
	fmt.Print(question)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer := strings.ToLower(scanner.Text())
		if answer == "yes" {
			return true
		} else if answer == "no" {
			return false
		} else {
			fmt.Print("Invalid input. Please enter 'yes' or 'no': ")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// Default to no if an error occurs
	return false
}

func ifYesContinue(question string) bool {

	fmt.Print(question)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		answer := strings.ToLower(scanner.Text())
		switch answer {
		case "days":
			if answer == "days" {
				daysUntilChristmas()
			}
		case "minutes":
			if answer == "minutes" {
				minutesUntilChristmas()
			}
		case "seconds":
			if answer == "seconds" {
				secondsUntilChristmas()
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

	// Default to no if an error occurs
	return false

}

func secondsUntilChristmas() {

	hoy := time.Now().Local()

	christmas := time.Date(hoy.Year(), time.December, 25, 0, 0, 0, 0, time.Local)

	untilChristmas := christmas.Sub(hoy)

	// Extract the number of days from the duration
	daysUntilChristmas := int(untilChristmas.Seconds())

	fmt.Printf("Seconds until Christmas: %d\n", daysUntilChristmas)

}

func minutesUntilChristmas() {

	hoy := time.Now().Local()

	christmas := time.Date(hoy.Year(), time.December, 25, 0, 0, 0, 0, time.Local)

	untilChristmas := christmas.Sub(hoy)

	// Extract the number of days from the duration
	daysUntilChristmas := int(untilChristmas.Minutes())

	fmt.Printf("Minutes until Christmas: %d\n", daysUntilChristmas)

}

// Days until Christmas
func daysUntilChristmas() {

	hoy := time.Now().Local()

	christmas := time.Date(hoy.Year(), time.December, 25, 0, 0, 0, 0, time.Local)

	untilChristmas := christmas.Sub(hoy)

	// Extract the number of days from the duration
	daysUntilChristmas := int(untilChristmas.Hours() / 24)

	fmt.Printf("Days until Christmas: %d\n", daysUntilChristmas)

}

// UNIX timestamp 1703484000 (until christmas)
