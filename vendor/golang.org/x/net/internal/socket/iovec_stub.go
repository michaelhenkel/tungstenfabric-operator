// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// +build !darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris
=======
// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!solaris
>>>>>>> v0.0.4

package socket

type iovec struct{}

func (v *iovec) set(b []byte) {}
