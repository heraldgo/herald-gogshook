package selector

import (
	"net/url"
	"strings"

	"github.com/heraldgo/herald-gogshook/util"
)

// GogsHook is a selector only pass with specified repo and branch
type GogsHook struct {
	util.BaseLogger
}

// Select will only pass with specified repo
func (slt *GogsHook) Select(triggerParam, jobParam map[string]interface{}) bool {
	mainParam := util.GogsParam(triggerParam)

	event := mainParam["event"]
	gogsEvent, _ := util.GetStringParam(jobParam, "gogs_event")
	if gogsEvent != "" && gogsEvent != event {
		slt.Debugf(`Event does not match: "%s"`, event)
		return false
	}

	repoParsed, err := url.Parse(mainParam["clone_url"].(string))
	if err != nil {
		slt.Errorf(`Invalid clone url: "%s"`, mainParam["clone_url"])
		return false
	}
	host := strings.SplitN(repoParsed.Host, ":", 2)[0]

	gogsHost, _ := util.GetStringParam(jobParam, "gogs_host")
	gogsName, _ := util.GetStringParam(jobParam, "gogs_name")
	gogsBranch, _ := util.GetStringParam(jobParam, "gogs_branch")

	if gogsHost != "" && gogsHost != host {
		slt.Debugf(`Host does not match: "%s"`, host)
		return false
	}

	branch := mainParam["branch"]
	if gogsBranch != "" && gogsBranch != branch {
		slt.Debugf(`Branch does not match: "%s"`, branch)
		return false
	}

	if gogsName != "" {
		if strings.ContainsAny(gogsName, "/") {
			repoFullName := mainParam["full_name"]
			if gogsName != repoFullName {
				slt.Debugf(`Name does not match: "%s"`, repoFullName)
				return false
			}
		} else {
			repoName := mainParam["name"]
			if gogsName != repoName {
				slt.Debugf(`Full name does not match: "%s"`, repoName)
				return false
			}
		}
	}

	return true
}
