package main

import (
  "github.com/gorilla/mux"
  "os"
  "log"
  "net/http"
  "net"
  "flag"
  "time"
  "github.com/ajays20078/go-ini"
  "github.com/ajays20078/go_http_logger"
 
)


func check(e error) {
	if e != nil {
		panic(e)
	}
}


func check_boolean(e bool){
	if !e {
		panic("Unexpected Error!!!")
	}
}


func main() {


  wordPtr := flag.String("c", "/etc/goserver/config.ini", "config file")

  

  file, err := ini.LoadFile(*wordPtr)
  check(err)
	
  port, ok := file.Get("Server", "Port")
  check_boolean(ok)
  virtual_server,ok := file.Get("Server","VirtualHost")
  check_boolean(ok)
  access_log,ok := file.Get("Server","AccessLog")
  check_boolean(ok)


   access_file_handler, err := os.OpenFile(access_log, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
   if err != nil {
                panic(err)
   }
   defer access_file_handler.Close()


  rtr := mux.NewRouter()


  rtr.HandleFunc("/user/{name:[a-z]+}/profile", func(w http.ResponseWriter, r *http.Request) {
              profile(w, r, access_file_handler)
       }).Host(virtual_server).Methods("GET")


  http.Handle("/", rtr)

  log.Println("Listening...")
  http.ListenAndServe(":"+port, nil)


  log.Println("AFter the call...")
}

func profile(w http.ResponseWriter, r *http.Request,access_file_handler *os.File) {
  start_time := time.Now().UnixNano()
  params := mux.Vars(r)
  name := params["name"]
  return_string := "Hello " + name
  ip,port,_ := net.SplitHostPort(r.RemoteAddr)
  return_string = return_string + "\t IP is" +ip + ":" +port
  http_return_status := 200
  w.Write([]byte(return_string))
  defer go_http_logger.Log_line(r,http_return_status,return_string,access_file_handler,start_time,time.Now().UnixNano())


}
