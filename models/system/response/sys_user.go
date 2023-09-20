package Response

import "time"

type User struct {
	ID       uint      `json:"id"`
	UserName string    `json:"userName"`
	RoleKey  uint      `json:"roleKey"`
	CreatAt  time.Time `json:"creatAt"`
}
