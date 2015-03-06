package main
import (
	"net/http"
	"github.com/vaughan0/go-ini"
	"flag"
	"log"
)
var (
	bindAddress = "127.0.0.1:8800"
	pathServe = "/mnt/my-data"
)
var conf = flag.String("f", "", "Config file")

func main() {
	flag.Parse()
	if len(*conf) > 0 {
		confFile, e := ini.LoadFile(*conf)
		if e != nil {
			log.Println(e)
		}
		if bindAddr, ok := confFile.Get("", "bindaddress"); ok {
			bindAddress = bindAddr
		}
		if pserv, ok := confFile.Get("", "pathserve"); ok {
			pathServe = pserv
		}
	}
        panic(http.ListenAndServe(bindAddress, http.FileServer(http.Dir(pathServe))))
}
