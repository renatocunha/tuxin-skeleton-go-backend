package main

import (
	"log"
	"github.com/go-ini/ini"
)

const (
	configIniPath string = "./config/config.ini"
)

type SslCert struct {
	CertificateFile string
	PrivateKeyFile string
	OtherCertificates []string
}

type Server struct {
	IsProduction bool
	ServerName string
}

type Auth0 struct {
	AccountDomain string
}

type NewRelic struct {
	Licensekey string
	AppName string
}

type WebServer struct {
	ClientPath string
	Routes []string
	HttpsPort uint64
	ReadTimeout uint
	WriteTimeout uint
}

type ConfigIni struct {
	SslCert
	Server
	NewRelic
	Auth0
	WebServer
}

var CfgIni *ConfigIni

func parseIni(configIniPath string) (*ConfigIni) {
	cfg, err := ini.Load(configIniPath);
	if (err != nil) {
		log.Fatalf("error loading config.ini: %v", err);
	}
	cfgIni := new(ConfigIni);
	err = cfg.MapTo(cfgIni);
	if (err != nil) {
		log.Fatalf("error parsing config.ini: %v", err);
	}
	return cfgIni
}

