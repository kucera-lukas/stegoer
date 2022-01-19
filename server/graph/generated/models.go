// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"stegoer/ent"
	"stegoer/ent/image"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type Auth struct {
	Ok      bool      `json:"ok"`
	Token   string    `json:"token"`
	Expires time.Time `json:"expires"`
}

type AuthUser struct {
	Auth *Auth     `json:"auth"`
	User *ent.User `json:"user"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type NewImage struct {
	Channel image.Channel  `json:"channel"`
	File    graphql.Upload `json:"file"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type UpdateUser struct {
	Name     *string `json:"name"`
	Password *string `json:"password"`
}