package proxy

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

// takes target url(string) and returns reverse proxy of type http.HandlerFunc
func ReverseProxy(Targethost string) http.HandlerFunc {

	// parser the target url(string) to url struct
	targetUrl, err := url.Parse(Targethost)

	if err != nil {
		log.Println(err)
	}

	reverseProxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new request to be forwaded to target server

		// set request host to target host
		r.Host = targetUrl.Host

		// set request url host to target host
		r.URL.Host = targetUrl.Host

		// set request scheme to target host scheme
		r.URL.Scheme = targetUrl.Scheme

		// set request URI to empty string
		r.RequestURI = ""

		// forward the newly created request to target server
		response, err := http.DefaultClient.Do(r)

		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// writer proxy-reponse response code to http.StatusOk(200)
		w.WriteHeader(http.StatusOK)

		// copy server-reponse body to proxy-response writer
		io.Copy(w, response.Body)

		// copy reponse headers to w.Headers
		for key, value := range response.Header {
			for _, value := range value {
				w.Header().Set(key, value)
			}
		}

	})
	return reverseProxy
}
