package main

import (
	"fmt"
)

type IAttack interface {
	Attack()
}

type HumanLowLevel struct {
	name string
	level  int
}

func (a *HumanLowLevel) Attack(){
	fmt.Println("woshi:", a.name,",dengjiwei:",a.level)
}

type HumanHighLevel struct {
	name string
	level  int
}

func (a *HumanHighLevel) Attack(){
	fmt.Println("woshi:", a.name,",dengjiwei:",a.level)
}

//定义一个多态的通用接口,传入不同的对象，调用同样的方法，就是实现了多态
func DoAttack(a IAttack){
	a.Attack()
}


func main(){
	lowLevel:= HumanLowLevel{
		name:"hello",
		level: 1,
	}

	highLevel:= HumanLowLevel{
		name:"hello1",
		level: 10,
	}
	lowLevel.Attack()
	var player IAttack
	player = &lowLevel
	player.Attack()

	player = &highLevel
	player.Attack()
	//多态
	DoAttack(&lowLevel)
	DoAttack(&highLevel)

}