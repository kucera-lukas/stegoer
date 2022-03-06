// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"time"

	"github.com/kucera-lukas/stegoer/ent/image"
	"github.com/kucera-lukas/stegoer/ent/predicate"
	"github.com/kucera-lukas/stegoer/ent/schema/ulid"
	"github.com/kucera-lukas/stegoer/ent/user"
)

// ImageWhereInput represents a where input for filtering Image queries.
type ImageWhereInput struct {
	Not *ImageWhereInput   `json:"not,omitempty"`
	Or  []*ImageWhereInput `json:"or,omitempty"`
	And []*ImageWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *ulid.ID  `json:"id,omitempty"`
	IDNEQ   *ulid.ID  `json:"idNEQ,omitempty"`
	IDIn    []ulid.ID `json:"idIn,omitempty"`
	IDNotIn []ulid.ID `json:"idNotIn,omitempty"`
	IDGT    *ulid.ID  `json:"idGT,omitempty"`
	IDGTE   *ulid.ID  `json:"idGTE,omitempty"`
	IDLT    *ulid.ID  `json:"idLT,omitempty"`
	IDLTE   *ulid.ID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "channel" field predicates.
	Channel      *image.Channel  `json:"channel,omitempty"`
	ChannelNEQ   *image.Channel  `json:"channelNEQ,omitempty"`
	ChannelIn    []image.Channel `json:"channelIn,omitempty"`
	ChannelNotIn []image.Channel `json:"channelNotIn,omitempty"`

	// "user" edge predicates.
	HasUser     *bool             `json:"hasUser,omitempty"`
	HasUserWith []*UserWhereInput `json:"hasUserWith,omitempty"`
}

