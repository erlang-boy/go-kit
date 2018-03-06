package binder

import (
	. "shell/config"
	"time"
)

const BindTimeout = 5 * time.Second

func binderFac(spec *BinderSpec) Binder {
	if spec.BindServerName == HDBD {
		return NewHdbd(spec)
	} else {
		return NewHEngined(spec)
	}
}
