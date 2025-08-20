package gid

import (
	"encoding/json"
	"fmt"
)

// CustomerID represents a GID for a Shopify Customer.
type CustomerID int

func (cid CustomerID) ID() int {
	return int(cid)
}

func (cid CustomerID) String() string {
	return fmt.Sprintf("gid://shopify/Customer/%d", cid)
}

func (cid CustomerID) Equal(ocid Identifier) bool {
	if other, ok := ocid.(CustomerID); ok {
		return cid == other
	}
	return false
}

func (cid CustomerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(cid.String())
}

func (cid *CustomerID) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*cid = New[CustomerID](val)
	return nil
}

func (cid CustomerID) IsValid() bool {
	return cid.ID() > 0
}

// CustomerIDs represents a slice of CustomerID.
type CustomerIDs []CustomerID

func (cids CustomerIDs) ToIDs() []int {
	return ToIDs(cids)
}

func (cids CustomerIDs) ToStrings() []string {
	return ToStrings(cids)
}

// NewCustomerID creates a new CustomerID from a value.
func NewCustomerID(val any) CustomerID {
	return New[CustomerID](val)
}

// NewCustomerIDValidated creates a new CustomerID from a value and validates it.
func NewCustomerIDValidated(val any) (CustomerID, error) {
	return NewValidated[CustomerID](val)
}
