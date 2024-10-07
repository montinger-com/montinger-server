package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/montinger-com/montinger-server/config"
	"github.com/montinger-com/montinger-server/lib/exceptions"
	"github.com/rashintha/logger"
)

const (
	Origin  byte = 0x0
	Replace      = 0x1
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	status byte
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	if r.status == 0x1 {
		r.body.Write(b)
		return r.ResponseWriter.Write(b)
	} else {
		return r.body.Write(b) //r.ResponseWriter.Write(b)
	}
}

func Interceptor() gin.HandlerFunc {
	return func(c *gin.Context) {

		if strings.HasPrefix(c.Request.URL.Path, "/metrics") {
			c.Next()
			return
		}

		wb := &responseBodyWriter{
			body:           &bytes.Buffer{},
			ResponseWriter: c.Writer,
			status:         Origin,
		}
		c.Writer = wb

		c.Next()
		wb.status = Replace
		originBytes := wb.body

		var result response_model.Result
		resp := c.Writer
		contentType := resp.Header().Get("Content-Type")

		if strings.HasPrefix(contentType, "application/json") {
			err := json.Unmarshal([]byte(originBytes.String()), &result)
			if err != nil {
				var jsonData any
				err := json.Unmarshal([]byte(originBytes.String()), &jsonData)
				if err != nil {
					result = response_model.Result{
						Data: originBytes.String(),
					}
				}
				result = response_model.Result{
					Data: jsonData,
				}

			}
		} else if strings.HasPrefix(contentType, "text/plain") {
			result = response_model.Result{
				Data: originBytes.String(),
			}
		} else {
			result = response_model.Result{
				Message: exceptions.ResourceNotFound.Error(),
				Errors:  []string{exceptions.ResourceNotFound.Error()},
			}
		}

		//omit fields from result which is not public

		result.Data = removeProtectedFields(result.Data, []string{
			"password",
		})

		result.Timestamp = time.Now().UTC().String()
		result.Environment = config.APP_ENV
		result.RequestID = requestid.Get(c)
		result.Path = c.Request.RequestURI
		result.Status = wb.Status()

		wb.body = &bytes.Buffer{}
		c.JSON(wb.Status(), result)
	}
}

func removeProtectedFields(data interface{}, fieldsToRemove []string) interface{} {

	// Convert the data to a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		return data
	}

	// Unmarshal the JSON string into a map[string]interface{}`
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(jsonData, &jsonMap); err != nil {
		return data
	}

	// Remove the fields from the map
	for _, fieldToRemove := range fieldsToRemove {
		delete(jsonMap, fieldToRemove)
	}

	// Marshal the map back to JSON
	jsonData, err = json.Marshal(jsonMap)
	if err != nil {
		return data
	}

	// Unmarshal the JSON back into the original data type
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data
	}

	return data

}

func Log() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		logger.Defaultln(fmt.Sprintf("ip: %s - time: %s | method: %s | path: %s | proto: %s | status: %d | latency: %s | userAgent: %s | error: %s",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage))

		return ""
	})
}
