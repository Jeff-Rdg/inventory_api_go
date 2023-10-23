package entity

import (
	"github.com/google/uuid"
	"time"
)

type base struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
}
