package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfig(t *testing.T) {
	_, err := LoadConfig()
	require.NoError(t, err)
}
