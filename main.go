package main

import (
    "context"
    "fmt"
    "net"
    "time"
    "os"
)

func results(r bool, ep string, t time.Duration) {
    if r {
        fmt.Printf("✔ %s [%s]\n", ep, t.Round(time.Millisecond))
    } else {
        fmt.Printf("✘ %s [%s]\n", ep, t.Round(time.Millisecond))
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
    var loop bool
    var endpoints []string
    for _, arg := range os.Args[1:] {
        if arg == "-l" {
            loop = true
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
