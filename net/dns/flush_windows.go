// Copyright (c) 2021 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dns

import (
	"fmt"
	"os/exec"
)

func flushCaches() error {
	out, err := exec.Command("ipconfig", "/flushdns").CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v (output: %s)", err, out)
	}
	return nil
}

// Flush clears the local resolver cache.
//
// Only Windows has a public dns.Flush, needed in router_windows.go. Other
// platforms like Linux need a different flush implementation depending on
// the Dns Manager. There is a FlushCaches method on the Manager which
// can be used on all platforms.
func Flush() error {
	return flushCaches()
}
