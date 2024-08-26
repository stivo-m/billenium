package domain

import (
	"database/sql/driver"
	"errors"
)

type Enumerations interface {
	Scan(value interface{}) error
	Value() (driver.Value, error)
}

type Enum[T ~string] struct {
	value T
}

func (e *Enum[T]) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		e.value = T(v)
	case []byte:
		e.value = T(v)
	default:
		return errors.New("invalid type for enum value")
	}
	return nil
}

func (e Enum[T]) Value() (driver.Value, error) {
	return string(e.value), nil
}

func (e Enum[T]) String() string {
	return string(e.value)
}

type TxnStatusOption string
type TxnTypeOption string

// Enum defs
type TxnStatus Enum[TxnStatusOption]
type TxnType Enum[TxnTypeOption]

// Enum values
const (
	INCOMING  TxnTypeOption = "incoming"
	OUTGOINGT TxnTypeOption = "outgoing"

	DRAFT     TxnStatusOption = "draft"
	PENDING   TxnStatusOption = "pending"
	COMPLETED TxnStatusOption = "completed"
	CANCELLED TxnStatusOption = "cancelled"
)
