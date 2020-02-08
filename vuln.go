package vuln

import (
    "fmt"
    "net/http"
    "crypto/tls"
)

func main() {
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
    _, err := http.Get("https://golang.org/")
    if err != nil {
        fmt.Println(err)
    }
}