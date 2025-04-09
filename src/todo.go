package main

import (
	"bufio"
	"fmt"
	"os"
)

var tasks map[int]*Task

func Clear() {
	fmt.Printf("\x1b[2J\x1b[H")
}
func Show_Options() {
	fmt.Println("[__________________________________]")
	fmt.Println("|           Task Options           |")
	fmt.Println("|__________________________________|")
	fmt.Println("|                                  |")
	fmt.Println("|[0]  Exit          [4]  Options   |")
	fmt.Println("|[1]  Read          [5]  New       |")
	fmt.Println("|[2]  Delete        [6]  List      |")
	fmt.Println("|[3]  Edit                         |")
	fmt.Println("[___________ ______________________]")

}

const DescriptionLimit int = 120
const TitileLimt int = 50

func Welcome() {
	fmt.Println("Welcome to the  Todo!")
	fmt.Printf("Select [Any Key] To Continue: ")
	var str string
	fmt.Scanf("%s", &str)
}

func ScanInput() string {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return string([]byte(str)[0])
}

func Run() {
	fmt.Print("What Do You Want To Do? ")
	input := ScanInput()
	task := new(Task)
	toExitAndShowOptions := func() {
		fmt.Printf("Enter Anything To Exit: ")
		ScanInput()
		Clear()
		Show_Options()
	}
	processInput := func() bool {
		switch input {
		case "0":
			fmt.Println("Todo: Exiting...")
			os.Exit(0)
			break

		case "5":
			task.New()
			toExitAndShowOptions()
			break

		case "6":
			Clear()
			task.ShowLists()
			toExitAndShowOptions()
			break

		case "3":
			Clear()
			task.Edit()
			break
		default:
			Clear()
			Show_Options()
			fmt.Printf("Todo: %q is invalid\n ", input)
			return false
		}

		return true
	}

	if !processInput() {
	}
	Run()
}
