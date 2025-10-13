package log

import (
	"log"
	"os"
	"net/http"
)

var writer = os.Stdout
var logger = log.New(writer, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)

const (
	PREFIX_INFO = "[INFO] "
	PREFIX_DEBUG = "[DEBUG] "
	PREFIX_ERROR = "[ERROR] "
	PREFIX_WARN = "[WARN] "
)

const HTTP_REQUEST_LOG_FORMAT = "{ url: %s, method: %s, host: %s, headers: %s }"

func Info(format string, v ...any) {
	updFormat := PREFIX_INFO + format
	logger.Printf(updFormat, v...)
}

func Debug(format string, v ...any) {
	updFormat := PREFIX_DEBUG + format
	logger.Printf(updFormat, v...)
}

func Error(format string, v ...any) {
	updFormat := PREFIX_ERROR + format
	logger.Printf(updFormat, v...)
}

func Warn(format string, v ...any) {
	updFormat := PREFIX_WARN + format
	logger.Printf(updFormat, v...)
}

func HttpRequest(req *http.Request) {
	Info(HTTP_REQUEST_LOG_FORMAT, req.URL, req.Method, req.Host, req.Header)
}
