package forms

type UserListForm struct {
	Page int `json:"page" form:"page" binding:"omitempty,min=0"`
	Size int `json:"size" form:"size" binding:"omitempty,min=0"`
}

type PasswordLoginForm struct {
	Mobile    string `json:"mobile" binding:"required,len=11,mobile"`
	Password  string `json:"password" binding:"required,min=3,max=20"`
	Captcha   string `json:"captcha" binding:"required,min=1,max=20"`
	CaptchaID string `json:"captcha_id" binding:"required,min=1,max=20"`
}

type UserCreateForm struct {
	Mobile   string `json:"mobile" binding:"required,len=11,mobile"`
	Nickname string `json:"nickname" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=3,max=20"`
}
