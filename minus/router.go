package minus

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type controllerInfo struct {
	// regex          *regexp.Regexp
	// params         map[int]string
	controllerType reflect.Type
}

type ControllerRegistor struct {
	routers []*controllerInfo
	App     *App
}

func NewControllerRegistor() *ControllerRegistor {
	return &ControllerRegistor{}
}

func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {
	t := reflect.Indirect(reflect.ValueOf(c)).Type()
	route := &controllerInfo{}
	route.controllerType = t

	p.routers = append(p.routers, route)
}

func (app *App) SetStaticPath(url string, path string) *App {
	Mapp.StaticDirs[url] = path
	return app
}

func (p *ControllerRegistor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("ServeHTTP:v%", p.routers)
	fmt.Printf("ServeHTTP:s%", r.Method)
	var started bool
	for prefix, staticDir := range Mapp.StaticDirs {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := staticDir + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			started = true
			return
		}
	}
	// requestPath := r.URL.Path

	//find a matching Route
	for _, route := range p.routers {
		//Invoke the request handler
		vc := reflect.New(route.controllerType)
		init := vc.MethodByName("Init")
		in := make([]reflect.Value, 2)
		ct := &Context{ResponseWriter: w, Request: r}
		in[0] = reflect.ValueOf(ct)
		in[1] = reflect.ValueOf(route.controllerType.Name())
		init.Call(in)
		in = make([]reflect.Value, 0)
		method := vc.MethodByName("Prepare")
		method.Call(in)
		if r.Method == "GET" {
			fmt.Printf("GET ==============")
			method = vc.MethodByName("Get")
			method.Call(in)
		} else if r.Method == "POST" {
			method = vc.MethodByName("Post")
			method.Call(in)
		} else if r.Method == "HEAD" {
			method = vc.MethodByName("Head")
			method.Call(in)
		} else if r.Method == "DELETE" {
			method = vc.MethodByName("Delete")
			method.Call(in)
		} else if r.Method == "PUT" {
			method = vc.MethodByName("Put")
			method.Call(in)
		} else if r.Method == "PATCH" {
			method = vc.MethodByName("Patch")
			method.Call(in)
		} else if r.Method == "OPTIONS" {
			method = vc.MethodByName("Options")
			method.Call(in)
		}
		// if AutoRender {
		// 	method = vc.MethodByName("Render")
		// 	method.Call(in)
		// }
		method = vc.MethodByName("Finish")
		method.Call(in)
		started = true
		break
	}

	//if no matches to url, throw a not found exception
	if started == false {
		http.NotFound(w, r)
	}
}
