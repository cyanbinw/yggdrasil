package router

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

const (
	Get    RequestEnum = 1
	Post               = 2
	Put                = 3
	Delete             = 4
)

type RequestEnum int

type Path struct {
	Action  gin.HandlerFunc
	Request RequestEnum
	Route   string
}

type Route struct {
	Controller string
	Work       []Path
}

var Routes []Route

func (t Route) GetWork() []Path {
	return t.Work
}

func (t *Route) GetController() {
	t.Controller = "Bill"
}

func Register(controller interface{}, controllerName string) bool {
	d := reflect.ValueOf(controller)
	route := Route{}
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + controllerName + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			route.Work = append(route.Work, data)
		}
	}
	Routes = append(Routes, route)
	return true
}

func (t *Route) SetWork() {
	if t.Controller == "" {
		t.GetController()
	}
	d := reflect.ValueOf(t).Elem()
	for i := 0; i < d.NumMethod(); i++ {
		k := d.Type()
		data := Path{}
		if !strings.Contains(k.Method(i).Name, "RE") {
			continue
		}
		name := k.Method(i)
		v := d.Method(i)
		z := v.Interface()
		strList := strings.Split(name.Name, "RE")
		if len(strList) > 0 {
			path := "/" + t.Controller + "/" + strList[0]
			data.Action = z.(func(ctx *gin.Context))
			data.Request = GetEnumValue(strList[1])
			data.Route = path
			t.Work = append(t.Work, data)
		}
	}
}

func GetEnumValue(str string) RequestEnum {
	switch str {
	case "Get":
		return Get
	case "Post":
		return Post
	case "Put":
		return Put
	case "Delete":
		return Delete
	default:
		return 0
	}
}

func LoadRouter(item []Path, c *gin.RouterGroup) {
	for _, i := range item {
		if i.Request == Get {
			c.GET(i.Route, i.Action)
		} else if i.Request == Post {
			c.POST(i.Route, i.Action)
		} else if i.Request == Put {
			c.PUT(i.Route, i.Action)
		} else if i.Request == Delete {
			c.DELETE(i.Route, i.Action)
		} else {

		}
	}
}
