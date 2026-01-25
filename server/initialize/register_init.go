package initialize

import (
	_ "gva/source/example"
	_ "gva/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
