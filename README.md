`trackr`
=======

`trackr` pulls tracking numbers out of unstructured text.

	trackr.Find("do you know anything about 012345678983?")
	// => trackr.TrackingNumber{Carrier: "FedEx", Number: "012345678983"}

`trackr` knows enough about the various formats to calculate and validate their checksums, making it much more precise
than (say) a regexp matching `^1Z.+`. It uses [Ragel](http://www.colm.net/open-source/ragel/) to extract and validate
tracking numbers using a single pass over the input.

Status
------

This package recognizes tracking numbers from:

* FedEx (both 12 and 15 digit Express formats and the 15 digit Ground format)
* UPS
* USPS (both USS128 and mod10/mod11 USS39)

This covers the needs at [LensRentals](https://www.lensrentals.com/), where it's been running in production since
2016.

Contributing
------------

Most of the work is done in `tracking_numbers.go.rl`, a Ragel file that simultaneously evaluates words for their
potential to match every tracking number format and calculates their checksums as it goes.

If you need to change the Ragel code, make sure you have `ragel` installed, make your changes, then `go generate`
to recreate `tracking_numbers.go`. Add some tests, make sure they pass, then commit all three together.
