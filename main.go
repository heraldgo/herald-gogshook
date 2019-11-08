package main

import (
	"fmt"

	"github.com/heraldgo/herald-gogshook/selector"
	"github.com/heraldgo/herald-gogshook/transformer"
	"github.com/heraldgo/herald-gogshook/trigger"
	"github.com/heraldgo/herald-gogshook/util"
)

const triggerGogsHookName = "gogs_hook"
const selectorGogsHookName = "gogs_hook"
const transformerGogsHookName = "gogs_hook"

// CreateTrigger create a new trigger
func CreateTrigger(name string, param map[string]interface{}) (interface{}, error) {
	if name != triggerGogsHookName {
		return nil, fmt.Errorf(`Trigger "%s" is not in plugin "gogshook"`, name)
	}

	unixSocket, _ := util.GetStringParam(param, "unix_socket")
	host, _ := util.GetStringParam(param, "host")
	port, _ := util.GetIntParam(param, "port")
	secret, _ := util.GetStringParam(param, "secret")

	if port == 0 && unixSocket == "" {
		host = "127.0.0.1"
		port = 8234
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
func CreateSelector(name string, param map[string]interface{}) (interface{}, error) {
	if name != selectorGogsHookName {
		return nil, fmt.Errorf(`Selector "%s" is not in plugin "gogshook"`, name)
	}

	return &selector.GogsHook{}, nil
}

// CreateTransformer create a new transformer
func CreateTransformer(name string, param map[string]interface{}) (interface{}, error) {
	if name != transformerGogsHookName {
		return nil, fmt.Errorf(`Transformer "%s" is not in plugin "gogshook"`, name)
	}

	return &transformer.GogsHook{}, nil
}
