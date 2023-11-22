package main

import (
    "context"
    "fmt"
    "net"
    "time"
    "os"
)

var loop bool
var colours bool
var endpoints []string

func results(r bool, ep string, t time.Duration) {
    if r {
        if colours {
            fmt.Printf("\033[0;32m✔\033[0m %s [%s]\n", ep, t.Round(time.Millisecond))
        } else {
            fmt.Printf("✔ %s [%s]\n", ep, t.Round(time.Millisecond))
        }
    } else {
        if colours {
            fmt.Printf("\033[0;31m✘\033[0m %s [%s]\n", ep, t.Round(time.Millisecond))
        } else {
            fmt.Printf("✘ %s [%s]\n", ep, t.Round(time.Millisecond))
        }
    }
}

func tcp(addr string) (time.Duration, error) {
    var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

    start := time.Now()
    conn, err := d.DialContext(ctx, "tcp", addr)
    end := time.Now()
    elapsed := end.Sub(start)
    if err != nil {
        return elapsed, err
    }
    conn.Close()

    return elapsed, nil
}

func main() {
    // Check we have enough arguments to do anything
    if len(os.Args) < 2 {
        fmt.Println("Expected a target\nExample: jonny.gg:443")
        os.Exit(1)
    }

    // check for flags and build slice of endpoints
    colours = true
    for _, arg := range os.Args[1:] {
        if arg == "-l" {
            loop = true
            continue
        } else if arg == "--no-colours" {
            colours = false
            continue
        }
        endpoints = append(endpoints, arg)
    }


    if loop {
        for {
            for _, ep := range endpoints {
                if t, err := tcp(ep); err == nil {
                    results(true, ep, t)
                } else {
                    results(false, ep, t)
                }
            }
        time.Sleep(time.Second * 1)
        }
    }

    for _, ep := range endpoints {
        if t, err := tcp(ep); err == nil {
            results(true, ep, t)
        } else {
            results(false, ep, t)
        }
    }
}
