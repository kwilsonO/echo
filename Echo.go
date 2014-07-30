package main

import(

  "./web"
  atlantis "./atlantis/types"
  "fmt"
  "log"
)

var(
   listenAddr string
   //db driver.Conn
)

func init() {
  cfg, err := atlantis.LoadAppConfig()
  if err != nil {
     log.Printf("error opening using default port")
     listenAddr = ":9998"
     return
  }
  listenAddr = fmt.Sprintf(":%d", cfg.HTTPPort)
  
} 

func echo(ctx *web.Context, val string){  
   ctx.ContentType("text/plain")
   ctx.ResponseWriter.Write([]byte("I am the echo server woot!\n"))
}
func healthz(ctx *web.Context, val string){
   ctx.ContentType("text/plain")
   //ctx.ResponseWriter.Header().Add("Server-Status", "OK")
   ctx.ResponseWriter.Write([]byte("OK\n"))
}

func main(){
   web.Get("/(healthz)", healthz)
   web.Get("/(.*)", echo)
   web.Run("0.0.0.0" + listenAddr)
}
