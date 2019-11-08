package util

import (
	"strings"
)

// GogsParam will choose some important params
func GogsParam(param map[string]interface{}) map[string]interface{} {
	event, _ := GetStringParam(param, "event")

	payloadParam, _ := GetMapParam(param, "payload")
	repositoryParam, _ := GetMapParam(payloadParam, "repository")

	repoName, _ := GetStringParam(repositoryParam, "name")
	repoFullName, _ := GetStringParam(repositoryParam, "full_name")
	cloneURL, _ := GetStringParam(repositoryParam, "clone_url")

	ref, _ := payloadParam["ref"]
	repoRef, _ := ref.(string)
	refFrag := strings.Split(repoRef, "/")
	branch := refFrag[len(refFrag)-1]

	return map[string]interface{}{
		"event":     event,
		"name":      repoName,
		"full_name": repoFullName,
		"branch":    branch,
		"clone_url": cloneURL,
	}
}
