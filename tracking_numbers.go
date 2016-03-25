
//line tracking_numbers.go.rl:1
package trackr

import "log"

type TrackingNumber struct {
	Carrier string
	Number  string
}


//line tracking_numbers.go:14
const tracking_number_start int = 0
const tracking_number_first_final int = 0
const tracking_number_error int = -1

const tracking_number_en_main int = 0


//line tracking_numbers.go.rl:13



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

	
//line tracking_numbers.go:44
	{
	cs = tracking_number_start
	}

//line tracking_numbers.go:49
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	}
	goto st_out
tr0:
//line tracking_numbers.go.rl:59

	goto st0
tr14:
//line tracking_numbers.go.rl:46

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
//line tracking_numbers.go.rl:59

	goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line tracking_numbers.go:110
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr1:
//line tracking_numbers.go.rl:39
 fe.start = p 
//line tracking_numbers.go.rl:37
 fe.sum += 3*(int(data[p]) - '0') 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line tracking_numbers.go:135
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr3
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr3:
//line tracking_numbers.go.rl:36
 fe.sum += 1*(int(data[p]) - '0') 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line tracking_numbers.go:158
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr4:
//line tracking_numbers.go.rl:38
 fe.sum += 7*(int(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line tracking_numbers.go:181
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr5:
//line tracking_numbers.go.rl:37
 fe.sum += 3*(int(data[p]) - '0') 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line tracking_numbers.go:204
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr6:
//line tracking_numbers.go.rl:36
 fe.sum += 1*(int(data[p]) - '0') 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line tracking_numbers.go:227
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr7
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr7:
//line tracking_numbers.go.rl:38
 fe.sum += 7*(int(data[p]) - '0') 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line tracking_numbers.go:250
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr8:
//line tracking_numbers.go.rl:37
 fe.sum += 3*(int(data[p]) - '0') 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line tracking_numbers.go:273
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr9:
//line tracking_numbers.go.rl:36
 fe.sum += 1*(int(data[p]) - '0') 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line tracking_numbers.go:296
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr10
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr10:
//line tracking_numbers.go.rl:38
 fe.sum += 7*(int(data[p]) - '0') 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line tracking_numbers.go:319
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr11:
//line tracking_numbers.go.rl:37
 fe.sum += 3*(int(data[p]) - '0') 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line tracking_numbers.go:342
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr12:
//line tracking_numbers.go.rl:36
 fe.sum += 1*(int(data[p]) - '0') 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line tracking_numbers.go:365
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr13:
//line tracking_numbers.go.rl:38
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:40

			log.Printf("feend");
			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line tracking_numbers.go:395
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr14
tr15:
//line tracking_numbers.go.rl:37
 fe.sum += 3*(int(data[p]) - '0') 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line tracking_numbers.go:418
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr16:
//line tracking_numbers.go.rl:36
 fe.sum += 1*(int(data[p]) - '0') 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line tracking_numbers.go:441
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr17:
//line tracking_numbers.go.rl:38
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:40

			log.Printf("feend");
			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line tracking_numbers.go:471
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr14
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st16
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto st0
	st_out:
	_test_eof0: cs = 0; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 12, 15:
//line tracking_numbers.go.rl:46

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
		case 0:
//line tracking_numbers.go.rl:59

//line tracking_numbers.go:535
		}
	}

	}

//line tracking_numbers.go.rl:68


	if cs < tracking_number_first_final {
		log.Printf("%q errored!", data)
	}

	if eof == pe {
		// unused, sure
	}

	log.Printf("%q => %v", data, found)

	return found
}
