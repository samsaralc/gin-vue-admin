package tryon

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tryon/request"
	tryres "github.com/flipped-aurora/gin-vue-admin/server/model/tryon/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TryOnApi struct{}

// UploadFile
// @Tags      tryon
// @Summary   换装
// @accept    multipart/form-data
// @Produce   application/json
// @Param     modelUrl  String  modelUrl
// @Success   200   {object}  response.Response{data=exampleRes.ExaFileResponse,msg=string}  "上传文件示例,返回包括文件详情"
// @Router    /fileUploadAndDownload/upload [post]
func (t *TryOnApi) submit(c *gin.Context) {
	req := new(request.TryOnRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		global.GVA_LOG.Error("参数绑定失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	r, err := tryOnService.ChangeClothes(req)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("换装失败，原因：%s", err.Error()), c)
		return
	}
	response.OkWithData(tryres.TryOnResponse{ResultImage: r}, c)
}
