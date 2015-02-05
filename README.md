# go-http-logger



This is a asynchronous non-blocking logger for go HTTP Server which logs the requests in the below format

%Date% %Time% %VirtualHost:Port% %ClientHost:Port% "%HttpMethod% %URL?Paramaters% %HttpVersion%" %HttpReturnCode% %ResponseByteSize% "%HttpClient%" %ResponseTime%

Where ResponseTime is in MilliSeconds



It takes the following arguments:

request - HTTP request from the client
return_status - HTTP return status from the server
return_string - HTTP response string
file_handler - File to which Log has to be written to
start_time - localtime in Nanoseconds (You can use time.Now().UnixNano()) when the HTTP handler starts its execution
end_time - localtime in Nanoseconds (You can use time.Now().UnixNano()) when the HTTP handler finishes its execution
