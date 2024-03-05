package pattern_matcher

import (
	"github.com/gobwas/glob"
	"github.com/rs/zerolog/log"
)

type PatternMatcher struct {
	include_patterns []glob.Glob
	exclude_patterns []glob.Glob
}

// Option is the type for functional options.
type Option func(*PatternMatcher)

func New(options ...Option) *PatternMatcher {
	pm := &PatternMatcher{
		include_patterns: []glob.Glob{glob.MustCompile("*")},
		exclude_patterns: []glob.Glob{},
	}

	for _, option := range options {
		option(pm)
	}
	return pm
}

func Replace(list []string, is_include bool) Option {
	return func(c *PatternMatcher) {
		if len(list) > 0 {
			patterns := []glob.Glob{}
			for _, pattern := range list {
				g, err := glob.Compile(pattern)
				if err != nil {
					log.Error().Err(err).Msgf("Error compiling pattern %s", pattern)
					continue
				}
				patterns = append(patterns, g)
			}
			if len(patterns) > 0 {
				c.include_patterns = patterns
				if !is_include {
					c.exclude_patterns = patterns
				}
			}			
		}
	}
}

func WithReplaceIncludes(list []string) Option {
	return Replace(list, true)
}

func WithReplaceExcludes(list []string) Option {
	return Replace(list, false)
}

func (pm *PatternMatcher) Match(value string) bool {

	is_included := false
	// Check if the value matches any pattern in the pm.include list
	for _, pattern := range pm.include_patterns {
		matched := pattern.Match(value)
		if matched {
			is_included = true
			break
		}
	}
	if !is_included {
		return false
	}
	// Check if the value matches any pattern in the excluded list
	for _, pattern := range pm.exclude_patterns {
		matched := pattern.Match(value)
		if matched {
			return false
		}
	}
	// If the value matches at least one pattern in the included list
	// and does not match any pattern in the excluded list, return true
	return true
}
