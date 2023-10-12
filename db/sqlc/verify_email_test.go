package db

import (
	"context"
	"testing"

	"github.com/maliByatzes/parte-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomVerifyEmail(t *testing.T) VerifyEmail {
	user := createRandomUser(t)
	require.NotEmpty(t, user)

	arg := CreateVerifyEmailParams{
		UserID:     user.ID,
		Email:      util.RandomEmail(),
		SecretCode: "858477",
	}

	vE, err := testStore.CreateVerifyEmail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, vE)

	require.Equal(t, arg.UserID, vE.UserID)
	require.Equal(t, arg.Email, vE.Email)
	require.Equal(t, arg.SecretCode, vE.SecretCode)
	require.NotZero(t, vE.ID)
	require.Equal(t, false, vE.IsUsed)
	require.NotZero(t, vE.ExpiredAt)
	require.NotZero(t, vE.CreatedAt)

	return vE
}

func TestCreateVerifyEmail(t *testing.T) {
	createRandomVerifyEmail(t)
}
