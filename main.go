package main

import (
    "context"
    "fmt"
    "net"
    "time"
    "os"
    "strings"
)

var loop bool
var colours bool
var endpoints []*endpoint

type endpoint struct {
    success bool
    addr    string
    time    time.Duration
}

func results(ep endpoint, p int) {
    var padding string
    padding = strings.Repeat(" ", p - len(ep.addr))
    if ep.success {
        if colours {
            fmt.Printf("\033[0;32m✔\033[0m %s%s [%s]\n", ep.addr, padding, ep.time.Round(time.Millisecond))
        } else {
            fmt.Printf("✔ %s%s [%s]\n", ep.addr, padding, ep.time.Round(time.Millisecond))
        }
    } else {
        if colours {
            fmt.Printf("\033[0;31m✘\033[0m %s%s [%s]\n", ep.addr, padding, ep.time.Round(time.Millisecond))
        } else {
            fmt.Printf("✘ %s%s [%s]\n", ep.addr, padding, ep.time.Round(time.Millisecond))
        }
    }
}

func tcp(ep *endpoint) (error) {
    var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

    start := time.Now()
    conn, err := d.DialContext(ctx, "tcp", ep.addr)
    end := time.Now()
    ep.time = end.Sub(start)
    if err != nil {
        ep.success = false
        return err
    }
    conn.Close()
    ep.success = true

    return nil
}

func main() {
    // Check we have enough arguments to do anything
    if len(os.Args) < 2 {
        fmt.Println("Expected a target\nExample: jonny.gg:443")
        os.Exit(1)
    }

    // check for flags and build slice of endpoints
    var padding int
    colours = true
    for _, arg := range os.Args[1:] {
        if arg == "-l" {
            loop = true
            continue
        } else if arg == "--no-colours" {
            colours = false
            continue
        }
        endpoints = append(endpoints, &endpoint{addr: arg})
        // set padding size
        if len(arg) > padding {
            padding = len(arg)
        }
    }

    // start tcptest loop if it's a loop
    if loop {
        for {
            for _, ep := range endpoints {
                err := tcp(ep)
                if err == nil {
                    results(*ep, padding)
                } else {
                    results(*ep, padding)
                }
            }
        time.Sleep(time.Second * 1)
        }
    }

    // if not just do the test once
    for _, ep := range endpoints {
        err := tcp(ep)
        if err == nil {
            results(*ep, padding)
        } else {
            results(*ep, padding)
        }
    }
}
