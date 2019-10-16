package main

import (
	"fmt"

	"github.com/heraldgo/herald-gogshook/filter"
	"github.com/heraldgo/herald-gogshook/trigger"
)

const triggerGogsHookName = "gogs_hook"
const filterGogsHookName = "gogs_hook"

// CreateTrigger create a new trigger
func CreateTrigger(name string) (interface{}, error) {
	if name != triggerGogsHookName {
		return nil, fmt.Errorf("Trigger \"%s\" is not in plugin \"gogshook\"", name)
	}

	return &trigger.GogsHook{}, nil
}

// CreateFilter create a new filter
func CreateFilter(name string) (interface{}, error) {
	if name != filterGogsHookName {
		return nil, fmt.Errorf("Filter \"%s\" is not in plugin \"gogshook\"", name)
	}

	return &filter.GogsHook{}, nil
}
