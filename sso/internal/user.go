package internal

import (
	"github.com/google/uuid"
	"time"
)

type Role string
type Email string
type PasswordBCrypto string

type UserRole struct {
	ID uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk" json:"id,omitempty"`

	Role Role `json:"role,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	ApplicationID uuid.UUID `json:"application_id,omitempty"`
	UserID        uuid.UUID `json:"user_id,omitempty"`
}

type UserRoles []*UserRole

type User struct {
	ID              uuid.UUID `bun:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	Email           Email     `json:"email"`
	PasswordBCrypto string    `json:"password_bcrypto"`

	FullName string `json:"full_name"`

	Roles UserRoles `json:"roles" bun:"rel:has-many,join:id=user_id"`

	CreatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"created_at" `
	UpdatedAt time.Time  `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at" `
	DeletedAt *time.Time `bun:",soft_delete,nullzero" json:"-,omitempty" `
}

type FindUserOptions struct {
	Email Email `json:"email,omitempty"`
	ID    uuid.UUID
}
