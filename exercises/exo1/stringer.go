package main

import (
	"fmt"
)

//création du type rectangle
type Rectangle struct {
	Length string
	Width  string
}

//création du type cercle
type Circle struct {
	Radius string
}

//interface pouvant être implémentée, donnant des noms de méthodes et pouvant être modifié selon la situation
type Shape interface {
}

// fonction permettant de printer le infos du rectangle
func (rect Rectangle) String() string {
	return fmt.Sprintf("Square of width %s and length %s", rect.Width, rect.Length)
}

// fonction permettant de printer le infos du cercle
func (circ Circle) String() string {
	return fmt.Sprintf("circle of radius %s", circ.Radius)
}

//fonction principale qui tourne
func main() {

	r := &Rectangle{"4", "6"} // permet de créer une instance de rectangle
	c := &Circle{"11"}

	s := []Shape{r, c}

	for _, s := range s {
		fmt.Println(s)
	}

}
