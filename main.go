package main

import (
	"fmt"
	"github.com/asherplotnik/golang-rest/myRouter"
	"github.com/asherplotnik/golang-rest/repository"
)

func main() {
	fmt.Println("Starting API")
	repository.Init()
	myRouter.Init()
}
