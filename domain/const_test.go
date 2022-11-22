package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURankIota(t *testing.T) {
	tests := []struct {
		UR    URank
		Index URank
		OK    bool
	}{
		{URContributor, 2, true},
		{URSubscriber, 3, true},
		{URSuperAdmin, 0, true},
		{URNomad, 4, true},
		{URAdmin, 1, true},

		{URContributor, 7, false},
		{URAdmin, 0, false},
		{URAdmin, 11, false},
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("[%v]", tc), func(t *testing.T) {
			var r URank = tc.UR
			if tc.OK {
				assert.Equal(t, r, tc.Index)
			} else {
				assert.NotEqual(t, r, tc.Index)
			}
		})
	}
}
