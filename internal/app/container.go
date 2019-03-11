package app

import (
	"github.com/a1ta1r/distributed-lab-server/internal/controller"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
)

var Connection = GetConnection()

// User
var UserStorage = storage.NewUserStorage(*Connection)
var UserService = service.NewUserService(UserStorage)
var UserController = controller.NewUserController(UserService)

// Team
var TeamStorage = storage.NewTeamStorage(*Connection)
var TeamService = service.NewTeamService(TeamStorage)
var TeamController = controller.NewTeamController(TeamService)

// Status
var StatusStorage = storage.NewStatusStorage(*Connection)
var StatusService = service.NewStatusService(StatusStorage)
var StatusController = controller.NewStatusController(StatusService)
