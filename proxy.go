package frogger

import(
	"io"
	"log"
	"net/http"
)

var headersNotForwarded = []string{"Host", "Content-Length", "Connection", "Proxy-Connection", "Accept-Encoding"}

type Proxy struct {
	Port int
}

func (p Proxy) Listen() error {
	http.HandleFunc("/", handleRequest)
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		return err
	}

	return nil
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	log.Println("Incoming request to " + req.URL.String())

	// Delete headers that proxy should not forward
	for _, h := range headersNotForwarded {
		req.Header.Del(h)
	}

	// Request actual page
	clientReq := &http.Request{Method: "GET", URL: req.URL, Header: req.Header}
	tr := &http.Transport{}
	resp, err := tr.RoundTrip(clientReq)	
	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Write server response headers back to client
	for k, v := range resp.Header {
		w.Header().Set(k, v[0])
	}
	w.Header().Set("X-Forwarded-For", "Frogger") // For debugging, should contain client and proxy id
	w.WriteHeader(resp.StatusCode)

	// Write body
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return
	}
}