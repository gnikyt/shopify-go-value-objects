package gid

import (
	"encoding/json"
	"testing"
)

func TestID(t *testing.T) {
	if got := NewCustomerID(123).ID(); got != 123 {
		t.Errorf("ID() = %d, want 123", got)
	}
	if got := NewCustomerID(0).ID(); got != 0 {
		t.Errorf("ID() = %d, want 0", got)
	}
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		gid  interface{ String() string }
		want string
	}{
		{"customer", NewCustomerID(123), "gid://shopify/Customer/123"},
		{"product", NewProductID(456), "gid://shopify/Product/456"},
		{"order", New[orderKind](789), "gid://shopify/Order/789"},
		{"zero", NewCustomerID(0), "gid://shopify/Customer/0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gid.String(); got != tt.want {
				t.Errorf("String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		gid  CustomerID
		want bool
	}{
		{"positive", NewCustomerID(1), true},
		{"zero", NewCustomerID(0), false},
		{"negative", New[customerKind](-1), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.gid.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	a := NewCustomerID(5)
	b := NewCustomerID(5)
	c := NewCustomerID(6)
	if !a.Equal(b) {
		t.Errorf("Equal() = false for identical GIDs, want true")
	}
	if a.Equal(c) {
		t.Errorf("Equal() = true for differing GIDs, want false")
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		val  any
		want int64
	}{
		{"int", 123, 123},
		{"int64", int64(123), 123},
		{"full gid string", "gid://shopify/Customer/123", 123},
		{"bare numeric string", "123456789", 123456789},
		{"wrong kind string yields zero", "gid://shopify/Product/123", 0},
		{"malformed string yields zero", "not-a-gid", 0},
		{"unsupported type yields zero", 1.5, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomerID(tt.val).ID(); got != tt.want {
				t.Errorf("New(%v) = %d, want %d", tt.val, got, tt.want)
			}
		})
	}
}

func TestNewValidated(t *testing.T) {
	tests := []struct {
		name    string
		val     any
		wantID  int64
		wantErr bool
	}{
		{"valid int", 123, 123, false},
		{"valid full gid", "gid://shopify/Customer/123", 123, false},
		{"valid bare numeric", "123", 123, false},
		{"zero is invalid", 0, 0, true},
		{"negative is invalid", -1, -1, true},
		{"wrong kind", "gid://shopify/Product/123", 0, true},
		{"malformed", "gid://shopify/Customer", 0, true},
		{"non-numeric id", "gid://shopify/Customer/abc", 0, true},
		{"wrong scheme", "urn://shopify/Customer/123", 0, true},
		{"wrong host segment", "gid://example/Customer/123", 0, true},
		{"unsupported type", 1.5, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewCustomerIDValidated(tt.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewValidated(%v) err = %v, wantErr %v", tt.val, err, tt.wantErr)
			}
			if got.ID() != tt.wantID {
				t.Errorf("NewValidated(%v) = %d, want %d", tt.val, got.ID(), tt.wantID)
			}
		})
	}
}

func TestMarshalJSON(t *testing.T) {
	got, err := json.Marshal(NewCustomerID(123))
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	if want := `"gid://shopify/Customer/123"`; string(got) != want {
		t.Errorf("Marshal() = %s, want %s", got, want)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		wantID  int64
		wantErr bool
	}{
		{"full gid", `"gid://shopify/Customer/123"`, 123, false},
		{"bare numeric", `"123"`, 123, false},
		{"wrong kind", `"gid://shopify/Product/123"`, 0, true},
		{"malformed gid", `"gid://shopify/Customer"`, 0, true},
		{"non-string json", `123`, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var g CustomerID
			err := json.Unmarshal([]byte(tt.data), &g)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unmarshal(%s) err = %v, wantErr %v", tt.data, err, tt.wantErr)
			}
			if !tt.wantErr && g.ID() != tt.wantID {
				t.Errorf("Unmarshal(%s) = %d, want %d", tt.data, g.ID(), tt.wantID)
			}
		})
	}
}

func TestMarshalRoundTrip(t *testing.T) {
	orig := NewProductID(987654321)
	data, err := json.Marshal(orig)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}
	var back ProductID
	if err := json.Unmarshal(data, &back); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if !orig.Equal(back) {
		t.Errorf("round trip: got %v, want %v", back, orig)
	}
}

func TestToIDs(t *testing.T) {
	gids := CustomerIDs{NewCustomerID(1), NewCustomerID(2), NewCustomerID(3)}
	got := gids.ToIDs()
	want := []int64{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("ToIDs() len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("ToIDs()[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestToIDsEmpty(t *testing.T) {
	if got := ToIDs(CustomerIDs{}); len(got) != 0 {
		t.Errorf("ToIDs(empty) = %v, want empty", got)
	}
}

func TestSliceHelpers(t *testing.T) {
	cids := CustomerIDs{NewCustomerID(1)}
	if got := cids.ToStrings(); len(got) != 1 || got[0] != "gid://shopify/Customer/1" {
		t.Errorf("CustomerIDs.ToStrings() = %v", got)
	}
	pids := ProductIDs{NewProductID(2)}
	if got := pids.ToIDs(); len(got) != 1 || got[0] != 2 {
		t.Errorf("ProductIDs.ToIDs() = %v", got)
	}
	oids := OrderIDs{New[orderKind](3)}
	if got := oids.ToIDs(); len(got) != 1 || got[0] != 3 {
		t.Errorf("OrderIDs.ToIDs() = %v", got)
	}
	if got := oids.ToStrings(); len(got) != 1 || got[0] != "gid://shopify/Order/3" {
		t.Errorf("OrderIDs.ToStrings() = %v", got)
	}
}

func TestNewProductIDValidated(t *testing.T) {
	if _, err := NewProductIDValidated(5); err != nil {
		t.Errorf("NewProductIDValidated(5) err = %v, want nil", err)
	}
	if _, err := NewProductIDValidated(0); err == nil {
		t.Errorf("NewProductIDValidated(0) err = nil, want error")
	}
}

func TestToStrings(t *testing.T) {
	gids := ProductIDs{NewProductID(1), NewProductID(2)}
	got := gids.ToStrings()
	want := []string{"gid://shopify/Product/1", "gid://shopify/Product/2"}
	if len(got) != len(want) {
		t.Fatalf("ToStrings() len = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("ToStrings()[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
