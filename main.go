package main

import (
	"consoleToDo/appcomands"
	"log"
	"os"
	"strconv"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp() // Create new stage of the app
	app.Name = "To-Do List"
	app.Usage = "For adding and tracking tasks" // Description
	app.Commands = []*cli.Command{              // List of the comands
		{
			Name:  "add", // First comand
			Usage: "Add a task",
			Action: func(c *cli.Context) error {
				task := appcomands.Task{Content: c.Args().First(), Complete: false}
				appcomands.AddTask(task, "tasks.json")
				return nil
			},
		},
		{
			Name:  "done", // Second comand
			Usage: "Complete a task",
			Action: func(c *cli.Context) error {
				idx, err := strconv.Atoi(c.Args().First())
				if err != nil {
					log.Fatal("Can't read arg or convert to int")
				}
				appcomands.CompleteTask(idx)
				return nil
			},
		},
		{
			Name:  "ls", // Third comand
			Usage: "Print all uncompleted tasks in list",
			Action: func(c *cli.Context) error {
				appcomands.ListTasks()
				return nil
			},
		},
		{
			Name:  "del", // Fourth comand
			Usage: "Delete specified task",
			Action: func(c *cli.Context) error {
				idx, err := strconv.Atoi(c.Args().First())
				if err != nil {
					log.Fatal("Can't read arg or convert to int")
				}
				appcomands.DeleteTask(idx)
				return nil
			},
		},
	}
	app.Run(os.Args) // Start this stage of the app
}
