package ws

import (
    "log"
)

const(
    WsRoom = "room"
    JsonRoom = "room"
    JsonMsg = "msg"
)

func Server(ip string, url string, cert string, key string, cmdchan <-chan string){

    log.Println("ws task starting...")
    log.Printf("ws server ip=%v url=%v cert=%v key=%v\n", ip, url, cert, key)

    go cmd(cmdchan)
    go wsserver(ip, url, cert, key)
}
