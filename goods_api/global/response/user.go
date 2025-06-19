package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	tmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02 15:04:05"))
	return []byte(tmp), nil
}

type UserResponse struct {
	Id       uint64   `json:"id,omitempty"`
	Mobile   string   `json:"Mobile,omitempty"`
	NickName string   `json:"nickName,omitempty"`
	Birthday JsonTime `json:"birthday,omitempty"`
	Gender   string   `json:"gender,omitempty"`
}
