package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func testJson(t *testing.T) {

}

func TestFixApp(t *testing.T) {
	err := FixApp()
	require.NoError(t, err)
}
