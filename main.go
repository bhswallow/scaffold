package main

import (
	"fmt"
	"scaffold/log"
	"scaffold/swallow"
)

func main()  {

	var modules []string = []string{"base","mysql","redis"}
	err := swallow.InitModule("./config/dev/",modules)
	if err != nil{
		log.Fatal("can not found config dir , "+ err.Error())
	}

	defer swallow.Destroy()

	fmt.Println("main end")
}



