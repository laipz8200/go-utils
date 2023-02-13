package memcache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_memcache(t *testing.T) {
	type args struct {
		key   string
		value any
		ttl   int64
	}
	tests := []struct {
		name   string
		args   args
		key    string
		want   any
		wantOk bool
	}{
		{
			name: "normal case",
			args: args{
				key:   "key 1",
				value: "value 1",
				ttl:   0,
			},
			key:    "key 1",
			want:   "value 1",
			wantOk: true,
		},
		{
			name: "not found",
			args: args{
				key:   "key 1",
				value: nil,
				ttl:   1,
			},
			key:    "key 2",
			want:   nil,
			wantOk: false,
		},
		{
			name: "outdate",
			args: args{
				key:   "key 1",
				value: "value 1",
				ttl:   -1,
			},
			key:    "key 1",
			want:   nil,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memcache{
				store: map[string]item{},
			}
			m.setWithTTL(tt.args.key, tt.args.value, tt.args.ttl)
			got, ok := m.get(tt.key)

			asserts := assert.New(t)
			asserts.Equal(tt.want, got)
			asserts.Equal(tt.wantOk, ok)
		})
	}
}

func Test_memcache_remove(t *testing.T) {
	type fields struct {
		store map[string]item
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   any
		wantOk bool
	}{
		{
			name: "normal case",
			fields: fields{
				store: map[string]item{
					"key1": {
						value: "value 1",
						ttl:   0,
					},
				},
			},
			args: args{
				key: "key1",
			},
			want:   nil,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memcache{
				store: tt.fields.store,
			}
			m.remove(tt.args.key)

			got, ok := m.get(tt.args.key)

			asserts := assert.New(t)
			asserts.Equal(tt.want, got)
			asserts.Equal(tt.wantOk, ok)
		})
	}
}
