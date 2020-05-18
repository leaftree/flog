package flog

import (
	"testing"

	"github.com/leaftree/flog/internal"
)

func TestInfo(t *testing.T) {
	Info("123", 123, map[string]interface{}{"ok": true}, "ok")
}

func TestInfof(t *testing.T) {
	Infof("%s %d %v %s", "123", 123, map[string]interface{}{"ok": true}, "ok")
	internal.OK()
}
