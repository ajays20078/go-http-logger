# go-http-logger



go-http-logger is a asynchronous non-blocking logger for go HTTP Server which logs the requests in the below format

**%YYYY/MM/DD% %HH:MM:SS% %VirtualHost:Port% %ClientHost:Port% "%HttpMethod% %URL?Paramaters% %HttpVersion%" %HttpReturnCode% %ResponseByteSize% "%HttpClient%" %ResponseTime%**

Where ResponseTime is in MilliSeconds



It takes the following arguments:  <br/>

* request - HTTP request from the client  <br/>
* return_status - HTTP return status from the server  <br/>
* return_string - HTTP response string  <br/>
* file_handler - File to which Log has to be written to  <br/>
* start_time - localtime in Nanoseconds (You can use time.Now().UnixNano()) when the HTTP handler starts its execution  <br/>
* end_time - localtime in Nanoseconds (You can use time.Now().UnixNano()) when the HTTP handler finishes its execution  <br/>
