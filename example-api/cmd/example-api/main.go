package main

// START OMIT
import (
	"flag"
	"log"

	"github.com/asasmoyo/halo123/api"
	"github.com/asasmoyo/techtalk93-jdv/example-api/http"
)

func main() {
	var (
		listenIP   string
		listenPort int
	)

	flag.StringVar(&listenIP, "listen-ip", "localhost", "Listen IP")
	flag.IntVar(&listenPort, "listen-port", 9000, "Listen Port")
	flag.Parse()
	// END OMIT
	log.Println("Starting server")

	var s api.Server = &http.Server{
		ListenIP:   listenIP,
		ListenPort: listenPort,
	}
	s.Run()
}
