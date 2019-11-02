package trigger

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/heraldgo/herald-gogshook/util"
)

// GogsHook is a trigger which will listen to http request
type GogsHook struct {
	util.HTTPServer
	Secret string
}

func (tgr *GogsHook) validateGogs(r *http.Request, body []byte) error {
	if r.Method != "POST" {
		return fmt.Errorf("Only POST request allowed")
	}

	sigHeader := r.Header.Get("X-Gogs-Signature")
	signature, err := hex.DecodeString(sigHeader)
	if err != nil {
		return fmt.Errorf("Invalid X-Gogs-Signature: %s", sigHeader)
	}
	key := []byte(tgr.Secret)

	if !util.ValidateMAC(body, signature, key) {
		return fmt.Errorf("Signature validation Error")
	}
	return nil
}

// Run the GogsHook trigger
func (tgr *GogsHook) Run(ctx context.Context, param chan map[string]interface{}) {
	tgr.ValidateFunc = tgr.validateGogs

	requestChan := make(chan map[string]interface{})

	tgr.ProcessFunc = func(w http.ResponseWriter, r *http.Request, body []byte) {
		bodyMap, err := util.JSONToMap(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Request body error: %s", err)))
			return
		}

		result := make(map[string]interface{})
		result["event"] = r.Header.Get("X-Gogs-Event")
		result["payload"] = bodyMap

		requestChan <- result
		w.Write([]byte("Gogs param received and trigger activated\n"))
	}

	tgr.Start()
	defer tgr.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case reqParam := <-requestChan:
			param <- reqParam
		}
	}
}
