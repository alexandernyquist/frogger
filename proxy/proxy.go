package frogger

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"strconv"
)

var dumpDelimiter = "------------------------------"
var headersNotForwarded = []string{"Host", "Content-Length", "Connection", "Proxy-Connection", "Accept-Encoding"}

type Proxy struct {
	Port      int
	DumpHosts []string
}

func (p Proxy) Listen() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handleRequest(w, r, p)
	})

	err := http.ListenAndServe(":" + strconv.Itoa(p.Port), nil)
	if err != nil {
		return err
	}

	return nil
}

func (p Proxy) shouldDump(host string) bool {
	for _, h := range p.DumpHosts {
		if h == host {
			return true
		}
	}

	return false
}

func joinHeaders(headers http.Header) string {
	var result string

	for k, v := range headers {
		result += k + ": " + strings.Join(v, ";") + "\n"
	}

	return result
}

func handleRequest(w http.ResponseWriter, req *http.Request, p Proxy) {
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

	host := req.URL.Host
	if p.shouldDump(host) {
		// Dump request to disk

		// Create directory if not exists
		os.Mkdir(host, 0666)

		// Write file to disk
		f, err := ioutil.TempFile(host, host+"-")
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		defer f.Close()

		// Write to dump file and response
		writers := io.MultiWriter(f, w)
		_, err = io.Copy(writers, resp.Body)
		if err != nil {
			log.Fatal("Could not write to dump file or response: %v", err)
		}

		// Append dump info to bottom of dump file
		f.WriteString("\n\n" + dumpDelimiter)
		f.WriteString("\n" + req.URL.String())
		f.WriteString("\n" + resp.Proto + " " + resp.Status)
		f.WriteString("\n" + joinHeaders(resp.Header))
	} else {
		// Write response directly
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			log.Fatal("Could not write to response: %v", err)
		}
	}
}
