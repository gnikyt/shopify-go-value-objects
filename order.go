package gid

import (
	"encoding/json"
	"fmt"
)

// OrderID represents a GID for a Shopify Order.
type OrderID int

func (oid OrderID) ID() int {
	return int(oid)
}

func (oid OrderID) String() string {
	return fmt.Sprintf("gid://shopify/Order/%d", oid)
}

func (oid OrderID) Equal(ooid Identifier) bool {
	if other, ok := ooid.(OrderID); ok {
		return oid == other
	}
	return false
}

func (oid OrderID) MarshalJSON() ([]byte, error) {
	return json.Marshal(oid.String())
}

func (oid *OrderID) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*oid = New[OrderID](val)
	return nil
}

func (oid OrderID) IsValid() bool {
	return oid.ID() > 0
}

// OrderIDs represents a slice of OrderID.
type OrderIDs []OrderID

func (oids OrderIDs) ToIDs() []int {
	return ToIDs(oids)
}

func (oids OrderIDs) ToStrings() []string {
	return ToStrings(oids)
}

// NewOrderID creates a new OrderID from a value.
func NewOrderID(val any) OrderID {
	return New[OrderID](val)
}

// NewOrderIDValidated creates a new OrderID from a value and validates it.
func NewOrderIDValidated(val any) (OrderID, error) {
	return NewValidated[OrderID](val)
}
