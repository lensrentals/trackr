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
	var usps struct {
		mod10a int
		mod10b int
		mod11 int
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
		action uspsstart { usps.mod10a = 0; usps.mod10b = 0; usps.mod11 = 0 }
		action uspsMod10a { usps.mod10a += 3*(int(fc) - '0'); usps.mod10b += 1*(int(fc) - '0') }
		action uspsMod10b { usps.mod10a += 1*(int(fc) - '0'); usps.mod10b += 3*(int(fc) - '0') }
		action uspsMod10a_end {
			//log.Printf("A: mod10a: %d, mod10b: %d, fc: %d", usps.mod10a, usps.mod10b, (int(fc) - '0'))
			if 10-(usps.mod10a % 10) == (int(fc) - '0') {
				usps.end = p+1
			}
		}
		action uspsMod10b_end {
			//log.Printf("B: mod10a: %d, mod10b: %d, fc: %d", usps.mod10a, usps.mod10b, (int(fc) - '0'))
			if 10-(usps.mod10b % 10) == (int(fc) - '0') {
				usps.end = p+1
			}
		}

		# USPS mod10 counts from the right, but we don't have a length yet
		# Tally even-length and odd-length strings simultaneously using separate counters, and
		# figure it out which to compare against in the final digit action
		# Match any strings 17-24 characters with the right numeric properties
		uspsMod10 =
			(digit@uspsMod10a digit@uspsMod10b){8,11}
			((digit@uspsMod10a digit@uspsMod10b_end) | digit@uspsMod10a_end);

		# USPS is collectively any of the schemes, with the common entry
		usps = (uspsMod10) >uspsstart;

		# Tracking numbers are any of our matchers
		tracking = fe | fg | ups | usps;

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

			if usps.end == p {
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
