package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/niko-greb/demosite/pkg"

	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().Unix())
	pkg.HandleFunc()
}
