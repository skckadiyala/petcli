package cmd

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/skckadiyala/petstore"
	"github.com/spf13/viper"
)

var (
	file     string // blob file location
	username string
	host     string
	port     string
	password string
)

type configAPI struct {
	PetHost       string `yaml:"petHost"`
	PetPort       string `yaml:"petPort"`
	Authorization string `yaml:"Authorization"`
}

func getConfig() *petstore.Configuration {
	// if viper.GetString("apimanagerhost") == "" || viper.GetString("apimanagerport") == "" {
	// 	utils.PrettyPrintErr("Please login to API Manager, use 'login' command")
	// 	os.Exit(0)
	// }
	fmt.Println("Value of host", viper.GetString("apimanagerhost"))

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
	}

	cfg := petstore.NewConfiguration()
	cfg.Host = "petstore.swagger.io" //viper.GetString("pethost") + ":" + viper.GetString("petport")
	cfg.Scheme = "https"
	// cfg.AddDefaultHeader("Authorization", "Basic "+viper.GetString("authorization"))
	cfg.HTTPClient = &http.Client{Transport: transCfg}
	return cfg
}
