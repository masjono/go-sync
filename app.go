package main

import (

    "github.com/google/go-github/github"
	"log"
	"github.com/darwinsimon/go-sync-toped/src"
	"net/http"
    "golang.org/x/oauth2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.Println("Starting Go Sync Toped...")
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: "7bc1dee7f0a122bfb1cd2416b3e5a9225532b9d5" },
    )
	tc := oauth2.NewClient(oauth2.NoContext, ts)

    client := github.NewClient(tc)
    src.InitCron(client)

	log.Fatal(http.ListenAndServe(":9999", nil))
}
