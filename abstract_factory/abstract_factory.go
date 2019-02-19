package abstract_factory

import "fmt"

type Product1 interface {
	Print1()
}

type Product2 interface {
	Print2()
}

type AbstractFactory interface {
	CreateProduct1() Product1
	CreateProduct2() Product2
}

type ConcreteProductA1 struct {
}

func (concreteProductA1 *ConcreteProductA1) Print1() {
	fmt.Println("ConcreteProductA1")
}

type ConcreteProductA2 struct {
}

func (concreteProductA1 *ConcreteProductA2) Print2() {
	fmt.Println("ConcreteProductA2")
}

type ConcreteFactoryA struct {
}

func (factory *ConcreteFactoryA) CreateProduct1() Product1 {
	return &ConcreteProductA1{}
}

func (factory *ConcreteFactoryA) CreateProduct2() Product2 {
	return &ConcreteProductA2{}
}

type ConcreteProductB1 struct {
}

func (concreteProductB1 *ConcreteProductB1) Print1() {
	fmt.Println("ConcreteProductB1")
}

type ConcreteProductB2 struct {
}

func (concreteProductB1 *ConcreteProductB2) Print2() {
	fmt.Println("ConcreteProductB2")
}

type ConcreteFactoryB struct {
}

func (factory *ConcreteFactoryB) CreateProduct1() Product1 {
	return &ConcreteProductB1{}
}

func (factory *ConcreteFactoryB) CreateProduct2() Product2 {
	return &ConcreteProductB2{}
}

func NewConcreteFactory(factoryType string) AbstractFactory {
	if factoryType == "A" {
		return &ConcreteFactoryA{}
	}
	if factoryType == "B" {
		return &ConcreteFactoryB{}
	}
	return nil
}

func AbstractFactoryExecute(factoryType string) {
	factory := NewConcreteFactory(factoryType)
	product1 := factory.CreateProduct1()
	product2 := factory.CreateProduct2()

	product1.Print1()
	product2.Print2()
}
