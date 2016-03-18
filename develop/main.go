package main

import (
	. "github.com/sczhaoyu/panda"
	_ "github.com/sczhaoyu/panda/develop/controller"
)

func main() {
	LocalAddress = ":9055"
	Run()
}
