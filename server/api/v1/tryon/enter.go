package tryon

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	TryOnApi
}

var (
	tryOnService = service.ServiceGroupApp.TryOnServiceGroup.TryOnService
)
