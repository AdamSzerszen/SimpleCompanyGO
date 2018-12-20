package servers

import "../tasks"

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
