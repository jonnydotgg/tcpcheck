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
        fmt.Printf("✔  %s [%s]\n", ep, t.Round(time.Millisecond))
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
    if len(os.Args) < 1 {
        fmt.Println("Expected a target\nExample: jonny.gg:443")
        os.Exit(1)
    }

    if len(os.Args) == 3 && os.Args[2] == "-l" {
        for {
            if t, err := tcp(os.Args[1]); err == nil {
                results(true, os.Args[1], t)
            } else {
                results(false, os.Args[1], t)
            }
            time.Sleep(time.Second * 1)
        }
    }

    if t, err := tcp(os.Args[1]); err == nil {
        results(true, os.Args[1], t)
    } else {
        results(false, os.Args[1], t)
    }
}
