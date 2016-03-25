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
	// FedEx Express
	var fe struct {
		sum int
		start int
		end int
	}

	// FedEx Ground
	var fg struct {
		sum int
		start int
		end int
	}

	// UPS
	var ups struct {
		sum int
		start int
		end int
	}

	%%{
		# FedEx Express uses a checksum that multiplies digits by 3 coefficients
		# Use one action for each type of digit
		action fe1 { fe.sum += 1*(int(fc) - '0') }
		action fe3 { fe.sum += 3*(int(fc) - '0') }
		action fe7 { fe.sum += 7*(int(fc) - '0') }
		action festart { fe.start = p; fe.sum = 0 }
		action feend {
			if ((fe.sum-7*(int(fc) - '0')) % 11) % 10 == (int(fc) - '0') {
				fe.end = p+1
			}
		}
		action feemit {
			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		}

		# FedEx Express is either 12 or 15 digits, in the 317317317317[317] pattern
		fe = ( ( digit@fe3 digit@fe1 digit@fe7 ){4,5} @feend) >festart %feemit;


		# FedEx Ground uses a checksum that multiplies digits by 2 coefficients
		action fg1 { fg.sum += 1*(int(fc) - '0') }
		action fg3 { fg.sum += 3*(int(fc) - '0') }
		action fgstart { fg.start = p; fg.sum = 0 }
		action fgend {
			if 10-(fg.sum % 10) == (int(fc) - '0') {
				fg.end = p+1
			}
		}
		action fgemit {
			if fg.end > fg.start && !(fe.end > fe.start) { // because we can overlap with FedEx Express
				found = append(found, TrackingNumber{"FedEx", data[fg.start:fg.end]})
			}
		}

		# Yes, FedEx Ground is really 15 digits, just like FedEx Express
		fg = ( (digit@fg1 digit@fg3){7} digit@fgend) >fgstart %fgemit;


		# UPS uses a checksum containing both alphabetic and numeric characters with two coefficients
		# Use one action for each {coefficient, chartype}
		action ups1n { ups.sum += 1*(int(fc) - '0') }
		action ups1a { ups.sum += 1*((int(fc) - '?') % 10) }
		action ups2n { ups.sum += 2*(int(fc) - '0') }
		action ups2a { ups.sum += 2*((int(fc) - '?') % 10) }
		action upsstart { ups.start = p; ups.sum = 0 }
		action upsend {
			if 10-(ups.sum % 10) == (int(fc) - '0') {
				ups.end = p+1
			}
		}
		action upsemit {
			if ups.end > ups.start {
				found = append(found, TrackingNumber{"UPS", data[ups.start:ups.end]})
			}
		}

		ups1 = (digit@ups1n) | ('A'..'Z'@ups1a);
		ups2 = (digit@ups2n) | ('A'..'Z'@ups2a);

		ups = ('1Z' (ups1 ups2){7} ups1 digit@upsend) >upsstart %upsemit;

		# Tracking numbers are any of our matchers
		tracking = fe | fg | ups;

		# Words could be tracking numbers (ignoring errors) or just alphanumeric strings
		word = (tracking $lerr{}) | (alnum+);

		# separators are non-alphanumeric strings
		separator = ^alnum+;

		main := separator? (word separator)* word?;

		write init;
		write exec;
	}%%

	log.Printf("%q => %v", data, found)

	return found
}
