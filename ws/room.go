package ws

import (
    "sync"
    "encoding/json"
    "log"
)

var connmap = make(map[string][]chan []byte)
var mapRwLock sync.RWMutex

func roomSendMsg(room string, msg interface{}){

    log.Printf("room=%s msg=%#v\n", room, msg)
    mapRwLock.RLock()
    msgstr,_ := json.Marshal(msg)
    connchans, ok := connmap[room]
    if !ok {goto ERR}

    for _,conn := range connchans{
        if conn == nil{continue}

        log.Printf("send msg to chan %#v\n", conn)
        conn<- msgstr
    }
    ERR:
    mapRwLock.RUnlock()
}

func roomAddChannel(room string, connchan chan []byte){

    mapRwLock.Lock()
    if _,ok := connmap[room]; !ok{
        connchans := [] chan []byte{connchan}
        connmap[room] = connchans
    }else {
        connmap[room] = append(connmap[room], connchan) 
    }
    mapRwLock.Unlock()
    log.Printf("chan %v add room %v list %v", connchan, room, connmap[room])
}

func roomDelChannel(room string, connchan chan []byte){

    mapRwLock.Lock()
    connchans,ok := connmap[room]
    newconnchans := []chan []byte{}
    if !ok{goto ERR}

    for _,value := range connchans {
        if connchan == value {
            continue
        }
        newconnchans = append(newconnchans, value)
    }
    connmap[room] = newconnchans
    ERR:
    mapRwLock.Unlock()
    log.Printf("chan %v del from room %v list %v", connchan, room, connmap[room])
}

