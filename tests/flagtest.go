package main

import (
	"flag"
	"fmt"
)

var (
	name string
	age int
	sex bool
)

func init()  {
	flag.StringVar(&name,"name","kaishidongle","input your name")
	flag.IntVar(&age,"age",20,"input your age")
	flag.BoolVar(&sex,  "sex", false, "input true or false")
	//flag.BoolVar(&sex, "boolflag", false, "bool flag value")

}

func main() {
	flag.Parse()
	fmt.Println("name:",name)
	fmt.Println("age:",age)
	fmt.Println("sex:",sex)
}