package util

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToArgs(t *testing.T) {
	tests := []struct {
		name string
		args MappedArgs
		want []string
	}{
		{
			"basic",
			MappedArgs{
				"foo":   "bar",
				"bar":   "baf",
				"empty": "",
			},
			[]string{"foo=bar", "bar=baf", "empty="},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.ToArgs()
			sort.Strings(got)
			sort.Strings(tt.want)
			assert.Equal(t, tt.want, got)
		})
	}
}
