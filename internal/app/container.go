package app

import (
	"github.com/a1ta1r/distributed-lab-server/internal/controller"
	"github.com/a1ta1r/distributed-lab-server/internal/service"
	"github.com/a1ta1r/distributed-lab-server/internal/storage"
	"github.com/jinzhu/gorm"
)

var Connection *gorm.DB = GetConnection()
var UserStorage storage.UserStorage = storage.NewUserStorage(*Connection)
var UserService service.UserService = service.NewUserService(UserStorage)
var UserController controller.UserController = controller.NewUserController(UserService)

