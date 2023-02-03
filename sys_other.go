//go:build !linux
// +build !linux

package keyctl

func keyctl_SetTimeout(id keyId, nsecs uint) error {
	return ErrNotImplemented
}

func keyctl_Read(id keyId, b *byte, size int) (int32, error) {
	return 0, ErrNotImplemented
}

func keyctl_Link(id, ring keyId) error {
	return ErrNotImplemented
}

func keyctl_Unlink(id, ring keyId) error {
	return ErrNotImplemented
}

func keyctl_Chown(id keyId, user, group int) error {
	return ErrNotImplemented
}

func keyctl_SetPerm(id keyId, perm uint32) error {
	return ErrNotImplemented
}

func add_key(keyType, keyDesc string, payload []byte, id int32) (int32, error) {
	return 0, ErrNotImplemented
}

func getfsgid() (int32, error) {
	return 0, ErrNotImplemented
}

func newKeyring(id keyId) (*keyring, error) {
	return nil, ErrNotImplemented
}

func createKeyring(parent keyId, name string) (*keyring, error) {
	return nil, ErrNotImplemented
}

func searchKeyring(id keyId, name, keyType string) (keyId, error) {
	return 0, ErrNotImplemented
}

func describeKeyId(id keyId) ([]byte, error) {
	return nil, ErrNotImplemented
}

func listKeys(id keyId) ([]keyId, error) {
	return nil, ErrNotImplemented
}

func updateKey(id keyId, payload []byte) error {
	return ErrNotImplemented
}
