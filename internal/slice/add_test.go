package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"toys4go/internal/errs"
)

func TestAdd1(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		addVal    int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "index 0",
			slice:     []int{111, 100},
			addVal:    222,
			index:     0,
			wantSlice: []int{222, 111, 100},
		},
		{
			name:      "index middle",
			slice:     []int{111, 112, 113},
			addVal:    222,
			index:     1,
			wantSlice: []int{111, 222, 112, 113},
		},
		{
			name:    "index out of range",
			slice:   []int{111, 100},
			index:   12,
			wantErr: errs.ErrorIndexOutOfRange(2, 12),
		},
		{
			name:    "index less than 0",
			slice:   []int{111, 100},
			index:   -1,
			wantErr: errs.ErrorIndexOutOfRange(2, -1),
		},
		{
			name:      "index last",
			slice:     []int{111, 100, 101, 102, 102, 102},
			addVal:    222,
			index:     5,
			wantSlice: []int{111, 100, 101, 102, 102, 222, 102},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Add(tc.slice, tc.addVal, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
