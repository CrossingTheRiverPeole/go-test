package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(60)
	fmt.Println(intn)
}
