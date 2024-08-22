package repository

import (
	"encoding/json"
	"fmt"
	"github.com/HoseinAsadolahi/task-tracker/pkg/model"
	"github.com/HoseinAsadolahi/task-tracker/pkg/utils"
	"github.com/HoseinAsadolahi/task-tracker/statics"
	"os"
	"path"
)

func getFilePath() string {
	wd, err := os.Getwd()
	if utils.CheckError(err, "Can't reach current working directory") {
		os.Exit(1)
	}
	return path.Join(wd, "tasks.json")
}

func setOriginToZero(file *os.File) {
	err := file.Truncate(0)
	if utils.CheckError(err, "Can't truncate the file!") {
		os.Exit(1)
	}
	_, err = file.Seek(0, 0)
	if utils.CheckError(err, "Can't set the pointer to the beginning of the file!") {
		os.Exit(1)
	}
}

func exists() bool {
	filePath := getFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func createIfNotExists() {
	filePath := getFilePath()
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println(statics.InfoStyle.Render("File does not exist. creating the tasks.json ..."))
		if _, err := os.Create(filePath); utils.CheckError(err, fmt.Sprintf("Can't Create File: %s!",
			filePath)) {
			os.Exit(1)
		}
	}
}

func closeFile(file *os.File) {
	err := file.Close()
	if utils.CheckError(err, "Can't Close File: "+getFilePath()+"!") {
		os.Exit(1)
	}
}

func open() *os.File {
	file, err := os.OpenFile(getFilePath(), os.O_RDWR, os.ModePerm)
	if utils.CheckError(err, "Can't Open File: "+getFilePath()+"!") {
		os.Exit(1)
	}
	return file
}

func readTasks(file *os.File) ([]model.Task, error) {
	var tasks []model.Task
	err := json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(desc string) {
	createIfNotExists()
	file := open()
	defer closeFile(file)
	tasks, err := readTasks(file)
	if utils.CheckError(err, "Can't read the file!") {
		os.Exit(1)
	}
	task := model.NewTask(utils.IfThenElse(len(tasks) == 0, 1, tasks[len(tasks)-1].Id+1).(int), desc)
	tasks = append(tasks, *task)
	setOriginToZero(file)
	err = json.NewEncoder(file).Encode(tasks)
	utils.CheckError(err, "Can't write into the file!")
}

func DeleteTask(id int) {
	if !exists() {
		fmt.Println(statics.WarningStyle.Render("Task Not Found!"))
		return
	}
	file := open()
	defer closeFile(file)
	tasks, err := readTasks(file)
	if utils.CheckError(err, "Can't read the file!") {
		os.Exit(1)
	}
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			setOriginToZero(file)
			err = json.NewEncoder(file).Encode(tasks)
			if utils.CheckError(err, "Can't write into the file!") {
				os.Exit(1)
			}
			fmt.Println(statics.InfoStyle.Render(fmt.Sprintf("Task:%d deleted", id)))
			return
		}
	}
	fmt.Println(statics.WarningStyle.Render("Task Not Found!"))
}

func UpdateTask(id int, desc, status string) {
	if !exists() {
		fmt.Println(statics.WarningStyle.Render("Task Not Found!"))
		return
	}
	file := open()
	defer closeFile(file)
	tasks, err := readTasks(file)
	if utils.CheckError(err, "Can't read the file!") {
		os.Exit(1)
	}
	for i, task := range tasks {
		if task.Id == id {
			if desc != "" {
				tasks[i].Description = desc
			} else {
				tasks[i].Status = status
			}
			setOriginToZero(file)
			err = json.NewEncoder(file).Encode(tasks)
			if utils.CheckError(err, "Can't write into the file!") {
				os.Exit(1)
			}
			fmt.Println(statics.InfoStyle.Render(fmt.Sprintf("Task:%d updated", id)))
			return
		}
	}
	fmt.Println(statics.WarningStyle.Render("Task Not Found!"))
}
