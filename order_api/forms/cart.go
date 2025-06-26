package forms

type CartCreate struct {
	GoodsId int `json:"goods_id" binding:"required,gt=0"`
	Nums    int `json:"nums" binding:"required,gt=0"`
}

type CartSelect struct {
	GoodsIds []int32 `json:"goods_ids" binding:"required"`
}
