//Package appcomands consists of comands for todo application
package appcomands

import (
	"bufio"
	"consoleToDo/filework"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

//Task needs for json
type Task struct {
	Content  string
	Complete bool
}

//ListTasks prints tasks to the console
func ListTasks() {
	file := filework.OpenTaskFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		output := scanner.Text()
		task := Task{}
		err := json.Unmarshal([]byte(output), &task)
		if err != nil {
			log.Fatal("Can't unmarshal information")
		}
		if !task.Complete {
			completeStat := "In process"
			fmt.Printf("[%d] %s - %s  /  %s\n", i, time.Now().Format("01-02-2006"), task.Content, completeStat)
			i++
		} else {
			completeStat := "Completed"
			fmt.Printf("[%d] %s - %s  /  %s\n", i, time.Now().Format("01-02-2006"), task.Content, completeStat)
			i++
		}
	}
}

//AddTask writes task in the file
func AddTask(task Task, filename string) {
	jsonMarchal, err := json.Marshal(task)
	if err != nil {
		log.Fatal("Can't marshal this json")
	}
	jsonMarchal = append(jsonMarchal, "\n"...)
	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if _, err = file.Write(jsonMarchal); err != nil {
		log.Fatal("Can't add a task")
	}
	defer file.Close()
}

//DeleteTask deletes task from list
func DeleteTask(idx int) {
	file := filework.OpenTaskFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	arrayOfTasks := []string{}
	for scanner.Scan() {
		output := scanner.Text()
		arrayOfTasks = append(arrayOfTasks, output)
	}
	if len(arrayOfTasks) == 1 {
		os.Remove("tasks.json")
		return
	}
	for j := range arrayOfTasks {
		if j+1 != idx {
			currentTask := Task{}
			err := json.Unmarshal([]byte(arrayOfTasks[j]), &currentTask)
			if err != nil {
				log.Fatal("Can't unmarshal information")
			}
			AddTask(currentTask, ".tempfile")
		}
	}
	os.Rename(".tempfile", "tasks.json")
	os.Remove(".tempfile")
}

//CompleteTask closes the task
func CompleteTask(idx int) {
	file := filework.OpenTaskFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		output := scanner.Text()
		currentTask := Task{}
		err := json.Unmarshal([]byte(output), &currentTask)
		if err != nil {
			log.Fatal("Can't unmarshal information")
		}
		if !currentTask.Complete {
			if idx == i {
				currentTask.Complete = true
			}
			i++
		}
		AddTask(currentTask, ".tempfile")
	}
	os.Rename(".tempfile", "tasks.json")
	os.Remove(".tempfile")
}
