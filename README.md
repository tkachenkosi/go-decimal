#Go-Decimal

The go-decimal package providers a Decimal type for multi-precision arithmetic.

##Usage

* create a Decimal instance by float64 or int64:

		dec := decimal.New(5.18)

        dec := new(decimal.Decimal).SetFloat(5.18)
	
        dec := new(decimal.Decimal).SetInt(5)

* create a Decimal instance by string:
	
		dec, err := decimal.Parse("5.18181818181818181818181818181818181818")
        
		dec, err := new(decimal.Decimal).SetString("5.18181818181818181818181818181818181818")
