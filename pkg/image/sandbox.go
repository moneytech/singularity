/*
  Copyright (c) 2018, Sylabs, Inc. All rights reserved.

  This software is licensed under a 3-clause BSD license.  Please
  consult LICENSE file distributed with the sources of this project regarding
  your rights to use or distribute this software.
*/

package image

import (
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type Sandbox struct {
	rootfs string
}

// SandboxFromPath returns a sandbox object of the directory located at path
func SandboxFromPath(path string) *Sandbox {
	return &Sandbox{}
}

/* RuntimeImage Interface Methods */

// Root returns the OCI specs.Root data type
func (i *Sandbox) Root() *specs.Root {
	return &specs.Root{}
}

/* BuildtimeImage Interface Methods */

// Rootfs returns the path of the rootfs of the sandbox
func (i *Sandbox) Rootfs() string {
	return i.rootfs
}

// isSandbox checks the "magic" of the given file and
// determines if the file is of sandbox type
func isSandbox(path string) bool {
	return false
}
