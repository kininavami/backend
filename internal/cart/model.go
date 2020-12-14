package cart

import "github.com/vmware/vending/internal/product"

type Cart struct {
	UserID uint
	Items []product.Item
}
