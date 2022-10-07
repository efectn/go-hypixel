package gohypixel

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Client_Key(t *testing.T) {
	client := New(Config{APIKey: "0e9128f0-4c84-41de-838e-77fb63e71ae8"})
	fmt.Print(client.Key())
	require.Equal(t, "w", "d")
}
