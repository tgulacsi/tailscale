// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by tailscale.com/cmd/cloner -type Persist; DO NOT EDIT.

package controlclient

import ()

// Clone makes a deep copy of Persist.
// The result aliases no memory with the original.
func (src *Persist) Clone() *Persist {
	if src == nil {
		return nil
	}
	dst := new(Persist)
	*dst = *src
	return dst
}
