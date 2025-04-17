package tryon

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	tryonReq "github.com/flipped-aurora/gin-vue-admin/server/model/tryon/request"
	aiart "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/aiart/v20221229"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"go.uber.org/zap"
)

type TryOnService struct{}

func (t *TryOnService) ChangeClothes(req *tryonReq.TryOnRequest) (*string, error) {
	// 创建腾讯云身份凭证
	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)

	// 客户端配置
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "aiart.tencentcloudapi.com"

	// 创建 client
	client, err := aiart.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		global.GVA_LOG.Error("创建腾讯云客户端失败", zap.Error(err))
		return nil, err
	}

	// 创建请求对象
	request := aiart.NewChangeClothesRequest()
	request.ModelUrl = req.ModelUrl
	request.ClothesUrl = req.ClothesUrl
	request.ClothesType = req.ClothesType

	// 发送请求
	response, err := client.ChangeClothes(request)
	if sdkErr, ok := err.(*errors.TencentCloudSDKError); ok {
		global.GVA_LOG.Error("腾讯云SDK返回错误", zap.Error(sdkErr))
		return nil, err
	}
	if err != nil {
		global.GVA_LOG.Error("请求失败", zap.Error(err))
		return nil, err
	}

	return response.Response.ResultImage, nil
}
