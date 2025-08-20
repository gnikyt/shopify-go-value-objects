package gid

import (
	"encoding/json"
	"fmt"
)

// VariantID represents a GID for a Shopify Product Variant.
type VariantID int

func (vid VariantID) ID() int {
	return int(vid)
}

func (vid VariantID) String() string {
	return fmt.Sprintf("gid://shopify/ProductVariant/%d", vid)
}

func (vid VariantID) Equal(ovid Identifier) bool {
	if other, ok := ovid.(VariantID); ok {
		return vid == other
	}
	return false
}

func (vid VariantID) MarshalJSON() ([]byte, error) {
	return json.Marshal(vid.String())
}

func (vid *VariantID) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*vid = New[VariantID](val)
	return nil
}

func (vid VariantID) IsValid() bool {
	return vid.ID() > 0
}

// VariantIDs represents a slice of VariantID.
type VariantIDs []VariantID

func (vids VariantIDs) ToIDs() []int {
	return ToIDs(vids)
}

func (vids VariantIDs) ToStrings() []string {
	return ToStrings(vids)
}

// NewVariantID creates a new VariantID from a value.
func NewVariantID(val any) VariantID {
	return New[VariantID](val)
}

// NewVariantIDValidated creates a new VariantID from a value and validates it.
func NewVariantIDValidated(val any) (VariantID, error) {
	return NewValidated[VariantID](val)
}
