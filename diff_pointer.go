package diff

import "reflect"

func (cl *Changelog) diffPtr(path []string, a, b reflect.Value) error {
	if a.Kind() != b.Kind() {
		return ErrTypeMismatch
	}

	if a.IsNil() && b.IsNil() {
		return nil
	}

	if a.IsNil() {
		cl.add(UPDATE, path, nil, b.Interface())
		return nil
	}

	if b.IsNil() {
		cl.add(UPDATE, path, a.Interface(), nil)
		return nil
	}

	return cl.diff(path, reflect.Indirect(a), reflect.Indirect(b))
}