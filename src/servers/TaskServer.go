package servers

import "../tasks"

type ITaskServer interface {
	AddTask(task tasks.TaskToDo) bool
	GetTask() tasks.TaskToDo
}

type TaskServer struct {
	tasksChannel chan tasks.TaskToDo
}

func (taskServer TaskServer) AddTask(task tasks.TaskToDo) bool {
	taskServer.tasksChannel <- task
	return true
}

func (taskServer TaskServer) GetTask() tasks.TaskToDo {
	taskToGet := <-taskServer.tasksChannel
	return taskToGet
}

func CreateTaskServer(size int) TaskServer {
	taskServer := TaskServer{
		tasksChannel: make(chan tasks.TaskToDo, size),
	}

	return taskServer
}
