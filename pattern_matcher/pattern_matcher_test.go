package pattern_matcher

import (
	"testing"
)

func TestPatternMatcher(t *testing.T) {
	type fields struct {
		include_patterns []string
		exclude_patterns []string
	}
	tests := []struct {
		name   string
		fields fields
		value  string
		want   bool
	}{
		{
			"test_1",
			fields{
				include_patterns: []string{"*"},
				exclude_patterns: []string{},
			},
			"hello_world",
			true,
		},
		{
			"test_2",
			fields{
				include_patterns: []string{"*"},
				exclude_patterns: []string{"hello*"},
			},
			"hello_world",
			false,
		},
		{
			"test_3",
			fields{
				include_patterns: []string{"*"},
				exclude_patterns: []string{"hello*"},
			},
			"hi_world",
			true,
		},
		{
			"test_4",
			fields{
				include_patterns: []string{"hello"},
				exclude_patterns: []string{"*"},
			},
			"hello",
			false,
		},
		{
			"test_5",
			fields{
				include_patterns: []string{"openstack_nova*", "openstack_cinder*"},
				exclude_patterns: []string{"openstack_cinder_attachment*"},
			},
			"openstack_cinder_attachments",
			false,
		},
		{
			"test_6",
			fields{
				include_patterns: []string{"openstack*"},
				exclude_patterns: []string{"openstack_identity*"},
			},
			"openstack_cinder_attachments",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := New(
				WithReplaceIncludes(tt.fields.include_patterns),
				WithReplaceExcludes(tt.fields.exclude_patterns),
			)
			if got := pm.Match(tt.value); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
