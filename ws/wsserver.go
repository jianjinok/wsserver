package ws

import (
    "net/http"
    "log"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:1024,
    WriteBufferSize:1024,
    CheckOrigin:func(r *http.Request) bool{
        return true
    },
}

func wsRead(wbscon *websocket.Conn) error{
    for{
        len,data,err := wbscon.ReadMessage()
        log.Printf("ws read len=%#v,data=%#v,err=%#v",len,data,err)
        if err !=nil{
            return err
        }
    }
}

func wsMsgProc(room string, wbscon *websocket.Conn)error{
    var err error
    wbschan := make(chan []byte, 5)

    roomAddChannel(room, wbschan)
    for
    {
        for msg := range wbschan{
            log.Printf("room=%s send msg=%s\n",room, string(msg))
            err = wbscon.WriteMessage(websocket.TextMessage, msg)
            if err != nil{goto ERR}
        }
    }
    ERR:
        roomDelChannel(room, wbschan)
        close(wbschan)
        return err
}

func wsHandler(w http.ResponseWriter, r *http.Request){
    var (
    wbsCon *websocket.Conn
    err error
    )

    log.Printf("wsHandler %v\n", r.URL)
    room,ok := r.URL.Query()[WsRoom]
    if !ok{
        log.Println("no find room")    
        return
    }

    if wbsCon, err = upgrader.Upgrade(w, r, nil); err != nil {
        return
    }
    
    go wsMsgProc(room[0], wbsCon)
    if err = wsRead(wbsCon); err != nil{goto ERR}

    ERR:
        wbsCon.Close()
}

func wsserver(ip string, url string, cert string, key string){
    http.HandleFunc(url, wsHandler)
    err := http.ListenAndServeTLS(ip, cert, key, nil)
    if err != nil{
        log.Fatal("Listen Server", err.Error())
    }
}

