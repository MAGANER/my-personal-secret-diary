package main

import(
	"fmt"
	"bufio"
	"os"
	"time"
)

func print_help() {
	fmt.Println("	My Secret Personal Diary")
	fmt.Println("		version 1.0")
	fmt.Println("		developed by github.com/MAGANER")
	fmt.Println("Commands:")
	fmt.Println("help == see this")
	fmt.Println("open == open existing diary")
	fmt.Println("add  == add new page to diary")
}
func get_data(file, question string) (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(file)
	path, _ := reader.ReadString('\n')
	fmt.Print(question)
	second, _ := reader.ReadString('\n')

	return clear_str(path),clear_str(second)
}
func run_menu() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		val, _ := reader.ReadString('\n')
		val = clear_str(val)

		if val == "help" {
			print_help()
			time.Sleep(5*time.Second)
		}
		if val == "open" {
			return 0
		}
		if val == "add" {
			return 1
		}

		execute_cmd_command("cls")
	}
}
