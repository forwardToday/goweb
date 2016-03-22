package minus

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/fcgi"
	"os"
	"time"
)

const (
	Version = "1.0.0"
)

var (
	Mapp    *App
	AppPath string
)

func init() {
	Mapp = NewApp(nil)
	AppPath, _ = os.Getwd()
}

type App struct {
	Handlers   *ControllerRegistor
	config     *Config
	StaticDirs map[string]string
	// TemplateRegister *TemplaterRegister
}

func NewApp(config *Config) *App {
	cr := NewControllerRegistor()
	app := &App{
		Handlers:   cr,
		config:     config,
		StaticDirs: make(map[string]string),
		// TemplateRegistor: NewTemplateRegistor(),
	}
	cr.App = app
	return app
}

func (app *App) Run() {
	if app.config.HttpAddr == "" {
		app.config.HttpAddr = "172.0.0.1"
	}
	addr := fmt.Sprintf("%s:%d", app.config.HttpAddr, app.config.HttpPort)
	var err error

	// err = httpListenAndServe(addr, app.Handlers, app.config.ReadTimeout, app.config.WriteTimeout)

	for {
		if app.config.UseFcgi {
			l, e := net.Listen("tcp", addr)
			if e != nil {
				log.Print("Listen: ", e)
			}
			//log.Print("UseFcgi, fcgi.Serve")
			err = fcgi.Serve(l, app.Handlers)
		} else {
			//log.Print("http.ListenAndServe")
			//err = http.ListenAndServe(addr, app.Handlers)
			err = httpListenAndServe(addr, app.Handlers, app.config.ReadTimeout, app.config.WriteTimeout)
		}
		if err != nil {
			log.Print("ListenAndServe: ", err)
			//panic(err)
		}
		time.Sleep(time.Second * 2)
	}
}

func httpListenAndServe(addr string, handler http.Handler, readTimeout time.Duration, writeTimeout time.Duration) error {
	if readTimeout == 0 {
		readTimeout = 5 * time.Second
	}

	server := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}

	return server.ListenAndServe()
}

func Run(config *Config) {
	Mapp.config = config
	Mapp.Run()
}
