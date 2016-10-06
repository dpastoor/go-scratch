package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kataras/iris"
)

func main() {
	api := iris.Party("/api")
	v1 := api.Party("/v1")
	v1.Get("/flu", func(ctx *iris.Context) {
		ctx.JSON(http.StatusOK, allFromOrg("all"))
	})
	iris.Listen(":8080")
}

// Person is a simple type for testing sqlite
type Person struct {
	Name         string
	Organization string
}

func allFromOrg(org string) []Person {

	fmt.Println("fetching data for org: ", org)
	start := time.Now()
	time.Sleep(1 * time.Second)
	fmt.Println("got data in: ", time.Since(start))
	return []Person{
		Person{"P1", "A"},
		Person{"P2", "A"},
		Person{"P1", "B"},
		Person{"P2", "B"},
	}
}
