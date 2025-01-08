package internal

import (
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ID               uuid.UUID  `pg:"type:uuid,default:uuid_generate_v4(),pk" json:"id"`
	User             string     `json:"user"`
	Request          string     `json:"request"`
	Response         string     `json:"response"`
	ModifiedResponse string     `json:"modified_response"`
	CreatedAt        time.Time  `pg:"default:current_timestamp" json:"created_at"`
	UpdatedAt        time.Time  `pg:"default:current_timestamp" json:"updated_at"`
	DeletedAt        *time.Time `pg:",soft_delete" json:"-,omitempty"`
}
