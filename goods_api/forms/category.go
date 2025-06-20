package forms

type CategoryCreate struct {
	Name  string `json:"name" binding:"required"`
	Pid   int    `json:"parent_category_id" binding:"omitempty"`
	Level int    `json:"level" binding:"required,oneof=1 2 3"`
	IsTab bool   `json:"is_tab" binding:"omitempty"`
}
