package chairman

import (
	"../../configuration"
	"../../servers"
	"../../tasks"
	"strconv"
)

type IChairmanFactory interface {
	CreateChairman() Chairman
}

type chairmanFactory struct {
	taskServerSize  int
	factoryName     string
	chairmanCounter *int
	taskFactory     tasks.ITaskFactory
	taskServer      servers.ITaskServer
}

func (factory chairmanFactory) CreateChairman() Chairman {
	chairman := Chairman{
		taskServer:  factory.taskServer,
		taskFactory: &factory.taskFactory,
		id: "CHAIRMAN_" + strconv.Itoa(*factory.chairmanCounter) +
			"::" + factory.factoryName,
	}

	*factory.chairmanCounter++

	return chairman
}

func CreateChairmanFactory(factoryName string) IChairmanFactory {
	config := configuration.GetConfiguration()
	taskFactory := tasks.CreateTaskFactory()
	taskFactoryPointer := &taskFactory

	taskServer := servers.CreateTaskServer(config.TaskServerSize)
	taskServerPointer := &taskServer

	chairmanFactory := chairmanFactory{
		taskFactory:     taskFactoryPointer,
		taskServer:      taskServerPointer,
		factoryName:     factoryName,
		chairmanCounter: new(int),
	}

	*chairmanFactory.chairmanCounter = 1

	return chairmanFactory
}

/*func (chairmanFactory ChairmanFactory) Create() company_members.ICompanyMember {
	createdChairman := company_members.CreateChairman(chairmanFactory.chairmanCounter)
	createdChairman.SelectTaskFactory(chairmanFactory.taskFactory)
	createdChairman.SelectTaskServer(chairmanFactory.taskServer)

	chairmanFactory.chairmanCounter++

	return createdChairman
}

func CreateCharimanFactory() ChairmanFactory {
	tasksServer := servers.CreateTaskServer(10);
	taskFactory := tasks.CreateTaskFactory()

	chairmanFactory := ChairmanFactory{
		chairmanCounter: 1,
		taskServer:      tasksServer,
		taskFactory:     taskFactory,
	}

	return chairmanFactory
}

func (chairman Chairman) SelectTaskServer(taskServer servers.TaskServer) {
	chairman.taskServer = taskServer
}

func (chairman Chairman) SelectTaskFactory(factory *tasks.TaskFactory) {
	chairman.taskFactory = factory
}

func (chairman Chairman) DoResponsibilities() {
	createdTask := (*chairman.taskFactory).CreateTask()
	chairman.taskServer.AddTask(createdTask)
}

func CreateChairman(id int) Chairman{
	chairman := Chairman{id: "CHAIRMAN_" + strconv.Itoa(id)}
	return chairman
}*/
