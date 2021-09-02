package main

import (
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"
	"gopkg.in/yaml.v3"

	"reverseProxy/internal/proxyApi"
)

type Configs struct {
	ConfigItems []struct {
		From string `yaml:"from"`
		To   string `yaml:"to"`
	} `yaml:"config"`
}

func main() {
	cont, err := os.ReadFile("./configs/config.yml")
	if err != nil {
		log.Fatalln(err)
	}
	var configs Configs
	if err = yaml.Unmarshal(cont, &configs); err != nil {
		log.Fatalln(err)
	}

	proxy := proxyApi.NewReverseProxy()
	for _, item := range configs.ConfigItems {
		fmt.Printf("from: %s to %s \n", item.From, item.To)
		proxy.AddHost(item.From, item.To)
	}
	handler := proxy.GenerateProxyHandler()
	if err := fasthttp.ListenAndServe(":9000", handler); err != nil {
		log.Fatal(err)
	}
}
