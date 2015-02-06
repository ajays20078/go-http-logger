package httpLogger

import (
	"log"
	"net/http"
	"os"
	"strconv"
)

// LogLine functions logs  the http request info and returnStatus/Time to fileHandler
func LogLine(request *http.Request, returnStatus int, returnString string, fileHandler *os.File, startTime int64, endTime int64) int {
	go logLineAsync(request, returnStatus, returnString, fileHandler, startTime, endTime)
	return 0
}

func logLineAsync(request *http.Request, returnStatus int, returnString string, fileHandler *os.File, startTime int64, endTime int64) {
	timeDiff := (float64(endTime) - float64(startTime)) / 1000000.0
	timeTakenString := strconv.FormatFloat(timeDiff, 'f', 3, 64)
	log.SetOutput(fileHandler)
	httpReturnStatus := strconv.FormatInt(int64(returnStatus), 10)
	responseLength := strconv.FormatInt(int64(len(returnString)), 10)
	if request.URL.RawQuery != "" {
		log.Println(request.Host + " " + request.RemoteAddr + " \"" + request.Method + " " + request.URL.Path + "?" + request.URL.RawQuery + " " + request.Proto + "\"" + " " + httpReturnStatus + " " + responseLength + " " + "\"" + request.Header.Get("User-Agent") + "\"" + " " + timeTakenString)
	} else {
		log.Println(request.Host + " " + request.RemoteAddr + " \"" + request.Method + " " + request.URL.Path + " " + request.Proto + "\"" + " " + httpReturnStatus + " " + responseLength + " " + "\"" + request.Header.Get("User-Agent") + "\"" + " " + timeTakenString)
	}
}
