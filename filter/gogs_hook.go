package filter

import (
	"net/url"
	"strings"

	"github.com/heraldgo/herald-gogshook/util"
)

// GogsHook is a filter only pass with specified repo and branch
type GogsHook struct {
	util.BaseLogger
}

// Filter will only pass with specified repo and branch
func (flt *GogsHook) Filter(triggerParam, filterParam map[string]interface{}) (map[string]interface{}, bool) {
	event, _ := util.GetStringParam(triggerParam, "event")
	gogsEvent, _ := util.GetStringParam(filterParam, "gogs_event")
	if gogsEvent != "" && gogsEvent != event {
		flt.Debugf(`[Filter(GogsHook)] Event does not match: "%s"`, event)
		return nil, false
	}

	payloadParam, _ := util.GetMapParam(triggerParam, "payload")
	repositoryParam, _ := util.GetMapParam(payloadParam, "repository")

	repoName, _ := util.GetStringParam(repositoryParam, "name")
	repoFullName, _ := util.GetStringParam(repositoryParam, "full_name")
	cloneURL, _ := util.GetStringParam(repositoryParam, "clone_url")

	ref, _ := payloadParam["ref"]
	repoRef, _ := ref.(string)
	refFrag := strings.Split(repoRef, "/")
	branch := refFrag[len(refFrag)-1]

	repoParsed, err := url.Parse(cloneURL)
	if err != nil {
		flt.Errorf(`[Filter(GogsHook)] Invalid clone url: "%s"`, cloneURL)
		return nil, false
	}
	host := strings.SplitN(repoParsed.Host, ":", 2)[0]

	gogsHost, _ := util.GetStringParam(filterParam, "gogs_host")
	gogsName, _ := util.GetStringParam(filterParam, "gogs_name")
	gogsBranch, _ := util.GetStringParam(filterParam, "gogs_branch")

	if gogsHost != "" && gogsHost != host {
		flt.Debugf(`[Filter(GogsHook)] Host does not match: "%s"`, host)
		return nil, false
	}

	if gogsBranch != "" && gogsBranch != branch {
		flt.Debugf(`[Filter(GogsHook)] Branch does not match: "%s"`, branch)
		return nil, false
	}

	if gogsName != "" {
		if strings.ContainsAny(gogsName, "/") {
			if gogsName != repoFullName {
				flt.Debugf(`[Filter(GogsHook)] Name does not match: "%s"`, repoFullName)
				return nil, false
			}
		} else {
			if gogsName != repoName {
				flt.Debugf(`[Filter(GogsHook)] Full name does not match: "%s"`, repoName)
				return nil, false
			}
		}
	}

	result := map[string]interface{}{
		"event":     branch,
		"clone_url": cloneURL,
		"branch":    branch,
	}

	return result, true
}
