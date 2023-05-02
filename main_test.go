package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// this is just a test I was running locally to ensure that I could kill a process.
// this should not be run in any test suite since this is a hardcoded pid
func Test_KillPort(t *testing.T) {
	err := logKillPort("81146")

	assert.Nil(t, err)
}
