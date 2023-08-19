package application

import (
	"net/http"

)

func (m *MessageProcessor) processLogMethod(method string, rw http.ResponseWriter, _ *http.Request, window Window, params QueryParams) {
	switch method {
	case "Log":
		/*var msg logger.Message
		err := params.ToStruct(&msg)
		if err != nil {
			m.httpError(rw, "error parsing log message: %s", err.Error())
			return
		}
		msg.Sender = window.Name()
		globalApplication.Log(&msg)
		m.ok(rw)
		*/
	default:
		m.httpError(rw, "Unknown log method: %s", method)
	}

}
