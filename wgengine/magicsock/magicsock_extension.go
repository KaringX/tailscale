// Copyright (c) Tailscale Inc & contributors
// SPDX-License-Identifier: BSD-3-Clause

// Package magicsock implements a socket that can change its communication path while
// in use, actively searching for the best way to communicate.
package magicsock

import (
	"github.com/sagernet/wireguard-go/conn"
)

func (s *connBind) SendWithoutModify(bufs [][]byte, endpoint conn.Endpoint) error { //karing hiddify
	return s.Send(bufs, endpoint, 0)
}
