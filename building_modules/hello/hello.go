package main

import (
	"fmt"

	"example.com/greetings"
	
)

func main(){
	message := greetings.Hello("OmSharma");
	fmt.Println(message);
}