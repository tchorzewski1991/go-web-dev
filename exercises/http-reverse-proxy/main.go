package main

import (
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"log"
	"bytes"
	"strings"
	url2 "net/url"
	"net/http/httputil"
)

type requestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

// Provides callback when value for PORT environment variable is missing.
func getPort(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return  value
	}

	return fallback
}

// Setups value for :PORT environment variable that will be used to listen on.
func getListenAddress() string {
	port := getPort("PORT", "1338")
	return ":" + port
}

// Logs every variable that will be used to run reverse proxy server.
func logSetup() {

	a_cond_url := os.Getenv("A_COND_URL")
	b_cond_url := os.Getenv("B_COND_URL")
	default_cond_url := os.Getenv("DEFAULT_COND_URL")

	fmt.Printf("Server will run on: %s\n", getListenAddress())
	fmt.Printf("Redirecting to A URI: %s \n", a_cond_url)
	fmt.Printf("Redirecting to B URI: %s \n", b_cond_url)
	fmt.Printf("Redirecting to DEFAULT URI: %s \n", default_cond_url)
}

// requestPayloadDecoder() creates new *json.Decoder from current request.
// It is worth to remember that *http.Request.Body is representation of
// io.ReadCloser interface. We cannot read from it more than ONCE.
// Workaround for this issue is relevant. We should create a new io.ReadCloser
// interface from buffered bytes that we already have, and reassign it to
// actual *http.Request.Body value. Unfortunately, while creating new
// io.Reader interface with bytes.NewBuffer is easy, we need to remember
// that Buffer doesn't define Close() method. To expand Buffer with
// Closer interface we can use ioutil.NopCloser() method. This method wraps
// io.Reader and returns io.ReadCloser interface with no-op Close() method.
// 'no-op' means it does nothing.
func requestPayloadDecoder(request *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalf("Error while reading body: %s\n", err)
		panic(err)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

// parseRequestPayload() creates new requestPayloadStruct from actual request.
// Good to remember that struct is a value type. It needs to be given to the
// decoder.Decode() function by memory address if we want to modify its content.
func parseRequestPayload(request *http.Request) requestPayloadStruct {
	decoder := requestPayloadDecoder(request)

	var requestPayload requestPayloadStruct
	err := decoder.Decode(&requestPayload)

	if err != nil {
		panic(err)
	}

	return requestPayload
}

func logRequestPayload(requestPayload requestPayloadStruct, proxyUrl string) {
	log.Printf("proxy_condition: %s, proxy_url: %s\n", requestPayload.ProxyCondition, proxyUrl)
}

func getProxyUrl(proxyConditionRaw string) string {
	proxy_condition := strings.ToUpper(proxyConditionRaw)

	a_condition_url := os.Getenv("A_COND_URL")
	b_condition_url := os.Getenv("B_COND_URL")
	default_condition_url := os.Getenv("DEFAULT_COND_URL")

	if proxy_condition == "A" {
		return a_condition_url
	}

	if proxy_condition == "B" {
		return b_condition_url
	}

	return default_condition_url
}

// Serves a reverse proxy for given url.
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	url, _ := url2.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme

	// 'X-Forwarded-Host' identifies original HOST requested by the client.
	// Hosts names and ports of reverse proxies (load-balancers, CND's) may differ,
	// so this header is useful to determine which Host was originally used.
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))

	req.Host = url.Host

	proxy.ServeHTTP(res, req)
}

// Given a request sends it to appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	requestPayload := parseRequestPayload(req)

	url := getProxyUrl(requestPayload.ProxyCondition)

	logRequestPayload(requestPayload, url)

	serveReverseProxy(url, res, req)
}


func main()  {
	// Log actual setup.
	logSetup()

	// Registers new handle function. Let reverse proxy server to decide
	// where to redirect appropriate requests.
	http.HandleFunc("/", handleRequestAndRedirect)

	// Starts new server using PORT environment variable.
	if err := http.ListenAndServe(getListenAddress(), nil); err != nil {
		panic(err)
	}
}
