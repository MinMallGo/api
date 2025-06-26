package forms

type OrderCreate struct {
	Name    string `json:"name" binding:"required"`
	Mobile  string `json:"mobile" binding:"required,mobile"`
	Address string `json:"address" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type OrderList struct {
	Page int `json:"page"`
	Size int `json:"size"`
}
