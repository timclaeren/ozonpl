package main

import "fmt"

//import "fmt"

type pist struct {
	On bool
	Ammo, Power int
}

func (c *pist) Shoot() bool{
	var n bool
	if c.Ammo >= 1{
		n = true
		c.Ammo --
	} else {n=false}
	return n
}
func (c *pist) RideBike() bool{
	var n bool
	if c.Power >= 1{
		n = true
		c.Power --
	} else {n=false}
	return n
}
func (c *pist) retry() bool{
	var b bool
	if c.Power >=1 && c.Ammo >=1{
		b = true
	} else {b=false}
	return b
}
func main() {
	//c := pist{true, 5, 5}

	c := pist{true, 5, 5}
	//testStruct := *new(pist)
	fmt.Println(c)

}




