package transformer

import (
	"github.com/heraldgo/herald-gogshook/util"
)

// GogsHook is a transformer generate simple gogs param
type GogsHook struct {
	util.BaseLogger
}

// Transform will generate simple gogs param
func (tfm *GogsHook) Transform(triggerParam map[string]interface{}) map[string]interface{} {
	return util.GogsParam(triggerParam)
}
