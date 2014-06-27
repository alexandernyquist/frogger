package main

import (
	"flag"
	"fmt"
	"github.com/alexandernyquist/frogger/proxy"
)

type stringArgs []string

func (s *stringArgs) String() string {
	return fmt.Sprintf("%s", *s)
}

func (s *stringArgs) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var dumpHosts stringArgs
	portFlag := flag.Int("port", 8082, "Port to listen on")
	noCacheFlag := flag.Bool("nocache", false, "Disable caching")
	dumpAllFlag := flag.Bool("dumpall", false, "Dump all hosts")
	dumpHeadersFlag := flag.Bool("dumpheaders", false, "Include response headers in dump files")

	flag.Var(&dumpHosts, "dump", "Hosts to dump")
	flag.Parse()

	proxy := frogger.Proxy{*portFlag, *noCacheFlag, *dumpAllFlag, dumpHosts,  *dumpHeadersFlag}

	fmt.Printf("Serving frogger on :%d\n", *portFlag)
	err := proxy.Listen()
	if err != nil {
		fmt.Println("Could not listen on port 8082. Port probably already in use.")
	}
}
