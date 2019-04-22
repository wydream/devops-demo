package ctrls

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shanexu/logn"
	"sync"
)

var log = logn.GetLogger()

type Route = func(r gin.IRouter)

type RouteModule struct {
	Name  string
	Route Route
}

var routes []RouteModule

func Register(name string, route Route) {
	for i := range routes {
		if name == routes[i].Name {
			panic(fmt.Sprintf("duplicated route module %q", name))
		}
	}
	routes = append(routes, RouteModule{name, route})
}

func initRoutes(r gin.IRouter) {
	log.Info("init routes...")
	for _, rm := range routes {
		log.Infof("init route: %q", rm.Name)
		rm.Route(r)
	}
}

var once sync.Once

func Init(r gin.IRouter) {
	once.Do(func() {
		initRoutes(r)
	})
}
