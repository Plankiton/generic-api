// Package main provides ...
package main

import (
	Sex "github.com/Plankiton/SexPistol"
	"github.com/plankiton/generic-api/pkg/env"
)

func main() {
	Sex.NewPistol().
		Add("/", func(Sex.Request) string {
			return "Hello, World!"
		}).
		Run(env.PORT)
}
