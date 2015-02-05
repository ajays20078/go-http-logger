package http_logger

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

func Log_line(request *http.Request, return_status int, return_string string, file_handler *os.File, start_time int64, end_time int64) int {

	go Log_line_async(request, return_status, return_string, file_handler, start_time, end_time)
	return 0

}

func Log_line_async(request *http.Request, return_status int, return_string string, file_handler *os.File, start_time int64, end_time int64) {

	time_diff := (float64(end_time) - float64(start_time)) / 1000000.0
	time_taken_string := strconv.FormatFloat(time_diff, 'f', 3, 64)
	log.SetOutput(file_handler)
	http_return_status := strconv.FormatInt(int64(return_status), 10)
	response_length := strconv.FormatInt(int64(len(return_string)), 10)
	if request.URL.RawQuery != "" {
		log.Println(request.Host + " " + request.RemoteAddr + " \"" + request.Method + " " + request.URL.Path + "?" + request.URL.RawQuery + " " + request.Proto + "\"" + " " + http_return_status + " " + response_length + " " + "\"" + request.Header.Get("User-Agent") + "\"" + " " + time_taken_string)
	} else {
		log.Println(request.Host + " " + request.RemoteAddr + " \"" + request.Method + " " + request.URL.Path + " " + request.Proto + "\"" + " " + http_return_status + " " + response_length + " " + "\"" + request.Header.Get("User-Agent") + "\"" + " " + time_taken_string)

	}
}
