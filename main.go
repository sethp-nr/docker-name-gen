package main

import (
	"fmt"
	"time"
	"math/rand"
	"strings"

	"github.com/docker/docker/pkg/namesgenerator"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(strings.ReplaceAll(namesgenerator.GetRandomName(0), "_", "-"))
}


