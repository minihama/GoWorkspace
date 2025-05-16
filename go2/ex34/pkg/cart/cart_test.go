package cart_test

import (
	. "ex34/pkg/cart"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("쇼핑카트", func() {
	var itemA, itemB Item
	BeforeEach(func() {
		itemA = Item{Id: "i9Z3", Name: "item1", Price: 3, Quantity: 0}
		itemB = Item{Id: "i9Z4", Name: "item2", Price: 4, Quantity: 0}
	})

	Context("초기화", func() {
		var cart Cart
		BeforeEach(func() {
			cart = NewCart()
		})
		It("0개의 아이템을 가지고 있다", func() {
			Expect(cart.TotalUniqueItems()).Should(BeZero())
		})
		It("0개의 유닛을 가지고 있다", func() {
			Expect(cart.TotalQuantity()).Should(BeZero())
		})
		It("총 가격은 0.00이다", func() {
			Expect(cart.TotalPrice()).Should(BeZero())
		})
	})

	Context("item1을 1번 추가", func() {
		var cart Cart
		BeforeEach(func() {
			cart = NewCart()
			cart.AddItem(itemA)
		})
		It("1개의 아이템을 가지고 있다", func() {
			Expect(cart.TotalUniqueItems()).Should(Equal(1))
		})
		It("1개의 유닛을 가지고 있다", func() {
			Expect(cart.TotalQuantity()).Should(Equal(1))
		})
		It("총 가격은 3이다", func() {
			Expect(cart.TotalPrice()).Should(Equal(3.0))
		})
	})

	Context("item1을 2번 추가", func() {
		var cart Cart
		BeforeEach(func() {
			cart = NewCart()
			cart.AddItem(itemA)
			cart.AddItem(itemA)
		})
		It("1개의 아이템을 가지고 있다", func() {
			Expect(cart.TotalUniqueItems()).Should(Equal(1))
		})
		It("2개의 유닛을 가지고 있다", func() {
			Expect(cart.TotalQuantity()).Should(Equal(2))
		})
		It("총 가격은 3이다", func() {
			Expect(cart.TotalPrice()).Should(Equal(6.0))
		})
	})

	Context("item1을 1번 추가, item2를 1번 추가, item1을 삭제", func() {
		var cart Cart
		BeforeEach(func() {
			cart = NewCart()
			cart.AddItem(itemA)
			cart.AddItem(itemB)
			cart.RemoveItem(itemA)
		})
		It("1개의 아이템을 가지고 있다", func() {
			Expect(cart.TotalUniqueItems()).Should(Equal(1))
		})
		It("1개의 유닛을 가지고 있다", func() {
			Expect(cart.TotalQuantity()).Should(Equal(1))
		})
		It("총 가격은 4이다", func() {
			Expect(cart.TotalPrice()).Should(Equal(4.0))
		})
	})
})
