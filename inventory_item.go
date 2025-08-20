package gid

import (
	"encoding/json"
	"fmt"
)

// InventoryItemID represents a GID for a Shopify Inventory Item.
type InventoryItemID int

func (iiid InventoryItemID) ID() int {
	return int(iiid)
}

func (iiid InventoryItemID) String() string {
	return fmt.Sprintf("gid://shopify/InventoryItem/%d", iiid)
}

func (iiid InventoryItemID) Equal(oiiid Identifier) bool {
	if other, ok := oiiid.(InventoryItemID); ok {
		return iiid == other
	}
	return false
}

func (iiid InventoryItemID) MarshalJSON() ([]byte, error) {
	return json.Marshal(iiid.String())
}

func (iiid *InventoryItemID) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*iiid = New[InventoryItemID](val)
	return nil
}

func (iiid InventoryItemID) IsValid() bool {
	return iiid.ID() > 0
}

// InventoryItemIDs represents a slice of InventoryItem.
type InventoryItemIDs []InventoryItemID

func (iiids InventoryItemIDs) ToIDs() []int {
	return ToIDs(iiids)
}

func (iiids InventoryItemIDs) ToStrings() []string {
	return ToStrings(iiids)
}

// NewInventoryItemID creates a new InventoryItemID from a value.
func NewInventoryItemID(val any) InventoryItemID {
	return New[InventoryItemID](val)
}

// NewInventoryItemIDValidated creates a new InventoryItemID from a value and validates it.
func NewInventoryItemIDValidated(val any) (InventoryItemID, error) {
	return NewValidated[InventoryItemID](val)
}
