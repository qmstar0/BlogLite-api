package auth

import (
	"context"
	"errors"
	"slices"
)

func FilterAuthWithUserType(ctx context.Context, usertype ...string) error {
	user, err := GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	if !slices.Contains(usertype, user.Type) {
		return errors.New("没有权限")
	}
	return nil
}
