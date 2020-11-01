package cart

import "github.com/vmware/vending/internal/item"

type Cart struct {
	UserID uint
	Items []item.Item
}
