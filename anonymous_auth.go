package anonymous_auth

import (
	"context"
	"github.com/applicaset/user-svc"
)

const Name = "anonymous"

type auth struct {
}

type response struct {
	id string
}

func (rsp response) Validated() bool {
	return rsp.id != ""
}

func (rsp response) ID() string {
	return rsp.id
}

func (aa *auth) Validate(_ context.Context, args map[string]interface{}) (user.ValidateResponse, error) {
	rsp := new(response)
	iGuestID, ok := args["guest_id"]
	if !ok {
		return rsp, nil
	}

	guestID, ok := iGuestID.(string)
	if !ok {
		return rsp, nil
	}

	rsp.id = guestID

	return rsp, nil
}

func NewAuthProvider() user.AuthProvider {
	aa := auth{}

	return &aa
}

func New() user.Option {
	return user.WithAuthProvider(Name, NewAuthProvider())
}
