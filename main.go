package main 
import (
        "fmt"
          "net/http"
          )
 func main() {
 http.Hundleifunc("/",func(whttp.ResponseWriter, r "http.Request){
 fmt.fprintf(w,"Hello word!")
 })
 http.ListenAndServer(":8080",n11)
 )
