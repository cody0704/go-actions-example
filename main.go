package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	var port string
	var ping bool

	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.BoolVar(&ping, "ping", false, "check service live")

	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	if ping {
		if err := pinger(port); err != nil {
			log.Printf("ping service err:%v\n", err)
		}

		return
	}

	http.HandleFunc("/", handle)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got http request. time: %v", time.Now())
	fmt.Fprintf(w, "I love %s!, %s", r.URL.Path[:1], HelloActions())
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("server returned not-200 status code")
	}

	return nil
}

// HelloActions for hello actions
func HelloActions() string {
	return "Hello Github Actions, traefik workshop!"
}
