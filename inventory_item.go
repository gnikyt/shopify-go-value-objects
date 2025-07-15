package gid

import (
	"encoding/json"
	"fmt"
)

// InventoryItem represents a GID for a Shopify Inventory Item.
type InventoryItem int

func (iiid InventoryItem) ID() int {
	return int(iiid)
}

func (iiid InventoryItem) String() string {
	return fmt.Sprintf("gid://shopify/InventoryItem/%d", iiid)
}

func (iiid InventoryItem) Equal(oiiid Identifier) bool {
	if other, ok := oiiid.(InventoryItem); ok {
		return iiid == other
	}
	return false
}

func (iiid InventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(iiid.String())
}

func (iiid *InventoryItem) UnmarshalJSON(data []byte) error {
	var gid string
	if err := json.Unmarshal(data, &gid); err != nil {
		return err
	}
	*iiid = New[InventoryItem](gid)
	return nil
}

func (iiid InventoryItem) IsValid() bool {
	return iiid.ID() > 0
}

// InventoryItems represents a slice of InventoryItem.
type InventoryItems []InventoryItem

func (iiids InventoryItems) ToIDs() []int {
	return ToIDs(iiids)
}

func (iiids InventoryItems) ToStrings() []string {
	return ToStrings(iiids)
}

// NewInventoryItem creates a new InventoryItem from a value.
func NewInventoryItem(val any) InventoryItem {
	return New[InventoryItem](val)
}

// NewInventoryItemValidated creates a new InventoryItem from a value and validates it.
func NewInventoryItemValidated(val any) (InventoryItem, error) {
	return NewValidated[InventoryItem](val)
}
