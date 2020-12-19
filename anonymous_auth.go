package anonymousauth

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

func (aa *auth) Validate(_ context.Context, args map[string]interface{}) (usersvc.ValidateResponse, error) {
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

func NewAuthProvider() usersvc.AuthProvider {
	aa := auth{}

	return &aa
}

func New() usersvc.Option {
	return usersvc.WithAuthProvider(Name, NewAuthProvider())
}
