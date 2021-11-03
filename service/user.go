package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/haleyrc/api"
	"github.com/haleyrc/api/errors"
)

type UserService struct {
	DB Database
}

type GetUserRequest struct {
	ID   api.ID
	Name string
}

type GetUserResponse struct {
	User api.User
}

func (s UserService) GetUser(ctx context.Context, req GetUserRequest) (*GetUserResponse, error) {
	req.Name = strings.TrimSpace(req.Name)

	if req.ID == "" && req.Name == "" {
		return nil, fmt.Errorf("get user failed: %w",
			errors.BadRequest{Message: "Either ID or name is required to get a user."})
	}

	var resp GetUserResponse
	err := s.DB.RunInTransaction(ctx, func(ctx context.Context, tx Tx) error {
		var err error
		if req.ID != "" {
			resp.User, err = tx.GetUserByID(ctx, req.ID)
		} else {
			resp.User, err = tx.GetUserByName(ctx, req.Name)
		}
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("get user failed: %w", err)
	}

	return &resp, nil
}
