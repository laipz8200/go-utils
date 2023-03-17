package rediscache

import (
	"context"
	"encoding"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ encoding.BinaryMarshaler = Foo{}

type Foo struct {
	Name string
}

// MarshalBinary implements encoding.BinaryMarshaler
func (f Foo) MarshalBinary() (data []byte, err error) {
	return json.Marshal(f)
}

func Test_WithContext(t *testing.T) {
	InitWithDSN("redis://default:redispw@localhost:32768")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := WithContext(ctx).Set("key", "value", 0)
	assert.Error(t, err)
}
