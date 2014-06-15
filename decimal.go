// Copyright 2014 Landjur. All rights reserved.

package decimal

import (
	"bytes"
	"errors"
	"io"
	"math/big"
	"strconv"
	"strings"
)

// New returns a new instance of Decimal.
func New(number float64) *Decimal {
	numberString := strconv.FormatFloat(number, 'f', -1, 64)
	decimal, _ := Parse(numberString)
	return decimal
}

// Parse returns a new instance of Decimal by parse decimal string.
func Parse(numberString string) (*Decimal, error) {
	var unscaledBuffer bytes.Buffer
	var scale int
	reader := strings.NewReader(numberString)
	index := 1
	for {
		ch, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		switch ch {
		case '+', '-':
			if index > 1 { // sign must be first character
				return nil, errors.New("decimal: invalid number string")
			}
			unscaledBuffer.WriteRune(ch)
		case '.':
			if scale != 0 {
				return nil, errors.New("decimal: invalid number string")
			}
			scale = len(numberString) - index
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			unscaledBuffer.WriteRune(ch)
		default:
			return nil, errors.New("decimal: invalid number string")
		}

		index++
	}

	integer, ok := new(big.Int).SetString(unscaledBuffer.String(), 10)
	if !ok {
		return nil, errors.New("decimal: invalid number string")
	}

	return &Decimal{
		integer: integer,
		scale:   scale,
	}, nil
}

// alignScale aligns the scale of two decimals.
func alignScale(a, b *Decimal) {
	switch {
	case a.scale < b.scale:
		a.integer.Mul(a.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(b.scale-a.scale)), nil))
		a.scale = b.scale
	case a.scale > b.scale:
		b.integer.Mul(b.integer, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(a.scale-b.scale)), nil))
		b.scale = a.scale
	}
}

// Decimal represents a decimal which can handing fixed precision.
type Decimal struct {
	integer *big.Int
	scale   int // scale represents the number of deciaml digits
}

// Cmp compares this and another and returns:
// -1 if this <  another
//  0 if this == another
// +1 if this >  another
func (this *Decimal) Cmp(another *Decimal) int {
	alignScale(this, another)
	return this.integer.Cmp(another.integer)
}

// Add sets this to the sum of this and another and returns this.
func (this *Decimal) Add(another *Decimal) *Decimal {
	alignScale(this, another)
	this.integer.Add(this.integer, another.integer)
	return this
}

// Sub sets this to the difference this-another and returns this.
func (this *Decimal) Sub(another *Decimal) *Decimal {
	alignScale(this, another)
	this.integer.Sub(this.integer, another.integer)
	return this
}

// Mul sets this to the product this*another and returns this.
func (this *Decimal) Mul(another *Decimal) *Decimal {
	this.integer.Mul(this.integer, another.integer)
	this.scale += another.scale
	return this
}

// Sign returns:
// -1: if this <  0
//  0: if this == 0
// +1: if this >  0
func (this Decimal) Sign() int {
	return this.integer.Sign()
}

// Float64 returns the nearest float64 value of decimal.
func (this Decimal) Float64() float64 {
	resultString := this.String()
	result, _ := strconv.ParseFloat(resultString, 64)
	return result
}

// FloatString returns a string representation of decimal form with precision digits of precision after the decimal point and the last digit rounded.
func (this Decimal) FloatString(precision int) string {
	x := new(big.Rat).SetInt(this.integer)
	y := new(big.Rat).Inv(new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(int64(10)), big.NewInt(int64(this.scale)), nil)))
	z := new(big.Rat).Mul(x, y)
	return z.FloatString(precision)
}

// String returns the string of Decimal.
func (this Decimal) String() string {
	unscaledString := strings.TrimLeft(this.integer.String(), "-")
	if this.scale == 0 {
		return unscaledString
	}

	pointIndex := len(unscaledString) - this.scale
	switch {
	case pointIndex < 0:
		if this.integer.Sign() == -1 {
			return "-0." + strings.Repeat("0", -1*pointIndex) + unscaledString
		}

		return "0." + strings.Repeat("0", -1*pointIndex) + unscaledString
	case pointIndex > 0:
		if this.integer.Sign() == -1 {
			return "-" + unscaledString[0:pointIndex] + "." + unscaledString[pointIndex:]
		}

		return unscaledString[0:pointIndex] + "." + unscaledString[pointIndex:]
	default: // pointIndex == 0
		if this.integer.Sign() == -1 {
			return "-0." + unscaledString
		}

		return "0." + unscaledString
	}
}
