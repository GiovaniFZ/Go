package main

import (
	"github.com/GiovaniFZ/Go/errorTest"
	"github.com/GiovaniFZ/Go/maps"
	"github.com/GiovaniFZ/Go/meet"
	"github.com/GiovaniFZ/Go/texts"
)

func main() {
	meet.SayHello()
	meet.Say("Hello")
	texts.StringsGo()
	errorTest.ErrorTest()
	maps.CheckPlayers()
}