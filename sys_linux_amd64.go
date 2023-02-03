//go:build linux && amd64
// +build linux,amd64

package keyctl

const (
	syscall_keyctl   uintptr = 250
	syscall_add_key  uintptr = 248
	syscall_setfsgid uintptr = 123
)
