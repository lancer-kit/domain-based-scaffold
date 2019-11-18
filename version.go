package main

import (
	"github.com/lancer-kit/domain-based-scaffold/info"
)

var (
	Version = "1.0.0-rc"
	Build   string
	Tag     string
)

func init() {
	info.App.Version = Version
	info.App.Build = Build
	info.App.Tag = Tag
}
