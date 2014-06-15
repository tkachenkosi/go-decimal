// Copyright 2014 Landjur. All rights reserved.

package decimal

import (
	"testing"
)

func TestDecimal(t *testing.T) {
	a1 := New(1000.1234567890)
	a2 := New(-1000.1234567890)
	a3 := New(0.1234567890)
	a4 := New(-0.1234567890)
	if a1.String() != "1000.123456789" {
		t.Fatal("create instance of decimal failed")
	}
	if a2.String() != "-1000.123456789" {
		t.Fatal("create instance of decimal failed")
	}
	if a3.String() != "0.123456789" {
		t.Fatal("create instance of decimal failed")
	}
	if a4.String() != "-0.123456789" {
		t.Fatal("create instance of decimal failed")
	}

	// Cmp
	b1 := New(100.1234567890)
	b2 := New(00000100.12345678900000)
	if b1.Cmp(b2) != 0 {
		t.Fatal("Cmp error")
	}

	// Add
	c1 := New(100.1234567890)
	c2 := New(100.9876543210)
	c3 := New(201.1111111100)
	if c1.Add(c2).Cmp(c3) != 0 {
		t.Fatal("Add error")
	}

	// Sub
	d1 := New(100.987654321)
	d2 := New(100.88888888800000)
	d3 := New(0.0987654330000000000000000000000000000000000)
	if d1.Sub(d2).Cmp(d3) != 0 {
		t.Fatalf("Sub error")
	}

	// Mul
	e1 := New(1.23456789)
	e2 := New(9.87654321)
	if e1.Mul(e1).Mul(e2).String() != "15.053411111487447638891241" {
		t.Fatalf("Mul error")
	}
}
