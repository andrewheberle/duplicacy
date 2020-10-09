// Copyright (c) Acrosync LLC. All rights reserved.
// Free for personal use and commercial trial
// Commercial use requires per-user licenses available from https://duplicacy.com

// +build !windows

package duplicacy

import "os"

// OpenFileReadOnly wraps os.OpenFile to open a file read-only
func OpenFileReadOnly(fullPath string) (*os.File, error) {
	return os.OpenFile(fullPath, os.O_RDONLY, 0)
}
