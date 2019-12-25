package main


import (
    "golang.org/x/net/websocket"
    "fmt"
    "log"
    "math/rand"
    "time"
)

var conNum = 0

func main() {
    var origin = "http://localhost/"
    var url = "ws://localhost/ws?device_id="
    for i:=0; i<10; i++{
        newUrl := url+RandStringBytes(10)
        fmt.Println(newUrl);
        go wsTest(newUrl, origin)
    }
    time.Sleep(50 * time.Second)
}


func wsTest(url string, origin string){
    ws, err := websocket.Dial(url, "", origin)
    if err != nil {
        log.Fatal(err)
    }
    conNum++
    go func(){
        for{
            message := []byte("hello, world!你好")
            _, err = ws.Write(message)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("Send: %s\n", message)
            time.Sleep(1 * time.Second)
        }
    }()

    go func(){
        for{
            var msg = make([]byte, 512)
            m, err := ws.Read(msg)
            if err != nil {
                log.Fatal(err)
            }
            fmt.Printf("Receive: %s\n", msg[:m])
        }
    }()
    fmt.Println(conNum);
}


const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}