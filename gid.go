package gid

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// kind defined a Shopify GID resource name.
// The name is the resource segment of a GID: "gid://shopify/<Name>/<ID>".
type kind interface {
	name() string
}

// GID is a Shopify GID resource of kind K.
type GID[K kind] int64

// ID returns the numeric portion of the GID.
func (g GID[K]) ID() int64 {
	return int64(g)
}

// String returns the full GID, such as "gid://shopify/Customer/123".
func (g GID[K]) String() string {
	var k K
	return fmt.Sprintf("gid://shopify/%s/%d", k.name(), int64(g))
}

// IsValid reports whether the GID refers to a real resource.
func (g GID[K]) IsValid() bool {
	return g > 0
}

// Equal reports whether two GIDs of the same kind are equal.
func (g GID[K]) Equal(other GID[K]) bool {
	return g == other
}

// MarshalJSON implements JSON marshalling.
func (g GID[K]) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}

// UnmarshalJSON implements unmarshalling JSON.
func (g *GID[K]) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	parsed, err := parse[K](s)
	if err != nil {
		return err
	}
	*g = parsed
	return nil
}

// New creates a GID from a value, ignoring any parse error (the zero value is
// returned on failure).
//
// Supported inputs: int, int64, a full GID string ("gid://shopify/<Name>/<ID>"),
// or a bare numeric string ("123456789").
func New[K kind](val any) GID[K] {
	g, _ := parse[K](val)
	return g
}

// NewValidated creates a GID from a value and returns an error if the value is
// malformed, of the wrong kind, or does not refer to a real resource.
func NewValidated[K kind](val any) (GID[K], error) {
	g, err := parse[K](val)
	if err != nil {
		return g, err
	}
	if !g.IsValid() {
		return g, fmt.Errorf("gid: invalid value: %v", val)
	}
	return g, nil
}

// ToIDs converts a slice of GIDs to their numeric IDs.
func ToIDs[K kind](gids []GID[K]) []int64 {
	ids := make([]int64, len(gids))
	for i, g := range gids {
		ids[i] = g.ID()
	}
	return ids
}

// ToStrings converts a slice of GIDs to their full GID strings.
func ToStrings[K kind](gids []GID[K]) []string {
	strs := make([]string, len(gids))
	for i, g := range gids {
		strs[i] = g.String()
	}
	return strs
}

// parse builds a GID[K] from an int, int64, or string. Strings may be a full
// GID ("gid://shopify/<Name>/<ID>") or a bare numeric string ("123456789").
func parse[K kind](val any) (GID[K], error) {
	switch v := val.(type) {
	case int:
		return GID[K](v), nil
	case int64:
		return GID[K](v), nil
	case string:
		return parseString[K](v)
	default:
		return 0, fmt.Errorf("gid: unsupported type %T", val)
	}
}

func parseString[K kind](s string) (GID[K], error) {
	// Bare numeric string: "123456789".
	if id, err := strconv.ParseInt(s, 10, 64); err == nil {
		return GID[K](id), nil
	}

	// Full GID: "gid://shopify/<name>/<id>" splits into exactly five parts.
	parts := strings.Split(s, "/")
	if len(parts) != 5 || parts[0] != "gid:" || parts[2] != "shopify" {
		return 0, fmt.Errorf("gid: malformed GID %q", s)
	}

	var k K
	if want := k.name(); parts[3] != want {
		return 0, fmt.Errorf("gid: expected kind %s, got %s", want, parts[3])
	}

	id, err := strconv.ParseInt(parts[4], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("gid: invalid id in %q", s)
	}
	return GID[K](id), nil
}

//
// Examples.
//

// customerKind marks a GID as a Shopify Customer.
type customerKind struct{}

func (customerKind) name() string {
	return "Customer"
}

type (
	// CustomerID is the GID of a Shopify Customer.
	CustomerID = GID[customerKind]

	// CustomerIDs is a slice of CustomerID.
	CustomerIDs []CustomerID
)

func (cids CustomerIDs) ToIDs() []int64 {
	return ToIDs(cids)
}
func (cids CustomerIDs) ToStrings() []string {
	return ToStrings(cids)
}

// NewCustomerID creates a new CustomerID from a value.
func NewCustomerID(val any) CustomerID {
	return New[customerKind](val)
}

// NewCustomerIDValidated creates a new CustomerID from a value and validates it.
func NewCustomerIDValidated(val any) (CustomerID, error) {
	return NewValidated[customerKind](val)
}

// productKind marks a GID as a Shopify Product.
type productKind struct{}

func (productKind) name() string {
	return "Product"
}

type (
	// ProductID is the GID of a Shopify Product.
	ProductID = GID[productKind]

	// ProductIDs is a slice of ProductID.
	ProductIDs []ProductID
)

func (pids ProductIDs) ToIDs() []int64 {
	return ToIDs(pids)
}
func (pids ProductIDs) ToStrings() []string {
	return ToStrings(pids)
}

// NewProductID creates a new ProductID from a value.
func NewProductID(val any) ProductID {
	return New[productKind](val)
}

// NewProductIDValidated creates a new ProductID from a value and validates it.
func NewProductIDValidated(val any) (ProductID, error) {
	return NewValidated[productKind](val)
}

// orderKind marks a GID as a Shopify Order.
type orderKind struct{}

func (orderKind) name() string {
	return "Order"
}

type (
	// OrderID is the GID of a Shopify Order.
	OrderID = GID[orderKind]

	// OrderIDs is a slice of OrderID.
	OrderIDs []OrderID
)

func (oids OrderIDs) ToIDs() []int64 {
	return ToIDs(oids)
}
func (oids OrderIDs) ToStrings() []string {
	return ToStrings(oids)
}
