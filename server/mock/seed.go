package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Task struct {
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func main() {
	dbConn := openConnection()

	// load data
	tasks := LoadTasksFromJson()

	arg := os.Args[1]

	if arg == "-i" {
		fmt.Println("Create Data")
		dbConn.Create(&tasks)
		return
	} else if arg == "-d" {
		fmt.Println("Destroy database")
		dbConn.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&tasks)
		return
	}
}

func openConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("taskTracker.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func ParseJSON(filename string) []byte {
	jsonFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened ", filename)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}

func LoadTasksFromJson() []Task {
	byteValue := ParseJSON("mock/tasks.json")
	var tasks []Task
	json.Unmarshal([]byte(byteValue), &tasks)
	return tasks
}
