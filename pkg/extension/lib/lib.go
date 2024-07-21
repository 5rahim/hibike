package lib

import "reflect"

// Symbols variable stores the map of lib symbols per package.
var Symbols = map[string]map[string]reflect.Value{}

func init() {
	Symbols["github.com/5rahim/hibike/pkg/extension/lib"] = map[string]reflect.Value{
		"Symbols": reflect.ValueOf(Symbols),
	}
}
