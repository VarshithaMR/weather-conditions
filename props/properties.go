package props

import (
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"sync"
)

type Properties struct {
	Server ServerProps `yaml:"server"`
}

type ServerProps struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	ContextRoot string `yaml:"contextroot"`
}

type Server struct {
	Host        string
	Port        int
	ContextRoot string
	DoOnce      sync.Once
}

func NewServer(viper *viper.Viper, properties *Properties) *Server {
	if err := viper.Unmarshal(&properties, func(c *mapstructure.DecoderConfig) {
		c.DecodeHook = mapstructure.StringToTimeDurationHookFunc()
	}); err != nil {
		log.Warnf("❌ Unable to read application.yaml file : %s", err)
	}
	server := new(Server)
	server.Host = properties.Server.Host
	server.Port = properties.Server.Port
	server.ContextRoot = properties.Server.ContextRoot
	return server
}
