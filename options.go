package diff

// SliceOrdering determines whether the ordering of items in a slice results in a change
func SliceOrdering(enabled bool) func(d *Differ) error {
	return func(d *Differ) error {
		d.SliceOrdering = enabled
		return nil
	}
}

// DisableStructValues disables populating a seperate change for each item in a struct,
// where the struct is being compared to a nil value
func DisableStructValues() func(d *Differ) error {
	return func(d *Differ) error {
		d.DisableStructValues = true
		return nil
	}
}

// SetMapIdentifier sets an optional identifier for maps
// Maps containing the specified key will be compared using the value at this key
// If the values are inequal, the differ appends a DELETE of the first map and a CREATE of the second map
// to the changelog, as the whole map-object hast changed
func SetMapIdentifier(identifierKey string) func(d *Differ) error {
	return func(d *Differ) error {
		d.MapIdentifierKey = identifierKey
		return nil
	}
}
