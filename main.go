package main

import (
    "context"
    "fmt"
    "net"
    "time"
    "os"
)

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
            if t, check := tcp(os.Args[1]); check != nil {
                fmt.Printf("✘ %s [%s]\n", os.Args[1], t.Round(time.Millisecond))
            } else {
                fmt.Printf("✔ %s [%s]\n", os.Args[1], t.Round(time.Millisecond))
            }
            time.Sleep(time.Second * 1)
        }
    }

    if t, check := tcp(os.Args[1]); check != nil {
        fmt.Printf("✘ %s [%s]\n", os.Args[1], t.Round(time.Millisecond))
    } else {
        fmt.Printf("✔ %s [%s]\n", os.Args[1], t.Round(time.Millisecond))
    }
}
