package gid

import (
	"encoding/json"
	"fmt"
)

// AbandonedCheckoutID represents a GID for a Shopify Abandoned Checkout.
type AbandonedCheckoutID int

func (acid AbandonedCheckoutID) ID() int {
	return int(acid)
}

func (acid AbandonedCheckoutID) String() string {
	return fmt.Sprintf("gid://shopify/AbandonedCheckout/%d", acid)
}

func (acid AbandonedCheckoutID) Equal(oacid Identifier) bool {
	if other, ok := oacid.(AbandonedCheckoutID); ok {
		return acid == other
	}
	return false
}

func (acid AbandonedCheckoutID) MarshalJSON() ([]byte, error) {
	return json.Marshal(acid.String())
}

func (acid *AbandonedCheckoutID) UnmarshalJSON(data []byte) error {
	var val any
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	*acid = New[AbandonedCheckoutID](val)
	return nil
}

func (acid AbandonedCheckoutID) IsValid() bool {
	return acid.ID() > 0
}

// AbandonedCheckoutIDs represents a slice of AbandonedCheckoutID.
type AbandonedCheckoutIDs []AbandonedCheckoutID

func (acids AbandonedCheckoutIDs) ToIDs() []int {
	return ToIDs(acids)
}

func (acids AbandonedCheckoutIDs) ToStrings() []string {
	return ToStrings(acids)
}

// NewAbandonedCheckoutID creates a new AbandonedCheckoutID from a value.
func NewAbandonedCheckoutID(val any) AbandonedCheckoutID {
	return New[AbandonedCheckoutID](val)
}

// NewAbandonedCheckoutIDValidated creates a new AbandonedCheckoutID from a value and validates it.
func NewAbandonedCheckoutIDValidated(val any) (AbandonedCheckoutID, error) {
	return NewValidated[AbandonedCheckoutID](val)
}
