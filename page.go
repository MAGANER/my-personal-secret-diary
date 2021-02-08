package main

import(
	"io/ioutil"
	"fmt"
	"os"
	"time"
	"strings"
)

type Page struct {
	Body  string
	Topic string
	Data  string
}

func is_file_ok(path string, err error) {
	if err != nil {
		fmt.Println("something went wrong with file ",path, err)
		os.Exit(-1)
	} 
}
func get_diary_password(path,key string) string {
	file, err := ioutil.ReadFile(path)
	is_file_ok(path,err)
	

	_file:= string(file)

	begin := strings.Index(_file,"[")
	end   := strings.Index(_file,"]")

	if begin == -1 || end == -1 {
		return "";
	}

	_key := []byte(key)
	password := _file[begin+1:end]
	password, _ = decrypt(_key,password)
	return  password

}
func get_file_data(path, password string) string {
	file, err := ioutil.ReadFile(path)
	is_file_ok(path,err)
	key   := []byte(password)
	lines := string(file)
	lines, _  = decrypt(key,lines) 
	return lines
}
func read_file_lines(path string) []string {
	file, err := ioutil.ReadFile(path)
	is_file_ok(path,err)
	return strings.Split(string(file),"\n")
}
func add_to_file(path string, strings []string, topic, password string) {

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644) 
	if err != nil {
		fmt.Println("can not open file!",err)
	}

	key := []byte(password)
	for i := 0; i < len(strings); i++ {
		data := strings[i]+"\n"
		if i == 0 {
			data = "\n <border>" + data
		}
		data, err = encrypt(key,data)
		if err != nil {
			fmt.Println("something went wrong!")
		}
		f.WriteString(data)
	}
	_topic, _ := encrypt(key,"topic:"+topic+"\n <border>")
	curr_time := time.Now().String()
	date,   _ := encrypt(key,"date:"+curr_time+"\n")
	f.WriteString(date)
	f.WriteString(_topic)
}