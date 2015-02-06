# go-http-logger



go-http-logger is a wrapper around http handler for go HTTP Server which logs the requests in the below format

**%YYYY/MM/DD% %HH:MM:SS% %VirtualHost:Port% %ClientHost:Port% "%HttpMethod% %URL?Paramaters% %HttpVersion%" %HttpReturnCode% %ResponseByteSize% "%HttpClient%" %ResponseTime%**


It takes the following arguments:  <br/>

* request - HTTP request from the client  <br/>
* file_handler - File to which Log has to be written to  <br/>


**Syntax:**

http.ListenAndServe(Virtual_Host+":"+port, httpLogger.WriteLog(http.DefaultServeMux,fileHandler)) <br/>
