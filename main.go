package main

import "fmt"
import "io/ioutil"
import "os"
import "time"


func run() {
	//TODO: generate key
	var key = "LKHlhb899Y09olUi"
	
	result := run_menu()
	switch {
	case result == 0:
		//open diary
		path, password := get_data("enter diary path:","enter password:")
		real_password  := get_diary_password(path,key)
		fmt.Println(password,real_password)

	case result == 1:
		diary, password, topic, new_page := get_data_to_add("enter diary path:",
															"enter password:",
														    "enter page topic:",
														    "enter path to page:")
		real_password := get_diary_password(diary,key)
		if real_password != clear_str(password) {
			fmt.Println("password is incorrect!", real_password)
			time.Sleep(2*time.Second)
			run()
		}

		new_strings := read_file_lines(new_page)
		add_to_file(diary, new_strings,topic,key)
		
	case result == 2:
		name, password := get_data("enter new diary name:","enter password for save data:")
		password, _ = encrypt([]byte(key),password)
		password = "["+password+"]"+"\n"
		err := ioutil.WriteFile(name,[]byte(password),0777)
		if err != nil {
			fmt.Println("can not create new diary!", err)
			os.Exit(-1)
		}
	}
}
func main() {
	execute_cmd_command("title My Secret Personal Diary")
	execute_cmd_command("cls")

	run()
}
