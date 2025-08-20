package gid

import (
	"fmt"
	"strconv"
	"strings"
)

// Identifiers represents a set of GID types.
type Identifiers interface {
	CustomerID | ProductID | VariantID | OrderID | InventoryItemID | AbandonedCheckoutID
}

// Identifier represents a GID.
type Identifier interface {
	ID() int
	String() string
	Equal(ogid Identifier) bool
	IsValid() bool
}

// ToIDs converts GIDs to their int representation.
func ToIDs[T Identifiers](gids []T) []int {
	ids := make([]int, len(gids))
	for i, gid := range gids {
		ids[i] = Identifier(gid).ID()
	}
	return ids
}

// ToStrings converts GIDs to their string representation.
func ToStrings[T Identifiers](gids []T) []string {
	strs := make([]string, len(gids))
	for i, gid := range gids {
		strs[i] = Identifier(gid).String()
	}
	return strs
}

// typeFrom returns the type name for a given GID type.
func typeFrom[T Identifiers]() string {
	switch any(*new(T)).(type) {
	case CustomerID:
		return "Customer"
	case ProductID:
		return "Product"
	case VariantID:
		return "ProductVariant"
	case OrderID:
		return "Order"
	case InventoryItemID:
		return "InventoryItem"
	case AbandonedCheckoutID:
		return "AbandonedCheckout"
	default:
		return ""
	}
}

// commonNew creates a new GID from a value.
// It supports int64, int, float64, float32, and string formats.
// String formats supported: full GID ("gid://shopify/Type/ID") or numeric string ("123456789").
func commonNew[T Identifiers](val any) (T, error) {
	switch v := val.(type) {
	case int64:
		return T(int(v)), nil
	case int:
		return T(int(v)), nil
	case float64:
		return T(int(v)), nil
	case float32:
		return T(int(v)), nil
	case string:
		parts := strings.Split(v, "/")
		if len(parts) >= 5 {
			// Handle full GID format: "gid://shopify/Type/ID"
			typ := typeFrom[T]()
			if parts[3] != typ {
				return T(0), fmt.Errorf("expected type %s got %s", typ, parts[3])
			}
			cint, err := strconv.ParseInt(parts[4], 10, 64)
			if err != nil {
				return T(0), fmt.Errorf("invalid ID in GID: %s", v)
			}
			return T(int(cint)), nil
		} else {
			// Handle numeric string format: "123456789"
			cint, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return T(0), fmt.Errorf("invalid GID format or numeric ID: %v", v)
			}
			return T(int(cint)), nil
		}
	default:
		return T(0), fmt.Errorf("unsupported type for GID: %T", val)
	}
}

// New creates a new GID from a value.
// It supports int64, int, float64, float32, and string formats.
// For strings, it accepts either full GID format "gid://shopify/Type/ID" or numeric string "123456789".
// If the value is not recognized, it returns a zero value of the type.
// It ignores errors in validation, for validation use NewValidated.
func New[T Identifiers](val any) T {
	n, _ := commonNew[T](val)
	return n
}

// NewValidated creates a new GID from a value and validates it.
// It returns an error if the GID is not valid.
func NewValidated[T Identifiers](val any) (T, error) {
	gid, err := commonNew[T](val)
	if err != nil {
		return gid, err
	}
	if !Identifier(gid).IsValid() {
		return gid, fmt.Errorf("invalid %T value: %v", gid, val)
	}
	return gid, nil
}
