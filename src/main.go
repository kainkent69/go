package main

func main() {
	Clear()
	tasks = make(map[int]*Task)
	Welcome()
	Show_Options()
	Run()
}
