package main

import (
	"fmt"

	"github.com/heraldgo/herald-gogshook/selector"
	"github.com/heraldgo/herald-gogshook/trigger"
	"github.com/heraldgo/herald-gogshook/util"
)

const triggerGogsHookName = "gogs_hook"
const selectorGogsHookName = "gogs_hook"
const transformerGogsHookName = "gogs_hook"

// CreateTrigger create a new trigger
func CreateTrigger(typeName string, param map[string]interface{}) (interface{}, error) {
	if typeName != triggerGogsHookName {
		return nil, fmt.Errorf(`Trigger "%s" is not in plugin "gogshook"`, typeName)
	}

	unixSocket, _ := util.GetStringParam(param, "unix_socket")
	host, _ := util.GetStringParam(param, "host")
	port, _ := util.GetIntParam(param, "port")
	secret, _ := util.GetStringParam(param, "secret")

	if port == 0 && unixSocket == "" {
		port = 8234
	}

	if port != 0 && host == "" {
		host = "127.0.0.1"
	}

	tgr := &trigger.GogsHook{
		HTTPServer: util.HTTPServer{
			UnixSocket: unixSocket,
			Host:       host,
			Port:       port,
		},
		Secret: secret,
	}
	return tgr, nil
}

// CreateSelector create a new selector
func CreateSelector(typeName string, param map[string]interface{}) (interface{}, error) {
	if typeName != selectorGogsHookName {
		return nil, fmt.Errorf(`Selector "%s" is not in plugin "gogshook"`, typeName)
	}

	return &selector.GogsHook{}, nil
}
