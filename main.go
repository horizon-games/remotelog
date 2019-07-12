package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var (
	flags = flag.NewFlagSet("remotelog", flag.ExitOnError)
	host  = flags.String("host", "0.0.0.0", "http server hostname")
	port  = flags.String("port", "1111", "http server port")
)

func main() {
	flags.Parse(os.Args[1:])

	addr := fmt.Sprintf("%s:%s", *host, *port)

	err := startServer(addr)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer(addr string) error {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("."))
	})

	webrpcHandler := NewRemoteLogServer(&RemoteLogRPC{})
	r.Handle("/*", webrpcHandler)

	return http.ListenAndServe(addr, r)
}

type RemoteLogRPC struct {
}

func (s *RemoteLogRPC) Ping(ctx context.Context) (bool, time.Time, error) {
	return true, time.Now(), nil
}

func (s *RemoteLogRPC) Log(ctx context.Context, msg string, object map[string]interface{}) (bool, error) {
	if len(object) > 0 {
		fmt.Printf("%s -- %s -- %v\n", time.Now().Format(time.RFC1123), msg, object)
	} else {
		fmt.Printf("%s -- %s\n", time.Now().Format(time.RFC1123), msg)
	}
	return true, nil
}
