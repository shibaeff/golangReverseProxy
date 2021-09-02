package proxyApi

import (
	"log"
	"time"

	loglib "github.com/yeqown/log"

	"github.com/valyala/fasthttp"
	proxy "github.com/yeqown/fasthttp-reverse-proxy/v2"
)

type ReversePrxoy interface {
	AddHost(string, string)
	GenerateProxyHandler() func(ctx *fasthttp.RequestCtx)
}

type reverseProxy struct {
	proxies []*proxy.ReverseProxy
	index   map[string]int
}

func (r *reverseProxy) AddHost(from string, to string) {
	proxyServer := proxy.NewReverseProxy(to, proxy.WithTimeout(time.Second*time.Duration(2)))
	r.proxies = append(r.proxies, proxyServer)
	r.index[from] = len(r.proxies) - 1
}

func trim(s string) int {
	for i := 1; i < len(s); i++ {
		if s[i] == '/' {
			return i + 1
		}
	}
	return -1
}

func (r *reverseProxy) GenerateProxyHandler() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		requestURI := string(ctx.RequestURI())
		loglib.Info("requestURI=", requestURI)
		slashPos := trim(requestURI)
		if slashPos == -1 {
			hostNotFoundError(ctx)
			return
		}

		if ind, ok := r.index[requestURI[0:slashPos-1]]; ok {
			ctx.Request.SetRequestURI(requestURI[slashPos-1:])
			r.proxies[ind].ServeHTTP(ctx)
		} else {
			hostNotFoundError(ctx)
		}
	}
}

func hostNotFoundError(ctx *fasthttp.RequestCtx) {
	ctx.Error("No such host prefix", 500)
	log.Print("No such host prefix")
}

func NewReverseProxy() ReversePrxoy {
	return &reverseProxy{
		index: make(map[string]int),
	}
}
