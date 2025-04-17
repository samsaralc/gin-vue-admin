package request

type TryOnRequest struct {
	ModelUrl    *string `json:"modelUrl" binding:"required"`
	ClothesUrl  *string `json:"clothesUrl" binding:"required"`
	ClothesType *string `json:"clothesType" binding:"required,oneof=Upper-body Lower-body Full-body"`
}
