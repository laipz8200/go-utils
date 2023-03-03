package memcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	type args struct {
		key   string
		value any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				key:   "key 1",
				value: "value 1",
			},
		},
		{
			name: "case 2",
			args: args{
				key:   "key 2",
				value: "value 2",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.key, tt.args.value, 0)
			got, ok := Get(tt.args.key)
			asserts := assert.New(t)
			asserts.True(ok)
			asserts.Equal(tt.args.value, got)
		})
	}
}

func TestSetWithTTL(t *testing.T) {
	type args struct {
		key    string
		useKey string
		value  any
		ttl    time.Duration
	}
	tests := []struct {
		name   string
		args   args
		want   any
		wantOk bool
	}{
		{
			name: "case 1",
			args: args{
				key:    "key 1",
				useKey: "key 1",
				value:  "value 1",
				ttl:    1 * time.Second,
			},
			want:   "value 1",
			wantOk: true,
		},
		{
			name: "case 2",
			args: args{
				key:    "key 2",
				useKey: "key 2",
				value:  "value 2",
				ttl:    1 * time.Second,
			},
			want:   "value 2",
			wantOk: true,
		},
		{
			name: "outdated",
			args: args{
				key:    "key 1",
				useKey: "key 1",
				value:  "value 1",
				ttl:    -1 * time.Second,
			},
			want:   nil,
			wantOk: false,
		},
		{
			name: "no value",
			args: args{
				key:    "key 1",
				useKey: "key 2",
				value:  "value 1",
				ttl:    10 * time.Second,
			},
			want:   nil,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache = &memCache{
				store: map[string]item{},
			}
			Set(tt.args.key, tt.args.value, tt.args.ttl)
			got, ok := Get(tt.args.useKey)
			asserts := assert.New(t)
			asserts.Equal(tt.wantOk, ok)
			asserts.Equal(tt.want, got)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				key: "key 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Set(tt.args.key, "some value", 0)
			got, ok := Get(tt.args.key)

			asserts := assert.New(t)
			asserts.True(ok)
			asserts.NotEmpty(got)

			Remove(tt.args.key)
			got, ok = Get(tt.args.key)
			asserts.False(ok)
			asserts.Nil(got)
		})
	}
}
