package ws

import (
    "log"
    "encoding/json"
)

func cmd(wschan <-chan string){
    var room, msg interface{}
    var ok bool

    for jsonstr := range wschan{
        log.Printf("ws channel recv json=%s\n", jsonstr)
        jsonMap := make(map[string]interface{})
        json.Unmarshal([]byte(jsonstr), &jsonMap)

        if room, ok = jsonMap[JsonRoom];!ok{continue}
        if msg, ok = jsonMap[JsonMsg];!ok{continue}

        roomSendMsg(room.(string), msg)
   }
}
