package gid

import (
	"encoding/json"
	"fmt"
)

// ProductID represents a GID for a Shopify Product.
type ProductID int

func (pid ProductID) ID() int {
	return int(pid)
}

func (pid ProductID) String() string {
	return fmt.Sprintf("gid://shopify/Product/%d", pid)
}

func (pid ProductID) Equal(opid Identifier) bool {
	if other, ok := opid.(ProductID); ok {
		return pid == other
	}
	return false
}

func (pid ProductID) MarshalJSON() ([]byte, error) {
	return json.Marshal(pid.String())
}

func (pid *ProductID) UnmarshalJSON(data []byte) error {
	var gid string
	if err := json.Unmarshal(data, &gid); err != nil {
		return err
	}
	*pid = New[ProductID](gid)
	return nil
}

func (pid ProductID) IsValid() bool {
	return pid.ID() > 0
}

// ProductIDs represents a slice of ProductID.
type ProductIDs []ProductID

func (pids ProductIDs) ToIDs() []int {
	return ToIDs(pids)
}

func (pids ProductIDs) ToStrings() []string {
	return ToStrings(pids)
}

// NewProductID creates a new ProductID from a value.
func NewProductID(val any) ProductID {
	return New[ProductID](val)
}

// NewProductIDValidated creates a new ProductID from a value and validates it.
func NewProductIDValidated(val any) (ProductID, error) {
	return NewValidated[ProductID](val)
}
