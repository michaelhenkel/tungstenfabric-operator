// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build arm64 amd64 ppc64 ppc64le mips64 mips64le s390x
<<<<<<< HEAD
// +build darwin dragonfly freebsd linux netbsd openbsd
=======
// +build aix darwin dragonfly freebsd linux netbsd openbsd
>>>>>>> v0.0.4

package socket

import "unsafe"

func (v *iovec) set(b []byte) {
	l := len(b)
	if l == 0 {
		return
	}
	v.Base = (*byte)(unsafe.Pointer(&b[0]))
	v.Len = uint64(l)
}
