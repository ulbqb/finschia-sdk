package slessdb

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNodeHash(t *testing.T) {
	testCases := []struct {
		name  string
		key   string
		ok    bool
		store []byte
		hash  []byte
	}{
		{
			name:  "valid",
			key:   "s/k:store/nhashxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			ok:    true,
			store: []byte("store"),
			hash:  []byte("hashxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
		},
		{
			name:  "empty",
			key:   "",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "not enough length key",
			key:   "s/k:/",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "wrong store prefix",
			key:   "xxxx/n",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "no slash",
			key:   "s/k:xxxxx",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "no node prefix",
			key:   "s/k:store/",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "wrong node prefix",
			key:   "s/k:store/x",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "no node hash",
			key:   "s/k:store/n",
			ok:    false,
			store: nil,
			hash:  nil,
		},
		{
			name:  "not enough length node hash",
			key:   "s/k:store/nnotenouglength",
			ok:    false,
			store: nil,
			hash:  nil,
		},
	}

	for _, tc := range testCases {
		ok, store, hash := nodeHash([]byte(tc.key))
		require.Equal(t, tc.ok, ok)
		require.Equal(t, []byte(tc.store), store)
		require.Equal(t, []byte(tc.hash), hash)
	}
}
