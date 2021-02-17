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
func get_diary_password(path,key string) (string, int) {
	file, err := ioutil.ReadFile(path)
	is_file_ok(path,err)
	

	_file:= string(file)

	begin := strings.Index(_file,"[")
	end   := strings.Index(_file,"]")

	if begin == -1 || end == -1 {
		return "", -1;
	}

	_key := []byte(key)
	password := _file[begin+1:end]
	password, _ = decrypt(_key,password)
	return  password, end

}
func get_file_data(path string) string {
	file, err := ioutil.ReadFile(path)
	is_file_ok(path,err)
	lines := string(file)
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
	united_str := ""
	for i := 0; i < len(strings); i++ {
		united_str += strings[i] + "\n"
	}
	curr_time := time.Now().String()
	data := united_str+"topic:"+topic+"\n"+"date:"+curr_time+"\n"+"<border>"+"\n"
	data, err = encrypt(key,data)
	data += "<bor>"
	if err != nil {
		fmt.Println("can not encrypt data!")
	}
	f.WriteString(data)
}