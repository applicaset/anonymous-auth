package anonymous_auth_test

import (
	"context"
	"github.com/applicaset/anonymous-auth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAuth_Validate(t *testing.T) {
	auth := anonymous_auth.NewAuthProvider()

	ctx := context.Background()

	t.Run("Test Valid Arguments", func(t *testing.T) {
		guestID := "1"
		res, err := auth.Validate(ctx, map[string]interface{}{"guest_id": guestID})
		require.NoError(t, err)

		assert.True(t, res.Validated())
		assert.Equal(t, guestID, res.ID())
	})

	t.Run("Test Invalid Argument Type", func(t *testing.T) {
		res, err := auth.Validate(ctx, map[string]interface{}{"guest_id": 1})
		require.NoError(t, err)

		assert.False(t, res.Validated())
		assert.Zero(t, res.ID())
	})

	t.Run("Test Invalid Arguments", func(t *testing.T) {
		res, err := auth.Validate(ctx, nil)
		require.NoError(t, err)

		assert.False(t, res.Validated())
		assert.Zero(t, res.ID())
	})

	t.Run("Test Nil Arguments", func(t *testing.T) {
		res, err := auth.Validate(ctx, nil)
		require.NoError(t, err)

		assert.False(t, res.Validated())
		assert.Zero(t, res.ID())
	})
}
