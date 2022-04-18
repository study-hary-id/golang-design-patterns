package prototype

import (
	"testing"
)

func TestCloneWhiteShirt(t *testing.T) {
	// Get the ShirtCache instance.
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("GetShirtCloner() = nil, shirtCache shouldn't be nil")
	}

	// Clone Shirt as an ItemInfoGetter type.
	item, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}
	if item == whitePrototype {
		t.Error("error: item can't be equal to whitePrototype")
	}

	// Assert shirt from ItemInfoGetter to Shirt type.
	shirt, ok := item.(*Shirt)
	if !ok {
		t.Fatal("error: type assertion for shirt has failed")
	}
	shirt.SKU = "abc"

	/* Second prototyping process */
	// Create a second clone for white shirt.
	secondItem, err := shirtCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	secondShirt, ok := secondItem.(*Shirt)
	if !ok {
		t.Fatal("error: type assertion for secondShirt has failed")
	}
	if shirt.SKU == secondShirt.SKU {
		t.Error("error: sku shirt and secondShirt must be different")
	}
	if shirt == secondShirt {
		t.Error("error: shirt instance cannot be equal to secondShirt")
	}

	// Log the information of each shirt.
	t.Logf("log: %s", shirt.GetInfo())
	t.Logf("log: %s", secondShirt.GetInfo())
	t.Logf(
		"log: shirt memory positon: %p, secondShirt memory postion %p",
		&shirt,
		&secondShirt,
	)
}

func TestCloneBlackShirt(t *testing.T) {
	// Get the ShirtCache instance.
	shirtCache := GetShirtCloner()
	if shirtCache == nil {
		t.Fatal("GetShirtCloner() = nil, shirtCache shouldn't be nil")
	}

	// Clone Shirt as an ItemInfoGetter type.
	item, err := shirtCache.GetClone(Black)
	if err != nil {
		t.Error(err)
	}
	if item == whitePrototype {
		t.Error("error: item can't be equal to whitePrototype")
	}

	// Assert shirt from ItemInfoGetter to Shirt type.
	shirt, ok := item.(*Shirt)
	if !ok {
		t.Fatal("error: type assertion for shirt has failed")
	}
	shirt.SKU = "xyz"

	/* Second prototyping process */
	// Create a second clone for white shirt.
	secondItem, err := shirtCache.GetClone(Black)
	if err != nil {
		t.Error(err)
	}

	secondShirt, ok := secondItem.(*Shirt)
	if !ok {
		t.Fatal("error: type assertion for secondShirt has failed")
	}
	if shirt.SKU == secondShirt.SKU {
		t.Error("error: sku shirt and secondShirt must be different")
	}
	if shirt == secondShirt {
		t.Error("error: shirt instance cannot be equal to secondShirt")
	}

	// Log the information of each shirt.
	t.Logf("log: %s", shirt.GetInfo())
	t.Logf("log: %s", secondShirt.GetInfo())
	t.Logf(
		"log: shirt memory positon: %p, secondShirt memory postion %p",
		&shirt,
		&secondShirt,
	)
}
