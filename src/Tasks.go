package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

var SIZE int = 0

type Task struct {
	Title       string
	Description string
	Duration    time.Duration
	Done        bool
	ID          int
}

// / Show The list of Tasks
func (t *Task) ShowLists() {
	if SIZE <= 0 {
		fmt.Println("The Token List: \"EMPTY\"")
		return
	}

	fmt.Println("The Tokens List: ")
	for i := range SIZE {
		task := tasks[i]
		if task != nil {
			fmt.Printf("[%d] %q\n", task.ID, task.Title)
		}
	}
}

func (t *Task) New() {
	fmt.Println("Task.New: Creating a new Task")
	// getting the title
	readTitle := func() bool {
		fmt.Print("Task.Title?: ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Task.New: Error or content is too long")
			reader.Discard(0)
			return false
		}
		t.Title = string(line)
		return true
	}

	for {
		done := readTitle()
		if done {
			break
		}
	}
	// read the body

	setBody := func(str string) {
		t.Description = str
	}
	readBody := func() bool {
		fmt.Print("Task.Body?: ")
		buf := make([]byte, 1)

		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			if err := scanner.Err(); err != nil {
				if err == io.EOF {
					setBody(string(buf))
					return true
				}
				return false
			} else if line == "done" || line == "end" {
				setBody(string(buf))
				return true
			}
			buf = append(buf, []byte(line)...)
		}
		setBody(string(buf))
		return true
	}

	// read the body
	for {
		done := readBody()
		if !done {
			fmt.Printf("Task.New: Sorry Something went wrong")
		} else {
			break
		}
	}

	t.ID = SIZE
	tasks[SIZE] = t
	SIZE++
	fmt.Printf("Task.New: Yay! A new Task [%d] %q is Added", t.ID, t.Title)
	fmt.Print("\n")

}

func (t *Task) Edit() {
	if SIZE <= 0 {
		fmt.Println("Task.Edit: List Is Empty  ")
		return
	}

	t.ShowLists()
	fmt.Println("")
	fmt.Printf("Choose The ID: ")
	ID := ScanInput()
	_ = ID
	fmt.Println("[__________________________________]")
	fmt.Println("|           Task Options           |")
	fmt.Println("|__________________________________|")
	fmt.Println("|                                  |")
	fmt.Println("|[0]  Exit          [4]  Mark Done |")
	fmt.Println("|[1]  Edit Title    [5]  Delete    |")
	fmt.Println("|[2]  Edit Body                    |")
	fmt.Println("|[3]  Edit Time                    |")
	fmt.Println(F_BOTTOM)

	input := ScanInput()
	switch input {
	default:
		fmt.Printf("Task.Edit %q Is Invalid\n", input)
	}

}
