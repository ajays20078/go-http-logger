# go-http-logger

WriteLog Logs the Http Status for a request into fileHandler
Syntax : http.ListenAndServe(Virtual_Host+":"+port, httpLogger.WriteLog(http.DefaultServeMux,fileHandler))

