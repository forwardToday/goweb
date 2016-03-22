package minus

import (
	"time"
)

type Config struct {
	HttpAddr     string
	HttpPort     int
	TemplatePath string
	RecoverPanic bool
	RunMode      int8 //0=prod，1=dev
	UseFcgi      bool
	ReadTimeout  time.Duration // maximum duration before timing out read of the request, 默认:5*time.Second(5秒超时)
	WriteTimeout time.Duration // maximum duration before timing out write of the response, 默认:0(不超时)
}

const (
	RunModeProd int8 = 0
	RunModeDev  int8 = 1
)
