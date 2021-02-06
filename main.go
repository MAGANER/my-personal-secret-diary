package main

import "fmt"
import "io/ioutil"
import "os"
import "time"

func run() {
	result := run_menu()
	switch {
	case result == 0:
		path, password := get_data("enter diary path:","enter password:")
		fmt.Println(path, password)
	case result == 1:
		diary, password, topic, new_page := get_data_to_add("enter diary path:",
															"enter password:",
														    "enter page topic:",
														    "enter path to page:")
		strings := get_file_data(diary)
		real_password := strings[0]
		real_password = real_password[1:len(real_password)-1]
		if real_password != clear_str(password) {
			fmt.Println("password is incorrect!")
			time.Sleep(2*time.Second)
			run()
		}

		new_strings := get_file_data(new_page)
		add_to_file(diary, new_strings,topic)
		
	case result == 2:
		name, password := get_data("enter new diary name:","enter password for save data:")
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
