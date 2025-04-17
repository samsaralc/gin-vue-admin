package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/tryon"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	tryOnGroup      tryon.ServiceGroup
}
