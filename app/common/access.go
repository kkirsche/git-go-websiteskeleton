// Note: inspiration for this from https://gist.github.com/cespare/3985516
package common

import (
    "time"
    "strings"
    "strconv"
    "net/http"

    "github.com/golang/glog"
)

const (
    logFile = "access_log.txt"
)

type accessLog struct {
    ip, method, uri, protocol, host     string
    elapsedTime                         time.Duration
}

// LogAccess is used to log HTTP requests made to a golang web server, record information about the request, then allow it to continue.
func LogAccess(w http.ResponseWriter, req *http.Request, duration time.Duration) {
    clientIP := req.RemoteAddr

    if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
        clientIP = clientIP[:colon]
    }

    record := &accessLog{
        ip:             clientIP,
        method:         req.Method,
        uri:            req.RequestURI,
        protocol:       req.Proto,
        host:           req.Host,
        elapsedTime:    duration,
    }

    writeAccessLog(record)
}

func writeAccessLog(r *accessLog) {
    logRecord := fmt.Sprintf("%s %s %s %s: %s host: %s (load time: %d seconds)", r.ip, r.protocol, r.method, r.uri, r.host, strconv.FormatFloat(r.elapsedTime.Seconds(), 'f', 5, 64))
    glog.Infoln(logRecord)
    glog.Flush()
}
