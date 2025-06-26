package structure

type CartListResp struct {
	UserID   int    `json:"user_id"`
	GoodsID  int    `json:"goods_id"`
	Nums     int    `json:"num"`
	Checked  bool   `json:"checked"`
	GoodsImg string `json:"goods_img"`
}
