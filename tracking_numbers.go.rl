package trackr

import "log"

type TrackingNumber struct {
	Carrier string
	Number  string
}

%%{
	machine tracking_number;
	write data;
}%%


func findTrackingNumbers(data string) []TrackingNumber {
	// Ragel counters
	cs, p, pe, eof := 0, 0, len(data), len(data)

	// Output
	found := make([]TrackingNumber, 0, 4)

	// Matcher states
	var wordStart int
	// FedEx Express
	var fe struct {
		sum int
		end int
	}

	// FedEx Ground
	var fg struct {
		sum int
		end int
	}

	// UPS
	var ups struct {
		sum int
		end int
	}

	// USPS uses several schemes, which we check simultaneously
	// First is USS128:
	var uss128 struct {
		sum int
		end int
	}

	// Second is USS39:
	var uss39 struct {
		sum10, sum11 int
		sumOK bool
		end int
	}

	%%{
		# FedEx Express uses a checksum that multiplies digits by 3 coefficients
		# Use one action for each type of digit
		action fe1 { fe.sum += 1*(int(fc) - '0') }
		action fe3 { fe.sum += 3*(int(fc) - '0') }
		action fe7 { fe.sum += 7*(int(fc) - '0') }
		action festart { fe.sum = 0 }
		action feend {
			if ((fe.sum-7*(int(fc) - '0')) % 11) % 10 == (int(fc) - '0') {
				fe.end = p+1
			}
		}

		# FedEx Express is either 12 or 15 digits, in the 317317317317[317] pattern
		fe = ( ( digit@fe3 digit@fe1 digit@fe7 ){4,5} @feend) >festart;


		# FedEx Ground uses a checksum that multiplies digits by 2 coefficients
		action fg1 { fg.sum += 1*(int(fc) - '0') }
		action fg3 { fg.sum += 3*(int(fc) - '0') }
		action fgstart { fg.sum = 0 }
		action fgend {
			if 10-(fg.sum % 10) == (int(fc) - '0') {
				fg.end = p+1
			}
		}

		# Yes, FedEx Ground is really 15 digits, just like FedEx Express
		fg = ( (digit@fg1 digit@fg3){7} digit@fgend) >fgstart;


		# UPS uses a checksum containing both alphabetic and numeric characters with two coefficients
		# Use one action for each {coefficient, chartype}
		action ups1n { ups.sum += 1*(int(fc) - '0') }
		action ups1a { ups.sum += 1*((int(fc) - '?') % 10) }
		action ups2n { ups.sum += 2*(int(fc) - '0') }
		action ups2a { ups.sum += 2*((int(fc) - '?') % 10) }
		action upsstart { ups.sum = 0 }
		action upsend {
			if 10-(ups.sum % 10) == (int(fc) - '0') {
				ups.end = p+1
			}
		}

		ups1 = (digit@ups1n) | ('A'..'Z'@ups1a);
		ups2 = (digit@ups2n) | ('A'..'Z'@ups2a);

		ups = ('1Z' (ups1 ups2){7} ups1 digit@upsend) >upsstart;


		# USPS uses a couple different checksum schemes
		# We can match them simultaneously
		# USS128 matches on either 20 or 22-digit strings:
		action uss128start { uss128.sum = 0 }
		action uss128a { uss128.sum += 3*(int(fc) - '0') }
		action uss128b { uss128.sum += 1*(int(fc) - '0') }
		action uss128end {
			if 10-(uss128.sum % 10) == (int(fc) - '0') {
				uss128.end = p+1
			}
		}

		uss128 = ((digit@uss128a digit@uss128b){9,10} digit@uss128a digit@uss128end) >uss128start;

		# USS39 matches sets of 8 digits, with a 2-alpha prefix and a mandatory "US" suffix,
		# and can use either of two checksums
		action uss39start { uss39.sum10 = 0; uss39.sum11 = 0; uss39.sumOK = false }
		action uss39d1 { uss39.sum11 += 8*(int(fc) - '0'); uss39.sum10 += 1*(int(fc) - '0') }
		action uss39d2 { uss39.sum11 += 6*(int(fc) - '0'); uss39.sum10 += 3*(int(fc) - '0') }
		action uss39d3 { uss39.sum11 += 4*(int(fc) - '0'); uss39.sum10 += 1*(int(fc) - '0') }
		action uss39d4 { uss39.sum11 += 2*(int(fc) - '0'); uss39.sum10 += 3*(int(fc) - '0') }
		action uss39d5 { uss39.sum11 += 3*(int(fc) - '0'); uss39.sum10 += 1*(int(fc) - '0') }
		action uss39d6 { uss39.sum11 += 5*(int(fc) - '0'); uss39.sum10 += 3*(int(fc) - '0') }
		action uss39d7 { uss39.sum11 += 9*(int(fc) - '0'); uss39.sum10 += 1*(int(fc) - '0') }
		action uss39d8 { uss39.sum11 += 7*(int(fc) - '0'); uss39.sum10 += 3*(int(fc) - '0') }
		action uss39check {
			{
				var checkDigit10, checkDigit11 int

				checkDigit10 = 10-(uss39.sum10 % 10)

				remainder := uss39.sum11 % 11
				if remainder == 0 {
					checkDigit11 = 5
				} else if remainder == 1 {
					checkDigit11 = 0
				} else {
					checkDigit11 = 11 - remainder
				}

				uss39.sumOK = (checkDigit10 == (int(fc) - '0')) || (checkDigit11 == (int(fc) - '0'))
			}
		}
		action uss39complete {
			if uss39.sumOK {
				uss39.end = p+1
			}
		}

		uss39 = (
			('A'..'Z'{2})
			digit@uss39d1
			digit@uss39d2
			digit@uss39d3
			digit@uss39d4
			digit@uss39d5
			digit@uss39d6
			digit@uss39d7
			digit@uss39d8
			digit@uss39check
			'US'@uss39complete
			) >uss39start;

		# Tracking numbers are any of our matchers
		tracking = fe | fg | ups | uss128 | uss39;

		# Match and emit on whole words only
		action start { wordStart = p }
		action emit {
			if fe.end == p {
				found = append(found, TrackingNumber{"FedEx", data[wordStart:p]})
			} else if fg.end == p { // don't emit both FedEx Express and FedEx Ground
				found = append(found, TrackingNumber{"FedEx", data[wordStart:p]})
			}

			if ups.end == p {
				found = append(found, TrackingNumber{"UPS", data[wordStart:p]})
			}

			if uss128.end == p {
				found = append(found, TrackingNumber{"USPS", data[wordStart:p]})
			} else if uss39.end == p {
				found = append(found, TrackingNumber{"USPS", data[wordStart:p]})
			}
		}

		# Words could be tracking numbers (ignoring errors) or just alphanumeric strings
		word = (tracking)>start%emit | (alnum+);

		# separators are non-alphanumeric strings
		separator = ^alnum+;

		main := separator? (word separator)* word?;

		write init;
		write exec;
	}%%

	log.Printf("%q => %v", data, found)

	return found
}
