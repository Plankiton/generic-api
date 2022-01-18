// Package main provides ...
package main

import Sex "github.com/Plankiton/SexPistol"

func main() {
	Sex.NewPistol().
		Add("/", func(Sex.Request) string {
			return "Hello, World!"
		}).
		Run()
}
