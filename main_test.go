package main

import (
	"fmt"
	"testing"
)

func TestReplaceCorrectKeyWhenKeyIsAccessed(t *testing.T) {
	c := New(3)
	c.Set("a", 3)
	c.Set("b", 2)
	c.Set("c", 1)
	c.Get("a")
	c.Get("c")
	c.Get("a")
	c.Get("b")
	c.Set("d", 4)

	expectedValues := map[string]any{"a": 3, "b": 2, "d": 4}
	removedKey := "c"
	for i := range expectedValues {
		if _, ok := c.data[i]; !ok {
			t.Errorf("expected key %s, found nil.", i)
			for i := range c.data {
				fmt.Println(i)
			}

			return
		}
		if c.data[i].value != expectedValues[i] {
			t.Errorf("Expected %v, got %v", expectedValues[i], c.data[i].value)
		}
	}

	if _, ok := c.data[removedKey]; ok {
		t.Errorf("expected key a to be removed and it was not")
	}

	if len(c.data) != 3 {
		t.Errorf("expected to have 3 data found %d", len(c.data))
	}
}

func TestGetNotfoundValue(t *testing.T) {
	c := New(3)
	c.Set("a", 3)
	c.Set("b", 2)
	c.Set("c", 1)
	if value := c.Get("a"); value != 3 {
		t.Errorf("expected to get 3, got %v", value)
	}
	if value := c.Get("a"); value != 3 {
		t.Errorf("expected to get 3, got %v", value)
	}
	if value := c.Get("c"); value != 1 {
		t.Errorf("expected to get 1, got %v", value)
	}
	c.Set("d", 4)
	if value := c.Get("d"); value != 4 {
		t.Errorf("expected to get d, got %v", value)
	}
}

func TestGetCorrectValues(t *testing.T) {
	c := New(3)
	c.Set("a", 3)
	c.Set("b", 2)
	c.Set("c", 1)
	if value := c.Get("x"); value != -1 {
		t.Errorf("expected to get -1, got %v", value)
	}
}

func TestWithSameKeyDifferentValue(t *testing.T) {
	c := New(3)
	c.Set("a", 3)
	c.Set("a", 2)
	c.Set("a", 1)

	if value := c.Get("a"); value != 1 {
		t.Errorf("expected to get 1, got %v", value)
	}

	if len(c.data) != 1 {
		t.Errorf("expected to have 1 data, found %d", len(c.data))
	}

	if lenItems(c) != 1 {
		t.Errorf("expected to have 1 item, found %d", lenItems(c))
	}
}

func lenItems(c *Cache) int {
	count := 0
	for i := c.head; i != nil; i = i.next {
		// fmt.Println(i.key)
		count++
	}
	return count
}
