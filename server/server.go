package server

import (
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"strconv"
	"sync"
	"weather-conditions/props"
	"weather-conditions/server/api"
)

type Server struct {
	host        string
	port        int
	contextRoot string
	doOnce      sync.Once
}

func NewWeatherServer(properties *viper.Viper, prop *props.Properties) *Server {
	if err := properties.Unmarshal(&prop, func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.StringToTimeDurationHookFunc()
	}); err != nil {
		// Handle the error
	}
	server := new(Server)
	server.host = prop.Server.Host
	server.port = prop.Server.Port
	server.contextRoot = prop.Server.ContextRoot
	return server
}

func (s *Server) ConfigureAPI() {
	s.doOnce.Do(func() {
		api.ConfigureApi(s.contextRoot, strconv.Itoa(s.port))
	})
}
