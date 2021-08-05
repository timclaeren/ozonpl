package main

import "fmt"

type pist struct {
	On bool
	Ammo, Power int
}

func (c *pist) Shoot() bool{
	var n bool
	if c.Ammo > 0 && c.On == true{
		n = true
		c.Ammo --
	} else{n=false}
	return n
}
func (c *pist) RideBike() bool{
	var n bool
	if c.Power > 0 && c.On == true {
		n = true
		c.Power --
	} else {n = false}
	return n
}


func main() {
	c := pist{true, 1, 1}
	testStruct := &c
	fmt.Println(testStruct)
}




