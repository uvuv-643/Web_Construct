package internal

import (
	"github.com/google/uuid"
	"time"
)

type Role string
type Email string
type PasswordBCrypto string

type UserRole struct {
	ID            uuid.UUID `pg:"type:uuid,default:uuid_generate_v4(),pk" json:"id,omitempty"`
	Role          Role      `json:"role,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
	ApplicationID string    `json:"application_id,omitempty"`
	UserID        uuid.UUID `pg:"type:uuid" json:"user_id,omitempty"`
}

type UserRoles []*UserRole

type User struct {
	ID              uuid.UUID  `pg:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	Email           Email      `json:"email"`
	PasswordBCrypto string     `json:"-"`
	Roles           UserRoles  `pg:"rel:has-many,join:id=user_id" json:"roles"`
	CreatedAt       time.Time  `pg:"default:current_timestamp" json:"created_at"`
	UpdatedAt       time.Time  `pg:"default:current_timestamp" json:"updated_at"`
	DeletedAt       *time.Time `pg:",soft_delete" json:"-,omitempty"`
}

type FindUserOptions struct {
	Email Email     `json:"email,omitempty"`
	ID    uuid.UUID `json:"id"`
}
