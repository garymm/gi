// Copyright 2018 The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// based on golang.org/x/exp/shiny:
// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !dragonfly && !openbsd && !3d
// +build !linux,!dragonfly,!openbsd,!3d

package x11driver

import (
	"fmt"
	"runtime"
	"unsafe"
)

func shmOpen(size int) (shmid uintptr, addr unsafe.Pointer, err error) {
	return 0, unsafe.Pointer(uintptr(0)),
		fmt.Errorf("unsupported GOOS/GOARCH %s/%s", runtime.GOOS, runtime.GOARCH)
}

func shmClose(p unsafe.Pointer) error {
	return fmt.Errorf("unsupported GOOS/GOARCH %s/%s", runtime.GOOS, runtime.GOARCH)
}
