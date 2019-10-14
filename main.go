package main

import (
	"fmt"

	"github.com/xianghuzhao/herald-gogshook/filter"
	"github.com/xianghuzhao/herald-gogshook/trigger"
)

const triggerGogsHookName = "gogs_hook"
const filterGogsHookName = "gogs_hook"

// CreateTrigger create a new trigger
func CreateTrigger(name string) (interface{}, error) {
	if name != triggerGogsHookName {
		return nil, fmt.Errorf("Trigger \"%s\" not found", name)
	}

	return &trigger.GogsHook{}, nil
}

// CreateFilter create a new filter
func CreateFilter(name string) (interface{}, error) {
	if name != filterGogsHookName {
		return nil, fmt.Errorf("Filter \"%s\" not found", name)
	}

	return &filter.GogsHook{}, nil
}
