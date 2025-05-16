package cart

import (
	"github.com/orcaman/concurrent-map/v2"
)

type Cart struct {
	Items cmap.ConcurrentMap[string, Item]
}

type Item struct {
	Id       string
	Name     string
	Price    float64
	Quantity int
}

func NewCart() Cart {
	return Cart{
		Items: cmap.New[Item](),
	}
}

func (c *Cart) TotalUniqueItems() int {
	return c.Items.Count()
}

func (c *Cart) TotalQuantity() int {
	sum := 0
	c.Items.IterCb(func(_ string, v Item) {
		sum += v.Quantity
	})
	return sum
}

func (c *Cart) TotalPrice() float64 {
	sum := 0.0
	c.Items.IterCb(func(_ string, v Item) {
		sum += float64(v.Quantity) * v.Price
	})
	return sum
}

func (c *Cart) AddItem(i Item) {
	c.Items.Upsert(i.Id, i, func(exist bool, oldValue Item, newValue Item) Item {
		if !exist {
			newValue.Quantity = 1
			return newValue
		}
		oldValue.Quantity++
		return oldValue
	})
}

func (c *Cart) RemoveItem(i Item) {
	oldItem, ok := c.Items.Get(i.Id)
	if ok {
		if oldItem.Quantity <= 1 {
			c.Items.Remove(i.Id)
		} else {
			oldItem.Quantity--
			c.Items.Set(i.Id, oldItem)
		}
	}
}
