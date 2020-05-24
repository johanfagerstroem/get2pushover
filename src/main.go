package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	defaultUser  string
	defaultToken string
	xVersion     string
)

func getRemoteFQDN(r *http.Request) (string, error) {
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}

	names, err := net.LookupAddr(host)
	if err != nil {
		return "", err
	}

	return names[0], nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		token := q.Get("token")
		user := q.Get("user")
		message := q.Get("message")
		title := q.Get("title")

		if message == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		/* if token not provided by caller; default to environment */
		if token == "" {
			if defaultToken == "" {
				log.Printf("[%s] No token provided in request and no default set\n", r.RemoteAddr)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			token = defaultToken
		}

		/* if user not provided by caller; default to environment */
		if user == "" {
			if defaultUser == "" {
				log.Printf("[%s] No user provided in request and no default set\n", r.RemoteAddr)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			user = defaultUser
		}

		/* if title not provided by caller; default to remote hostname */
		if title == "" {
			var err error
			title, err = getRemoteFQDN(r)
			if err != nil {
				log.Printf("[%s] Failed to get remote FQDN for request: %v\n", r.RemoteAddr, err)
				/* default to get2pushover */
				title = "get2pushover"
			}
		}

		form := url.Values{}
		form.Add("token", token)
		form.Add("user", user)
		form.Add("message", message)
		form.Add("title", title)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		pushReq, err := http.NewRequestWithContext(ctx, http.MethodPost,
			"https://api.pushover.net/1/messages.json", strings.NewReader(form.Encode()))
		if err != nil {
			log.Printf("[%s] Failed to create Pushover request: %v\n", r.RemoteAddr, err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		client := http.Client{}
		pushResp, err := client.Do(pushReq)
		if err != nil {
			log.Printf("[%s] Failed to send Pushover notification: %v\n", r.RemoteAddr, err)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

		if pushResp.StatusCode != 200 {
			log.Printf("[%s] Bad status code from Pushover: %d\n", r.RemoteAddr, pushResp.StatusCode)
			w.WriteHeader(http.StatusBadGateway)
			return
		}

	})

	defaultUser = os.Getenv("PUSHOVER_DEFAULT_USER")
	defaultToken = os.Getenv("PUSHOVER_DEFAULT_TOKEN")

	port := os.Getenv("LISTEN_PORT")
	if port == "" {
		log.Printf("Listening port must be specified!\n")
		os.Exit(1)
	}

	log.Printf("Starting get2pushover %s", xVersion)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
