package flog

import (
	"fylos/flog/internal"
	"testing"
)

func TestInfo(t *testing.T) {
	Info("123", 123, map[string]interface{}{"ok": true}, "ok")
}

func TestInfof(t *testing.T) {
	Infof("%s %d %v %s", "123", 123, map[string]interface{}{"ok": true}, "ok")
	internal.OK()
}
