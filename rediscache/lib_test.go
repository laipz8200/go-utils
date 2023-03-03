package rediscache

import (
	"context"
	"encoding"
	"encoding/json"
	"testing"
	"time"

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

func Test_RedisCache(t *testing.T) {
	var (
		addr     = "127.0.0.1:32768"
		password = "redispw"
		db       = 0
	)

	Init(addr, password, db)

	tests := []struct {
		name      string
		key       string
		value     any
		ttl       time.Duration
		want      string
		wantError bool
	}{
		{
			name:      "case 1",
			key:       "key 1",
			value:     "value 1",
			ttl:       0,
			want:      "value 1",
			wantError: false,
		},
		{
			name:      "case 2",
			key:       "key 2",
			value:     "value 2",
			ttl:       0,
			want:      "value 2",
			wantError: false,
		},
		{
			name: "Object",
			key:  "key 4",
			value: Foo{
				Name: "name",
			},
			ttl:       0,
			want:      "{\"Name\":\"name\"}",
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			asserts := assert.New(t)

			err := Set(tt.key, tt.value, tt.ttl)
			asserts.NoError(err)

			got, err := Get(tt.key)
			if tt.wantError {
				asserts.Error(err)
				return
			}
			asserts.Equal(tt.want, got)

			err = Remove(tt.key)
			asserts.NoError(err)

			_, err = Get(tt.key)
			asserts.Error(err)
		})
	}
}

func Test_WithContext(t *testing.T) {
	InitWithDSN("redis://default:redispw@localhost:32768")

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := WithContext(ctx).Set("key", "value", 0)
	assert.Error(t, err)
}
