package v1

import (
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Proxy struct{}

// Proxy Handler ...
func (*Proxy) Handler(rw http.ResponseWriter, req *http.Request) {

	log.Printf("[reverse proxy server] received request from %s at: %s\n", req.RemoteAddr, time.Now().Format(time.RFC822))

	// define json place holder origin server URL
	originServerURL, err := url.Parse("https://jsonplaceholder.typicode.com")
	if err != nil {
		log.Fatal("invalid origin server URL")
	}

	// set req Host, URL and Request URI to forward a request to the origin server
	req.Host = originServerURL.Host
	req.URL.Host = originServerURL.Host
	req.URL.Scheme = originServerURL.Scheme
	req.RequestURI = ""

	delHopHeaders(req.Header)

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		appendHostToXForwardHeader(req.Header, clientIP)
	}

	// save the response from the origin server
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		http.Error(rw, "Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	log.Printf("[reverse proxy server] received response status code: %s", resp.Status)

	delHopHeaders(resp.Header)

	copyHeader(rw.Header(), resp.Header)

	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)

}
