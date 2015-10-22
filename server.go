// Our webserver
package main

import (
    "fmt"
    "net/http"
    "github.com/guregu/kami"
    "golang.org/x/net/context"
    "github.com/r-fujiwara/kamigami-no-asobi/greeting" // see package greeting below
)

func greet(ctx context.Context, w http.ResponseWriter, r *http.Request) {
    hello := greeting.FromContext(ctx)
    name := kami.Param(ctx, "name")
    fmt.Fprintf(w, "%s, %s!", hello, name)
}

func main() {
    ctx := context.Background()
    ctx = greeting.WithContext(ctx, "Hello") // set default greeting
    kami.Context = ctx                       // set our "god context", the base context for all requests

    kami.Use("/hello/", greeting.Guess) // use this middleware for paths under /hello/
    kami.Get("/hello/:name", greet)     // add a GET handler with a parameter in the URL
    kami.Serve()                        // gracefully serve with support for einhorn and systemd
}
