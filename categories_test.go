package iab

import "testing"

func TestGetByCatT3(t *testing.T) {
	iabmap := NewIABTaxononmyMap()
	obj, err := iabmap.GetByCategories([]string{"Telecommunications Industry"})
	if err != nil {
		t.Error(err)
	}
	t.Log(obj)
}

func TestGetByCatT2(t *testing.T) {
	iabmap := NewIABTaxononmyMap()
	obj, err := iabmap.GetByCategories([]string{"Events and Attractions", "Comedy Events"})
	if err != nil {
		t.Error(err)
	}
	t.Log(obj.Parent)
}
