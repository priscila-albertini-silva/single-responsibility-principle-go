package main

import "fmt"

type Task struct {
	description string
	done        bool
}

type TaskList struct {
	tasks []Task
}

func (tl *TaskList) addTask(task Task) {
	tl.tasks = append(tl.tasks, task)
}

func (tl *TaskList) markTaskAsDone(index int) {
	if index >= 0 && index < len(tl.tasks) {
		tl.tasks[index].done = true
	}
}

type TaskPrinter struct {
	taskList TaskList
}

func (tp *TaskPrinter) printTasks() {
	for _, task := range tp.taskList.tasks {
		fmt.Println(task.description)
	}
}

func main() {
	taskList := TaskList{}
	taskList.addTask(Task{description: "Buy groceries", done: false})
	taskList.addTask(Task{description: "Clean the house", done: false})

	taskPrinter := TaskPrinter{taskList: taskList}

	fmt.Println("Tasks:")
	taskPrinter.printTasks()

	taskList.markTaskAsDone(0)

	fmt.Println("\nUpdated tasks:")
	taskPrinter.printTasks()
}
