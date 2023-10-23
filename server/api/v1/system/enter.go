package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DBApi
	JwtApi
	BaseApi
	SystemApi
	OperationRecordApi
}

var (
	jwtService             = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService            = service.ServiceGroupApp.SystemServiceGroup.UserService
	initDBService          = service.ServiceGroupApp.SystemServiceGroup.InitDBService
	systemConfigService    = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)
