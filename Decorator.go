package main

import "fmt"

type Order interface {
	GetPrice() int
	GetDiscription() string
}

type BaseOrder struct{}

func(b *BaseOrder) GetPrice() int{
 return 100
}
func(b *BaseOrder) GetDiscription() string{
	return "Базовый заказ"
   }

type OrderDecorator struct{
	order Order
}

func(od *OrderDecorator) GetPrice() int {
	return od.order.GetPrice()
}

func(od *OrderDecorator) GetDiscription() string{
	return od.order.GetDiscription()
}

type DeliveryDecorator struct{
	OrderDecorator
}

func(dd *DeliveryDecorator) GetPrice() int{
	return dd.order.GetPrice()+ 50
}
func(dd *DeliveryDecorator) GetDiscription() string{
	return dd.order.GetDiscription()+ ", c доставкой"
}

type GiftWrapDecorator struct{
	OrderDecorator
}

func(gwd *GiftWrapDecorator) GetPrice() int{
	return gwd.order.GetPrice()+ 25
}
func(gwd *GiftWrapDecorator) GetDiscription() string{
	return gwd.order.GetDiscription()+ ", c подарочной упоковкой"
}

func main(){
	baseOrder := &BaseOrder{}
	orderWithGiftWrap := &GiftWrapDecorator{OrderDecorator{order: baseOrder}}
	orderWithDelivery := &DeliveryDecorator{OrderDecorator{order: baseOrder}}
	orderWithBothOptions := &DeliveryDecorator{OrderDecorator{order: &GiftWrapDecorator{OrderDecorator{order: baseOrder}}}}
	fmt.Println("Базовый заказ:",baseOrder.GetDiscription(),
	baseOrder.GetPrice())
	
	

	fmt.Println("\nЗаказ с подарочной упаковкой:"	,orderWithGiftWrap.GetDiscription(),"; Цена:",
	orderWithGiftWrap.GetPrice())


	fmt.Println("\nЗаказ с доставкой:"	,orderWithDelivery.GetDiscription(), "; Цена:",
	orderWithDelivery.GetPrice())

	fmt.Println("\nЗаказ с обеими опциями:",	orderWithBothOptions.GetDiscription(),"; Цена:",
	orderWithBothOptions.GetPrice())


}



