package router

import (
	"github.com/VINDA-98/Tenacity-Solutions/server/router/example"
	"github.com/VINDA-98/Tenacity-Solutions/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
