Golang Decimal Package 
=======================

Golang Decimal Package providers a decimal type for multi-precision arithmetic.

Usage
=====

* create a decimal by float64:

		dec := decimal.New(5.18)
	
* create a decimal by string:
	
		dec, err := decimal.Parse("5.18181818181818181818181818181818181818")
		
		
COPYRIGHT & LICENSE
=====================

Copyright 2014 Landjur. Code released under the Apache License, Version 2.0.
