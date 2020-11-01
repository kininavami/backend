package user

import (
	"github.com/vmware/vending/internal/common"
)

type User struct {
	common.Model
	Name string `json:"name"`
	Username string `json:"username"`
	Password string	`json:"password"`
	Role string `json:"role"`
	Address string `json:"address"`
	Email string	`json:"email"`
}
