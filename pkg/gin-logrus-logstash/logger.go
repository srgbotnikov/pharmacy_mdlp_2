package ginlogruslogstash

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"strings"
	"time"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var timeFormat = "02/Jan/2006:15:04:05"

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// Logger is the logrus logger handler
func Logger(logger *logrus.Logger, connectionString, appName string) gin.HandlerFunc {

	conn, err := net.Dial("udp", connectionString)
	if err != nil {
		log.Println(err)
	}

	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": appName}))

	logger.Hooks.Add(hook)

	return func(c *gin.Context) {
		// other handler can change c.Path so:
		path := c.Request.URL.Path
		start := time.Now()

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		var buf bytes.Buffer
		tee := io.TeeReader(c.Request.Body, &buf)
		request, _ := ioutil.ReadAll(tee)
		c.Request.Body = ioutil.NopCloser(&buf)

		//	codeParus := c.Request.Header.Get("CodeParus")

		c.Next()
		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()

		bodyResp := blw.body.String()

		clientIP := c.ClientIP()

		dataLength := c.Writer.Size()

		if dataLength < 0 {
			dataLength = 0
		}

		if statusCode < http.StatusBadRequest {
			bodyResp = ""
		}

		requsetStr := string(request)
		if strings.Contains(path, "/v2/documents/") {
			requsetStr = ""
		}

		entry := logger.WithFields(logrus.Fields{
			"statusCode": statusCode,
			"latency":    latency, // time to process
			"clientIP":   clientIP,
			"method":     c.Request.Method,
			"path":       path,
			"dataLength": dataLength,
			"response":   bodyResp,
			"request":    requsetStr,
			//"codeParus":  codeParus,
			"RequestURI": c.Request.RequestURI,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			message := fmt.Sprintf("%s - [%s] \"%s %s\" %d (%dms)", clientIP, time.Now().Format(timeFormat), c.Request.Method, path, statusCode, latency)
			if statusCode >= http.StatusInternalServerError {
				entry.Error(message)
			} else if statusCode >= http.StatusBadRequest {
				entry.Warn(message)
			} else {
				entry.Info(message)
			}
		}
	}
}
