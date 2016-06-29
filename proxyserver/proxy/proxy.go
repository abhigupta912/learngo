package proxy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type ProxyServer struct {
	httpClient *http.Client
}

func NewProxyServer(client *http.Client) *ProxyServer {
	if client == nil {
		client = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	return &ProxyServer{client}
}

func (server *ProxyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")

	values := r.URL.Query()
	urlString := values.Get("url")
	if urlString == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Use query parameter url")
		return
	}

	finalUrl, err := url.Parse(urlString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		log.Println("Unable to parse url")
		return
	}

	resp, err := server.httpClient.Get(finalUrl.String())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		log.Println("Unable to GET url")
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err.Error())
		log.Println("Unable to read response from url")
		return
	}

	data := bytes.NewBuffer(body)
	fmt.Fprintln(w, data.String())
}
