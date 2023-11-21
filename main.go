package main

import (
    "context"
    "fmt"
    "net"
    "time"
    "os"
)

func tcp(addr string) error {
    var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()


    conn, err := d.DialContext(ctx, "tcp", addr)
    if err != nil {
        return err
    }
    conn.Close()
    return nil
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Expected a target\nExample: jonny.gg:443")
        os.Exit(1)
    }

    if check := tcp(os.Args[1]); check != nil {
        fmt.Printf("✘ %s\n", os.Args[1])
    } else {
        fmt.Printf("✔ %s\n", os.Args[1])
    }
}
