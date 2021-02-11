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
	fmt.Println("add  == add new page to existing diary")
	fmt.Println("make == make new diary")
	fmt.Println("read == read all diary")
	fmt.Println("quit == close diary")
}
func get_data(file, question string) (string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(file)
	path, _ := reader.ReadString('\n')
	fmt.Print(question)
	second, _ := reader.ReadString('\n')

	return clear_str(path),clear_str(second)
}
func get_data_to_add(file, question, topic, page_path string) (string, string, string, string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(file)
	path, _ := reader.ReadString('\n')
	fmt.Print(question)
	password, _ := reader.ReadString('\n')
	fmt.Print(topic)
	_topic, _ := reader.ReadString('\n')
	fmt.Print(page_path)
	page, _ := reader.ReadString('\n')

	return clear_str(path),clear_str(password),clear_str(_topic), clear_str(page)
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
		if val == "make" {
			return 2
		}
		if val == "read" {
			return 3
		}
		if val == "quit" {
			return 4
		}

		execute_cmd_command("cls")
	}
}
