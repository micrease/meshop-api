package entity

type Product struct {
	//ID新增时可不传或为0，更新时必传
	ProdId   int32  `form:"prodId" json:"prodId"`
	ProdName string `form:"prodName" json:"prodName" binding:"required,gte=4,lte=32" tips:"prodName参数不合法"`
}
