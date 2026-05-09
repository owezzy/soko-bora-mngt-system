package demo

import "testing"

func TestSpecIsDeterministicAcrossCalls(t *testing.T) {
	t.Parallel()

	first := Spec()
	second := Spec()

	if first.Customer.ID != second.Customer.ID || first.Customer.Name != second.Customer.Name || first.Customer.SMSNumber != second.Customer.SMSNumber {
		t.Fatal("expected demo customer spec to stay stable across calls")
	}

	if len(first.Stores) != 2 {
		t.Fatalf("expected 2 demo stores, got %d", len(first.Stores))
	}
	if len(second.Stores) != len(first.Stores) {
		t.Fatalf("expected stable store count, got %d and %d", len(first.Stores), len(second.Stores))
	}

	for i := range first.Stores {
		left := first.Stores[i]
		right := second.Stores[i]

		if left.ID != right.ID || left.Name != right.Name || left.Location != right.Location || left.Participating != right.Participating {
			t.Fatalf("expected store %d to be stable across calls", i)
		}
		if len(left.Products) != len(right.Products) {
			t.Fatalf("expected stable product count for store %s", left.ID)
		}

		for j := range left.Products {
			lp := left.Products[j]
			rp := right.Products[j]

			if lp != rp {
				t.Fatalf("expected product %d in store %s to be stable across calls", j, left.ID)
			}
		}
	}
}

func TestSpecReturnsFreshCopies(t *testing.T) {
	t.Parallel()

	first := Spec()
	first.Customer.Name = "Mutated"
	first.Stores[0].Name = "Changed Store"
	first.Stores[0].Products[0].Name = "Changed Product"

	second := Spec()

	if second.Customer.Name != "Demo Shopper" {
		t.Fatalf("expected fresh customer copy, got %q", second.Customer.Name)
	}
	if second.Stores[0].Name != "Fresh Harvest Grocers" {
		t.Fatalf("expected fresh store copy, got %q", second.Stores[0].Name)
	}
	if second.Stores[0].Products[0].Name != "Bananas" {
		t.Fatalf("expected fresh product copy, got %q", second.Stores[0].Products[0].Name)
	}
}
