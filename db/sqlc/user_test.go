package db

import (
	"context"
	"testing"
	"time"

	"github.com/maliByatzes/parte-server/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	arg := CreateUserParams{
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.ID)
	require.Equal(t, false, user.IsVerified)
	require.Equal(t, false, user.IsSuperuser)
	require.NotEmpty(t, user.Thumbnail)
	require.NotZero(t, user.UpdatedAt)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateuser(t *testing.T) {
	createRandomUser(t)
}

func TestGetuser(t *testing.T) {
	user1 := createRandomUser(t)
	require.NotEmpty(t, user1)

	user2, err := testStore.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.Equal(t, user1.IsSuperuser, user2.IsSuperuser)
	require.Equal(t, user1.Thumbnail, user2.Thumbnail)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	require.NotEmpty(t, user1)

	hashedPassword, err := util.HashPassword(util.RandomString(12))
	require.NoError(t, err)

	arg := UpdateUserParams{
		ID:             user1.ID,
		Username:       util.RandomUsername(),
		HashedPassword: hashedPassword,
		Email:          util.RandomEmail(),
		IsVerified:     user1.IsVerified,
		IsSuperuser:    user1.IsSuperuser,
		Thumbnail:      user1.Thumbnail,
	}

	user2, err := testStore.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, user1.IsVerified, user2.IsVerified)
	require.Equal(t, user1.IsSuperuser, user2.IsSuperuser)
	require.Equal(t, user1.Thumbnail, user2.Thumbnail)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
