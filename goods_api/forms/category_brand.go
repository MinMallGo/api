package forms

type CreateCategoryBrand struct {
	CategoryId int `json:"category_id" binding:"required"`
	BrandId    int `json:"brand_id" binding:"required"`
}
