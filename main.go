package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/locale"
	"flamingo.me/flamingo/v3/framework/config"
	"flamingo.me/flamingo/v3/framework/prefixrouter"
	"flamingo.me/helloworld/src/helloworld"
	"flamingo.me/pugtemplate"
)

func main() {
	flamingo.App([]dingo.Module{
		new(prefixrouter.Module),
		new(pugtemplate.Module),
		new(locale.Module),
		new(helloworld.Module),
	}, flamingo.ChildAreas(
		config.NewArea("de_de", nil),
	))
}
