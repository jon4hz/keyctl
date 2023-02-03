//go:build linux && 386
// +build linux,386

package keyctl

const (
	syscall_keyctl   uintptr = 288
	syscall_add_key  uintptr = 286
	syscall_setfsgid uintptr = 139
)
