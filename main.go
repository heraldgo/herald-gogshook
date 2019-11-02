package main

import (
	"fmt"

	"github.com/heraldgo/herald-gogshook/filter"
	"github.com/heraldgo/herald-gogshook/trigger"
	"github.com/heraldgo/herald-gogshook/util"
)

const triggerGogsHookName = "gogs_hook"
const filterGogsHookName = "gogs_hook"

// CreateTrigger create a new trigger
func CreateTrigger(name string, param map[string]interface{}) (interface{}, error) {
	if name != triggerGogsHookName {
		return nil, fmt.Errorf(`Trigger "%s" is not in plugin "gogshook"`, name)
	}

	unixSocket, _ := util.GetStringParam(param, "unix_socket")
	host, _ := util.GetStringParam(param, "host")
	port, _ := util.GetIntParam(param, "port")
	secret, _ := util.GetStringParam(param, "secret")

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

// CreateFilter create a new filter
func CreateFilter(name string, param map[string]interface{}) (interface{}, error) {
	if name != filterGogsHookName {
		return nil, fmt.Errorf(`Filter "%s" is not in plugin "gogshook"`, name)
	}

	return &filter.GogsHook{}, nil
}
