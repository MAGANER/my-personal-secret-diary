package main

import "fmt"
import "io/ioutil"
import "os"
import "time"
import "strings"
func run(pages *[]string, _key string) {
	var key = _key
	
	result := run_menu()
	switch {
	case result == 0:
		//open diary
		path, password := get_data("enter diary path:","enter password:")
		real_password, end  := get_diary_password(path,key)
		if password ==  real_password {
			data := get_file_data(path)
			if end+2 == len(data) {
				fmt.Println("diary is empty!")
				time.Sleep(2*time.Second)
				run(pages,key)
			} else {
				data = data[end+2:]
			}
			encrypted, err := decrypt([]byte(key), data)
			if err != nil {
				fmt.Println("something went wrong with decryption!")
			} else {
				split_result := strings.Split(encrypted,"<border>")
				split_result = split_result[:len(split_result)-1]
				pages = &split_result
				run(pages,key)
			}
		} else {
			fmt.Println("password is incorrect!", password)
			time.Sleep(2*time.Second)
			run(pages,key)
		}

	case result == 1:
		diary, password, topic, new_page := get_data_to_add("enter diary path:",
															"enter password:",
														    "enter page topic:",
														    "enter path to page:")
		real_password, _ := get_diary_password(diary,key)
		if real_password != clear_str(password) {
			fmt.Println("password is incorrect!", password)
			time.Sleep(2*time.Second)
			run(pages,key)
		}

		new_strings := read_file_lines(new_page)
		add_to_file(diary, new_strings,topic,key)
		os.Remove(new_page)
		run(pages,key)
		
	case result == 2:
		name, password := get_data("enter new diary name:","enter password for save data:")
		password, _ = encrypt([]byte(key),password)
		password = "["+password+"]"+"\n"
		err := ioutil.WriteFile(name,[]byte(password),0777)
		if err != nil {
			fmt.Println("can not create new diary!", err)
			os.Exit(-1)
		}
		run(pages,key)
	case result == 3:
		if len(*pages) == 0 {
			fmt.Println("no pages to print!Maybe diary is not open yet?")
			run(pages,key)
		}
		fmt.Println(len(*pages))
		for i := 0; i< len(*pages); i++ {
			fmt.Println("--------------")
			fmt.Println((*pages)[i])
		}
		run(pages,key)
	case result == 4:
		_pages := make([]string,0)
		run(&_pages,key)
	}
}
func main() {
	execute_cmd_command("title My Secret Personal Diary")
	execute_cmd_command("cls")

	var key = get_key()
	var pages []string
	run(&pages,key)
}
