package prototype

import "testing"

// TestClone
func TestClone(t *testing.T) {
	// Get the ShirtCache instance.
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("GetShirtCloner() = nil")
	}

	// Clone Shirt as an ItemInfoGetter type.
	item, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item == whitePrototype {
		t.Error("item can't be equal to whitePrototype")
	}

	// Assert shirt from ItemInfoGetter to Shirt type.
	shirt, ok := item.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for shirt has failed")
	}

	shirt.SKU = "abc"
	// Create a second clone for white shirt.
	secondItem, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	secondShirt, ok := secondItem.(*Shirt)
	if !ok {
		t.Fatal("Type assertion for secondShirt has failed")
	}
	if shirt.SKU == secondShirt.SKU {
		t.Error("SKU's shirt and secondShirt must be different")
	}
	if shirt == secondShirt {
		t.Error("shirt object cannot be equal to secondShirt object")
	}

	// Log the information of each shirt.
	t.Logf("log: %s\n", shirt.GetInfo())
	t.Logf("log: %s\n", secondShirt.GetInfo())
	t.Logf(
		"log: shirt memory positon: %p, secondShirt memory postion %p",
		&shirt,
		&secondShirt,
	)
}
