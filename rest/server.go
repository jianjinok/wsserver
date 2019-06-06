package rest

import(
    "io/ioutil"
    "log"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"
    )

var wschan chan<- string
var routes = [] *rest.Route {
    rest.Post("/screenws", screenws),
}

func restServer(ip string, url string){

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(routes...)
    if err != nil{
        log.Fatal(err)
    }
    api.SetApp(router)
    http.Handle(url+"/", http.StripPrefix(url, api.MakeHandler()))
    log.Fatal(http.ListenAndServe(ip, nil))
}

func screenws(w rest.ResponseWriter, req *rest.Request){
    w.WriteJson(map[string]string{"status":"ok"})
    jsonbytes, _ := ioutil.ReadAll(req.Body)
    jsonstr := string(jsonbytes)
    log.Printf("rest recv json=%s", jsonstr)
    wschan<-jsonstr
}

func Server(ip string, url string, ws chan<-string){
    log.Println("rest server start...")
    log.Printf("rest server ip=%v url=%v\n", ip, url)
    wschan = ws
    restServer(ip, url)
}

