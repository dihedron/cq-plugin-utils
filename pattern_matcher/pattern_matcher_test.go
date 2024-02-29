package pattern_matcher

import (
	"testing"

	"github.com/gobwas/glob"
)

func TestPatternMatcher(t *testing.T) {
	type fields struct {
		include_patterns []glob.Glob
		exclude_patterns []glob.Glob
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
				include_patterns: []glob.Glob{glob.MustCompile("*")},
				exclude_patterns: []glob.Glob{},
			},
			"hello_world",
			true,
		},
		{
			"test_2",
			fields{
				include_patterns: []glob.Glob{glob.MustCompile("*")},
				exclude_patterns: []glob.Glob{glob.MustCompile("hello*")},
			},
			"hello_world",
			false,
		},
		{
			"test_3",
			fields{
				include_patterns: []glob.Glob{glob.MustCompile("*")},
				exclude_patterns: []glob.Glob{glob.MustCompile("hello*")},
			},
			"hi_world",
			true,
		},
		{
			"test_4",
			fields{
				include_patterns: []glob.Glob{glob.MustCompile("hello")},
				exclude_patterns: []glob.Glob{glob.MustCompile("*")},
			},
			"hello",
			false,
		},
		{
			"test_5",
			fields{
				include_patterns: []glob.Glob{glob.MustCompile("openstack_nova*"), glob.MustCompile("openstack_cinder*")},
				exclude_patterns: []glob.Glob{glob.MustCompile("openstack_cinder_attachment*")},
			},
			"openstack_cinder_attachments",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pm := &PatternMatcher{
				include_patterns: tt.fields.include_patterns,
				exclude_patterns: tt.fields.exclude_patterns,
			}
			if got := pm.Match(tt.value); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
