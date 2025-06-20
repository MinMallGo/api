package forms

type BrandCreate struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo" binding:"required,url"`
}
