// https://www.jianshu.com/p/55c1540438f5
package main

import (
    "fmt"
    "log"
    "flag"
    "net/http"
    "io/ioutil"
    "crypto/tls"
    "crypto/x509"
    "encoding/json"
    // "github.com/gorilla/mux"
)

var (
    port       int
    hostname   string 
    caroots    string
    keyfile    string
    signcert   string
)

func init() {
    flag.IntVar(&port,          "port",     8080,       "The host port on which the REST server will listen")
    flag.StringVar(&hostname,   "hostname", "0.0.0.0",  "The host name on which the REST server will listen")
    flag.StringVar(&caroots,    "caroot",   "",         "Path to file containing PEM-encoded trusted certificate(s) for clients")
    flag.StringVar(&keyfile,    "key",      "",         "Path to file containing PEM-encoded key file for service")
    flag.StringVar(&signcert,   "signcert", "",         "Path to file containing PEM-encoded sign certificate for service")
}

func startServer(address string, caroots string, keyfile string, signcert string) {
    pool := x509.NewCertPool()

    caCrt, err := ioutil.ReadFile(caroots)
    if err != nil {
        log.Fatalln("ReadFile err:", err)
    }
    pool.AppendCertsFromPEM(caCrt)

    s := &http.Server{
            Addr:    address,
            // Handler: router,
            TLSConfig: &tls.Config{
                MinVersion: tls.VersionTLS12,
                ClientCAs:  pool,
                ClientAuth: tls.RequireAndVerifyClientCert,
            },
    }
    err = s.ListenAndServeTLS(signcert, keyfile)
    if err != nil {
        log.Fatalln("ListenAndServeTLS err:", err)
    }
}

func SayHello(w http.ResponseWriter, r *http.Request) {
    log.Println("Entry SayHello")
    res := map[string]string {"hello": "world"}

    b, err := json.Marshal(res)
    if err == nil {
        w.WriteHeader(http.StatusOK)
        w.Header().Set("Content-Type", "application/json")
        w.Write(b)
    }

    log.Println("Exit SayHello")
}

func main() {
    flag.Parse()

    // router := mux.NewRouter().StrictSlash(true)
    http.HandleFunc("/service/hello", SayHello)

    var address = fmt.Sprintf("%s:%d", hostname, port)
    fmt.Println("Server listen on", address)
    startServer(address, caroots, keyfile, signcert)
    
    fmt.Println("Exit main")
}

