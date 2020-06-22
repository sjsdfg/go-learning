package chap01

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestSomething(t *testing.T) {

	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)
	var nilObject reflect.Type = nil
	assert.Nil(t, nilObject)

	// assert for not nil (good when you expect something)
	var nonNilObject string = "Something"
	if assert.NotNil(t, nonNilObject) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "Something", nonNilObject)
	}

}
