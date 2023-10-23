package system

type RouterGroup struct {
	JwtRouter
	SysRouter
	BaseRouter
	InitRouter
	UserRouter
	OperationRecordRouter
}
