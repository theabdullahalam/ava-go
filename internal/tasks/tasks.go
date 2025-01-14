package tasks

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theabdullahalam/ava-go/internal/utils"
)

type Task struct {
	Name string
	Description string
	Target string
	Args []string
	Type string
}

func (task Task) Run(args []string) string {
	if task.Name == "" {
		return "No Task detected"
	}

	response := "Could not run task " + task.Name
	if task.Type == "internal" {
		response = RunInternalTask(task.Name, args)
	}

	if task.Type == "script" {
		fmt.Printf("\nRunning script %s on node %s\n", task.Name, task.Target)
	}

	return response
}

func GetTaskListString() string {
	ava_folder := utils.GetAvaFolder()
	tasks_file_path := filepath.Join(ava_folder, "tasklist.json")

	data, err := os.ReadFile(tasks_file_path)
	if err != nil {
		fmt.Println(err)
	}

	return string(data)
}

func GetTaskList() []Task {
	task_list_string := GetTaskListString()

	dtaa := []byte(task_list_string)

	var taskList []Task
	err := json.Unmarshal(dtaa, &taskList)
	if err != nil {
		fmt.Println(err)
	}
	return taskList
}

func GetTask(name string) Task {
	task_list := GetTaskList()
	for _, task := range task_list {
		if task.Name == name {
			return task
		}
	}
	return Task{}
}