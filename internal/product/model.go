package product

import "github.com/vmware/vending/internal/common"

type Product struct {
	common.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Cost uint `json:"cost"`
}
