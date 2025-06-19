package forms

type GoodsCreate struct {
	// CategoryID 商品类别ID，必须提供且大于0
	CategoryID int32 `json:"category_id" form:"category_id" binding:"required,gt=0"`

	// BrandID 品牌ID，必须提供且大于0
	BrandID int32 `json:"brand_id" form:"brand_id" binding:"required,gt=0"`

	// OnSale 是否在售，布尔类型，默认为false
	// binding:"omitempty" 表示允许为空，如果为空则通常使用默认值
	OnSale bool `json:"on_sale" form:"on_sale" binding:"omitempty"`

	// ShipFree 是否包邮，布尔类型，默认为false
	ShipFree bool `json:"ship_free" form:"ship_free" binding:"omitempty"`

	// IsNew 是否新品，布尔类型，默认为false
	IsNew bool `json:"is_new" form:"is_new" binding:"omitempty"`

	// Stock 库存数量，必须提供且大于等于0
	Stock int32 `json:"stock" form:"stock" binding:"required,gte=0"`

	// Name 商品名称，必须提供且长度在2到100字符之间
	Name string `json:"name" form:"name" binding:"required,min=2,max=100"`

	// GoodsSn 商品编号，必须提供且长度在2到100字符之间
	// binding:"required" 表示必填
	GoodsSn string `json:"goods_sn" form:"goods_sn" binding:"required,min=2,max=100"`

	// ClickNum 点击量，默认为0，允许客户端提供，但通常是后端统计
	ClickNum int32 `json:"click_num,omitempty" form:"click_num" binding:"omitempty,gte=0"`

	// SoldNum 销量，默认为0，允许客户端提供，但通常是后端统计
	SoldNum int32 `json:"sold_num,omitempty" form:"sold_num" binding:"omitempty,gte=0"`

	// FavNum 收藏量，默认为0，允许客户端提供，但通常是后端统计
	FavNum int32 `json:"fav_num,omitempty" form:"fav_num" binding:"omitempty,gte=0"`

	// MarketPrice 市场价格，必须提供且大于等于0
	MarketPrice float32 `json:"market_price" form:"market_price" binding:"required,gte=0"`

	// ShopPrice 销售价格，必须提供且大于等于0，且小于等于市场价格
	ShopPrice float32 `json:"shop_price" form:"shop_price" binding:"required,gte=0,ltefield=MarketPrice"` // ltefield 校验 shop_price <= MarketPrice

	// GoodsBrief 商品简述，必须提供且长度在5到100字符之间
	GoodsBrief string `json:"goods_brief" form:"goods_brief" binding:"required,min=5,max=100"`

	// ImageUrl 商品图片URL列表，可以为空，但如果提供则至少有一个URL
	// binding:"omitempty,min=1,dive,url" 表示可以为空，如果非空则至少一个元素，并且每个元素都是有效的URL
	ImageUrl []string `json:"image_url,omitempty" form:"image_url" binding:"omitempty,min=1,dive,url"`

	// Description 商品描述，可以为空，但如果提供则至少有一个元素
	// 这里假设 description 也是多行或多段文本的列表，如果是一个长字符串则改为 string
	Description []string `json:"description,omitempty" form:"description" binding:"omitempty"`

	// GoodsFrontImage 商品封面图URL，必须提供且为有效的URL
	GoodsFrontImage string `json:"goods_front_image" form:"goods_front_image" binding:"required,url"`
}
