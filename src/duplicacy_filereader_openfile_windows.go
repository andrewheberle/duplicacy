// Copyright (c) Acrosync LLC. All rights reserved.
// Free for personal use and commercial trial
// Commercial use requires per-user licenses available from https://duplicacy.com

package duplicacy

import (
	"os"
	"syscall"
)

// OpenFileReadOnly attempts to open a file read-only using the Win32 file backup API then falls
// back to a standard os.OpenFile
func OpenFileReadOnly(fullPath string) (*os.File, error) {
	pathp, err := syscall.UTF16PtrFromString(fullPath)
	if err == nil {
		access := syscall.GENERIC_READ
		sharemode := uint32(syscall.FILE_SHARE_READ | syscall.FILE_SHARE_WRITE)
		createmode := syscall.OPEN_EXISTING
		f, err := syscall.CreateFile(pathp, access, sharemode, nil, createmode, syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OPEN_REPARSE_POINT, 0)
		if err == nil {
			return f, err
		}
	}

	return os.OpenFile(fullPath, os.O_RDONLY, 0)
}
