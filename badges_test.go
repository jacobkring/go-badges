package main

import (
	"github.com/jacobkring/go-assert"
	"testing"
)

func TestCoverageBadge(t *testing.T) {
	coverageBadge, err := CoverageBadge("googly")
	assert.NotNil(t, err)
	assert.Equal(t, "strconv.ParseFloat: parsing \"googly\": invalid syntax", err.Error())
	coverageBadge, err = CoverageBadge("90")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/90%25/green)", coverageBadge)
	coverageBadge, err = CoverageBadge("75")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/75%25/yellow)", coverageBadge)
	coverageBadge, err = CoverageBadge("65")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/65%25/orange)", coverageBadge)
	coverageBadge, err = CoverageBadge("50")
	assert.Nil(t, err)
	assert.Equal(t, "![](https://badgen.net/badge/coverage/50%25/red)", coverageBadge)
}
