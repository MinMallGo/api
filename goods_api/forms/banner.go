package forms

type BannerCreate struct {
	Image string `json:"image" binding:"required,url"`
	Url   string `json:"url" binding:"required,url"`
	Index int    `json:"index" binding:"omitempty,min=0"`
}
