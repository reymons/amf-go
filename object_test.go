package amf

import "testing"

func TestObject_ReturnsCorrectValuesForUndefinedProp(t *testing.T) {
	obj := NewObject()
	prop := obj.Prop("some_random_prop_name123")
	value := prop.Value()

	if prop.key != blankPropKey {
		t.Errorf("invalid key: expected %s, got %s", blankPropKey, prop.key)
	}
	if value.value != nil {
		t.Errorf("invalid value: expected nil, got %v", value.value)
	}
}
