package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
)

var ip string
var token string

func handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func handleIP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header["Authorization"]
	if auth == nil || auth[0] != token {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method == http.MethodPut {
		b := bytes.Buffer{}
		_, err := io.Copy(&b, r.Body)
		handle(err)
		ip = b.String()
		log.Printf("got ip: %s", ip)
		return
	} else if r.Method == http.MethodGet {
		_, err := w.Write([]byte(ip))
		handle(err)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	token = os.Getenv("TOKEN")
	log.Printf("TOKEN: %s", token)
	http.HandleFunc("/ip", handleIP)
	tlsCert := os.Getenv("TLS_CERT")
	tlsKey := os.Getenv("TLS_KEY")
	tlsOn := tlsCert != "" && tlsKey != ""
	if tlsOn {
		log.Printf("cert: %s, key: %s", tlsCert, tlsKey)
		go func() {
			log.Print(http.ListenAndServeTLS(os.Getenv("LISTEN_ADDRESS_TLS"), tlsCert, tlsKey, nil))
		}()
	}
	log.Print(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
