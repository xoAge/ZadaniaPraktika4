package doublyList

import (
	"reflect"
	"testing"
)

func TestAddAndValues(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("Expected %v, got %v", expected, list.Values())
	}
}

func TestGet(t *testing.T) {
	list := New[string]()
	list.Add("a")
	list.Add("b")
	list.Add("c")

	tests := []struct {
		index    int
		expected string
		ok       bool
	}{
		{0, "a", true},
		{1, "b", true},
		{2, "c", true},
		{3, "", false},
		{-1, "", false},
	}

	for _, tt := range tests {
		val, ok := list.Get(tt.index)
		if ok != tt.ok || (ok && val != tt.expected) {
			t.Errorf("Get(%d): expected (%v, %v), got (%v, %v)", tt.index, tt.expected, tt.ok, val, ok)
		}
	}
}

func TestRemove(t *testing.T) {
	list := New[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Удаление по первому индексу
	if !list.Remove(0) {
		t.Error("Failed to remove element at index 0")
	}
	expected := []int{2, 3}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("After removal at 0, expected %v, got %v", expected, list.Values())
	}

	// Удаление по последнему индексу
	if !list.Remove(1) {
		t.Error("Failed to remove element at index 1")
	}
	expected = []int{2}
	if !reflect.DeepEqual(list.Values(), expected) {
		t.Errorf("After removal at 1, expected %v, got %v", expected, list.Values())
	}

	// Попытка удаления по несуществующему индексу
	if list.Remove(5) {
		t.Error("Remove should return false for invalid index")
	}
}

func TestEdgeCases(t *testing.T) {
	list := New[int]()
	// Удаление из пустого списка
	if list.Remove(0) {
		t.Error("Remove should return false for empty list")
	}
	// Получение из пустого списка
	if v, ok := list.Get(0); ok {
		t.Errorf("Get from empty list should return false, got %v", v)
	}
}
