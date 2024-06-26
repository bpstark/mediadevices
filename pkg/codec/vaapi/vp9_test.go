//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris
// +build dragonfly freebsd linux netbsd openbsd solaris

package vaapi

import (
	"testing"

	"github.com/carbonrobotics/mediadevices/pkg/codec"
)

func TestVP9ShouldImplementBitRateControl(t *testing.T) {
	t.SkipNow() // TODO: Implement bit rate control

	e := &encoderVP9{}
	if _, ok := e.Controller().(codec.BitRateController); !ok {
		t.Error()
	}
}

func TestVP9ShouldImplementKeyFrameControl(t *testing.T) {
	e := &encoderVP9{}
	if _, ok := e.Controller().(codec.KeyFrameController); !ok {
		t.Error()
	}
}
