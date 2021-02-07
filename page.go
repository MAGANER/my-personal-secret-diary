package main

import(
	"strings"
	"io/ioutil"
	"fmt"
	"os"
	"time"
)

type Page struct {
	Body  string
	Topic string
	Data  string
}

func get_file_data(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("something went wrong with file ",path, err)
		os.Exit(-1)
	} 
	lines := strings.Split(string(file),"\n")
	return lines
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