// Filter applies the ImageWhereInput filter on the ImageQuery builder.
func (i *ImageWhereInput) Filter(q *ImageQuery) (*ImageQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering images.
// An error is returned if the input is empty or invalid.
func (i *ImageWhereInput) P() (predicate.Image, error) {
	var predicates []predicate.Image
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, image.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Image, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, image.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Image, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, image.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, image.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, image.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, image.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, image.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, image.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, image.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, image.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, image.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, image.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, image.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, image.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, image.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, image.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, image.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, image.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, image.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, image.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, image.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, image.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, image.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, image.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, image.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, image.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, image.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Channel != nil {
		predicates = append(predicates, image.ChannelEQ(*i.Channel))
	}
	if i.ChannelNEQ != nil {
		predicates = append(predicates, image.ChannelNEQ(*i.ChannelNEQ))
	}
	if len(i.ChannelIn) > 0 {
		predicates = append(predicates, image.ChannelIn(i.ChannelIn...))
	}
	if len(i.ChannelNotIn) > 0 {
		predicates = append(predicates, image.ChannelNotIn(i.ChannelNotIn...))
	}

	if i.HasUser != nil {
		p := image.HasUser()
		if !*i.HasUser {
			p = image.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasUserWith) > 0 {
		with := make([]predicate.User, 0, len(i.HasUserWith))
		for _, w := range i.HasUserWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, image.HasUserWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/kucera-lukas/stegoer/ent: empty predicate ImageWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return image.And(predicates...), nil
	}
}

// UserWhereInput represents a where input for filtering User queries.
type UserWhereInput struct {
	Not *UserWhereInput   `json:"not,omitempty"`
	Or  []*UserWhereInput `json:"or,omitempty"`
	And []*UserWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *ulid.ID  `json:"id,omitempty"`
	IDNEQ   *ulid.ID  `json:"idNEQ,omitempty"`
	IDIn    []ulid.ID `json:"idIn,omitempty"`
	IDNotIn []ulid.ID `json:"idNotIn,omitempty"`
	IDGT    *ulid.ID  `json:"idGT,omitempty"`
	IDGTE   *ulid.ID  `json:"idGTE,omitempty"`
	IDLT    *ulid.ID  `json:"idLT,omitempty"`
	IDLTE   *ulid.ID  `json:"idLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "updated_at" field predicates.
	UpdatedAt      *time.Time  `json:"updatedAt,omitempty"`
	UpdatedAtNEQ   *time.Time  `json:"updatedAtNEQ,omitempty"`
	UpdatedAtIn    []time.Time `json:"updatedAtIn,omitempty"`
	UpdatedAtNotIn []time.Time `json:"updatedAtNotIn,omitempty"`
	UpdatedAtGT    *time.Time  `json:"updatedAtGT,omitempty"`
	UpdatedAtGTE   *time.Time  `json:"updatedAtGTE,omitempty"`
	UpdatedAtLT    *time.Time  `json:"updatedAtLT,omitempty"`
	UpdatedAtLTE   *time.Time  `json:"updatedAtLTE,omitempty"`

	// "name" field predicates.
	Name             *string  `json:"name,omitempty"`
	NameNEQ          *string  `json:"nameNEQ,omitempty"`
	NameIn           []string `json:"nameIn,omitempty"`
	NameNotIn        []string `json:"nameNotIn,omitempty"`
	NameGT           *string  `json:"nameGT,omitempty"`
	NameGTE          *string  `json:"nameGTE,omitempty"`
	NameLT           *string  `json:"nameLT,omitempty"`
	NameLTE          *string  `json:"nameLTE,omitempty"`
	NameContains     *string  `json:"nameContains,omitempty"`
	NameHasPrefix    *string  `json:"nameHasPrefix,omitempty"`
	NameHasSuffix    *string  `json:"nameHasSuffix,omitempty"`
	NameEqualFold    *string  `json:"nameEqualFold,omitempty"`
	NameContainsFold *string  `json:"nameContainsFold,omitempty"`

	// "email" field predicates.
	Email             *string  `json:"email,omitempty"`
	EmailNEQ          *string  `json:"emailNEQ,omitempty"`
	EmailIn           []string `json:"emailIn,omitempty"`
	EmailNotIn        []string `json:"emailNotIn,omitempty"`
	EmailGT           *string  `json:"emailGT,omitempty"`
	EmailGTE          *string  `json:"emailGTE,omitempty"`
	EmailLT           *string  `json:"emailLT,omitempty"`
	EmailLTE          *string  `json:"emailLTE,omitempty"`
	EmailContains     *string  `json:"emailContains,omitempty"`
	EmailHasPrefix    *string  `json:"emailHasPrefix,omitempty"`
	EmailHasSuffix    *string  `json:"emailHasSuffix,omitempty"`
	EmailEqualFold    *string  `json:"emailEqualFold,omitempty"`
	EmailContainsFold *string  `json:"emailContainsFold,omitempty"`

	// "password" field predicates.
	Password             *string  `json:"password,omitempty"`
	PasswordNEQ          *string  `json:"passwordNEQ,omitempty"`
	PasswordIn           []string `json:"passwordIn,omitempty"`
	PasswordNotIn        []string `json:"passwordNotIn,omitempty"`
	PasswordGT           *string  `json:"passwordGT,omitempty"`
	PasswordGTE          *string  `json:"passwordGTE,omitempty"`
	PasswordLT           *string  `json:"passwordLT,omitempty"`
	PasswordLTE          *string  `json:"passwordLTE,omitempty"`
	PasswordContains     *string  `json:"passwordContains,omitempty"`
	PasswordHasPrefix    *string  `json:"passwordHasPrefix,omitempty"`
	PasswordHasSuffix    *string  `json:"passwordHasSuffix,omitempty"`
	PasswordEqualFold    *string  `json:"passwordEqualFold,omitempty"`
	PasswordContainsFold *string  `json:"passwordContainsFold,omitempty"`

	// "last_login" field predicates.
	LastLogin      *time.Time  `json:"lastLogin,omitempty"`
	LastLoginNEQ   *time.Time  `json:"lastLoginNEQ,omitempty"`
	LastLoginIn    []time.Time `json:"lastLoginIn,omitempty"`
	LastLoginNotIn []time.Time `json:"lastLoginNotIn,omitempty"`
	LastLoginGT    *time.Time  `json:"lastLoginGT,omitempty"`
	LastLoginGTE   *time.Time  `json:"lastLoginGTE,omitempty"`
	LastLoginLT    *time.Time  `json:"lastLoginLT,omitempty"`
	LastLoginLTE   *time.Time  `json:"lastLoginLTE,omitempty"`

	// "images" edge predicates.
	HasImages     *bool              `json:"hasImages,omitempty"`
	HasImagesWith []*ImageWhereInput `json:"hasImagesWith,omitempty"`
}

// Filter applies the UserWhereInput filter on the UserQuery builder.
func (i *UserWhereInput) Filter(q *UserQuery) (*UserQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		return nil, err
	}
	return q.Where(p), nil
}

// P returns a predicate for filtering users.
// An error is returned if the input is empty or invalid.
func (i *UserWhereInput) P() (predicate.User, error) {
	var predicates []predicate.User
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, user.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.User, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			or = append(or, p)
		}
		predicates = append(predicates, user.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, err
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.User, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			and = append(and, p)
		}
		predicates = append(predicates, user.And(and...))
	}
	if i.ID != nil {
		predicates = append(predicates, user.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, user.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, user.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, user.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, user.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, user.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, user.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, user.IDLTE(*i.IDLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, user.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, user.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, user.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, user.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, user.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, user.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, user.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, user.CreatedAtLTE(*i.CreatedAtLTE))
	}
	if i.UpdatedAt != nil {
		predicates = append(predicates, user.UpdatedAtEQ(*i.UpdatedAt))
	}
	if i.UpdatedAtNEQ != nil {
		predicates = append(predicates, user.UpdatedAtNEQ(*i.UpdatedAtNEQ))
	}
	if len(i.UpdatedAtIn) > 0 {
		predicates = append(predicates, user.UpdatedAtIn(i.UpdatedAtIn...))
	}
	if len(i.UpdatedAtNotIn) > 0 {
		predicates = append(predicates, user.UpdatedAtNotIn(i.UpdatedAtNotIn...))
	}
	if i.UpdatedAtGT != nil {
		predicates = append(predicates, user.UpdatedAtGT(*i.UpdatedAtGT))
	}
	if i.UpdatedAtGTE != nil {
		predicates = append(predicates, user.UpdatedAtGTE(*i.UpdatedAtGTE))
	}
	if i.UpdatedAtLT != nil {
		predicates = append(predicates, user.UpdatedAtLT(*i.UpdatedAtLT))
	}
	if i.UpdatedAtLTE != nil {
		predicates = append(predicates, user.UpdatedAtLTE(*i.UpdatedAtLTE))
	}
	if i.Name != nil {
		predicates = append(predicates, user.NameEQ(*i.Name))
	}
	if i.NameNEQ != nil {
		predicates = append(predicates, user.NameNEQ(*i.NameNEQ))
	}
	if len(i.NameIn) > 0 {
		predicates = append(predicates, user.NameIn(i.NameIn...))
	}
	if len(i.NameNotIn) > 0 {
		predicates = append(predicates, user.NameNotIn(i.NameNotIn...))
	}
	if i.NameGT != nil {
		predicates = append(predicates, user.NameGT(*i.NameGT))
	}
	if i.NameGTE != nil {
		predicates = append(predicates, user.NameGTE(*i.NameGTE))
	}
	if i.NameLT != nil {
		predicates = append(predicates, user.NameLT(*i.NameLT))
	}
	if i.NameLTE != nil {
		predicates = append(predicates, user.NameLTE(*i.NameLTE))
	}
	if i.NameContains != nil {
		predicates = append(predicates, user.NameContains(*i.NameContains))
	}
	if i.NameHasPrefix != nil {
		predicates = append(predicates, user.NameHasPrefix(*i.NameHasPrefix))
	}
	if i.NameHasSuffix != nil {
		predicates = append(predicates, user.NameHasSuffix(*i.NameHasSuffix))
	}
	if i.NameEqualFold != nil {
		predicates = append(predicates, user.NameEqualFold(*i.NameEqualFold))
	}
	if i.NameContainsFold != nil {
		predicates = append(predicates, user.NameContainsFold(*i.NameContainsFold))
	}
	if i.Email != nil {
		predicates = append(predicates, user.EmailEQ(*i.Email))
	}
	if i.EmailNEQ != nil {
		predicates = append(predicates, user.EmailNEQ(*i.EmailNEQ))
	}
	if len(i.EmailIn) > 0 {
		predicates = append(predicates, user.EmailIn(i.EmailIn...))
	}
	if len(i.EmailNotIn) > 0 {
		predicates = append(predicates, user.EmailNotIn(i.EmailNotIn...))
	}
	if i.EmailGT != nil {
		predicates = append(predicates, user.EmailGT(*i.EmailGT))
	}
	if i.EmailGTE != nil {
		predicates = append(predicates, user.EmailGTE(*i.EmailGTE))
	}
	if i.EmailLT != nil {
		predicates = append(predicates, user.EmailLT(*i.EmailLT))
	}
	if i.EmailLTE != nil {
		predicates = append(predicates, user.EmailLTE(*i.EmailLTE))
	}
	if i.EmailContains != nil {
		predicates = append(predicates, user.EmailContains(*i.EmailContains))
	}
	if i.EmailHasPrefix != nil {
		predicates = append(predicates, user.EmailHasPrefix(*i.EmailHasPrefix))
	}
	if i.EmailHasSuffix != nil {
		predicates = append(predicates, user.EmailHasSuffix(*i.EmailHasSuffix))
	}
	if i.EmailEqualFold != nil {
		predicates = append(predicates, user.EmailEqualFold(*i.EmailEqualFold))
	}
	if i.EmailContainsFold != nil {
		predicates = append(predicates, user.EmailContainsFold(*i.EmailContainsFold))
	}
	if i.Password != nil {
		predicates = append(predicates, user.PasswordEQ(*i.Password))
	}
	if i.PasswordNEQ != nil {
		predicates = append(predicates, user.PasswordNEQ(*i.PasswordNEQ))
	}
	if len(i.PasswordIn) > 0 {
		predicates = append(predicates, user.PasswordIn(i.PasswordIn...))
	}
	if len(i.PasswordNotIn) > 0 {
		predicates = append(predicates, user.PasswordNotIn(i.PasswordNotIn...))
	}
	if i.PasswordGT != nil {
		predicates = append(predicates, user.PasswordGT(*i.PasswordGT))
	}
	if i.PasswordGTE != nil {
		predicates = append(predicates, user.PasswordGTE(*i.PasswordGTE))
	}
	if i.PasswordLT != nil {
		predicates = append(predicates, user.PasswordLT(*i.PasswordLT))
	}
	if i.PasswordLTE != nil {
		predicates = append(predicates, user.PasswordLTE(*i.PasswordLTE))
	}
	if i.PasswordContains != nil {
		predicates = append(predicates, user.PasswordContains(*i.PasswordContains))
	}
	if i.PasswordHasPrefix != nil {
		predicates = append(predicates, user.PasswordHasPrefix(*i.PasswordHasPrefix))
	}
	if i.PasswordHasSuffix != nil {
		predicates = append(predicates, user.PasswordHasSuffix(*i.PasswordHasSuffix))
	}
	if i.PasswordEqualFold != nil {
		predicates = append(predicates, user.PasswordEqualFold(*i.PasswordEqualFold))
	}
	if i.PasswordContainsFold != nil {
		predicates = append(predicates, user.PasswordContainsFold(*i.PasswordContainsFold))
	}
	if i.LastLogin != nil {
		predicates = append(predicates, user.LastLoginEQ(*i.LastLogin))
	}
	if i.LastLoginNEQ != nil {
		predicates = append(predicates, user.LastLoginNEQ(*i.LastLoginNEQ))
	}
	if len(i.LastLoginIn) > 0 {
		predicates = append(predicates, user.LastLoginIn(i.LastLoginIn...))
	}
	if len(i.LastLoginNotIn) > 0 {
		predicates = append(predicates, user.LastLoginNotIn(i.LastLoginNotIn...))
	}
	if i.LastLoginGT != nil {
		predicates = append(predicates, user.LastLoginGT(*i.LastLoginGT))
	}
	if i.LastLoginGTE != nil {
		predicates = append(predicates, user.LastLoginGTE(*i.LastLoginGTE))
	}
	if i.LastLoginLT != nil {
		predicates = append(predicates, user.LastLoginLT(*i.LastLoginLT))
	}
	if i.LastLoginLTE != nil {
		predicates = append(predicates, user.LastLoginLTE(*i.LastLoginLTE))
	}

	if i.HasImages != nil {
		p := user.HasImages()
		if !*i.HasImages {
			p = user.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasImagesWith) > 0 {
		with := make([]predicate.Image, 0, len(i.HasImagesWith))
		for _, w := range i.HasImagesWith {
			p, err := w.P()
			if err != nil {
				return nil, err
			}
			with = append(with, p)
		}
		predicates = append(predicates, user.HasImagesWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, fmt.Errorf("github.com/kucera-lukas/stegoer/ent: empty predicate UserWhereInput")
	case 1:
		return predicates[0], nil
	default:
		return user.And(predicates...), nil
	}
}
