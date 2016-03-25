
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

	
//line tracking_numbers.go:68
	{
	cs = tracking_number_start
	}

//line tracking_numbers.go:73
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
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	}
	goto st_out
tr16:
//line tracking_numbers.go.rl:175

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
		
	goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line tracking_numbers.go:212
		if data[p] == 49 {
			goto tr2
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr3
		}
		goto st0
tr1:
//line tracking_numbers.go.rl:174
 wordStart = p 
//line tracking_numbers.go.rl:63
 fe.sum = 0 
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:77
 fg.sum = 0 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:110
 uss128.sum = 0 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line tracking_numbers.go:250
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr5:
//line tracking_numbers.go.rl:60
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line tracking_numbers.go:277
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr6
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr6:
//line tracking_numbers.go.rl:62
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line tracking_numbers.go:304
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr7
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr7:
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line tracking_numbers.go:331
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr8
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr8:
//line tracking_numbers.go.rl:60
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line tracking_numbers.go:358
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr9
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr9:
//line tracking_numbers.go.rl:62
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line tracking_numbers.go:385
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr10
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr10:
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line tracking_numbers.go:412
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr11
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr11:
//line tracking_numbers.go.rl:60
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line tracking_numbers.go:439
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr12:
//line tracking_numbers.go.rl:62
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line tracking_numbers.go:466
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr13:
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line tracking_numbers.go:493
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr14:
//line tracking_numbers.go.rl:60
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line tracking_numbers.go:520
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr15
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr15:
//line tracking_numbers.go.rl:62
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:64

			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line tracking_numbers.go:553
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr17
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr16
tr17:
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line tracking_numbers.go:580
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr18
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr18:
//line tracking_numbers.go.rl:60
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:76
 fg.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line tracking_numbers.go:607
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr19
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr19:
//line tracking_numbers.go.rl:62
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:64

			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
//line tracking_numbers.go.rl:78

			if 10-(fg.sum % 10) == (int(data[p]) - '0') {
				fg.end = p+1
			}
		
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line tracking_numbers.go:644
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr20
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr16
tr20:
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line tracking_numbers.go:667
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr21
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr21:
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line tracking_numbers.go:690
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr22:
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line tracking_numbers.go:713
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr23:
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line tracking_numbers.go:736
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr24:
//line tracking_numbers.go.rl:112
 uss128.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:113

			if 10-(uss128.sum % 10) == (int(data[p]) - '0') {
				uss128.end = p+1
			}
		
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line tracking_numbers.go:765
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr25
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr16
tr25:
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line tracking_numbers.go:788
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr26:
//line tracking_numbers.go.rl:113

			if 10-(uss128.sum % 10) == (int(data[p]) - '0') {
				uss128.end = p+1
			}
		
	goto st22
tr58:
//line tracking_numbers.go.rl:95

			if 10-(ups.sum % 10) == (int(data[p]) - '0') {
				ups.end = p+1
			}
		
	goto st22
tr70:
//line tracking_numbers.go.rl:150

			if uss39.sumOK {
				uss39.end = p+1
			}
		
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line tracking_numbers.go:831
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto tr16
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr2:
//line tracking_numbers.go.rl:174
 wordStart = p 
//line tracking_numbers.go.rl:63
 fe.sum = 0 
//line tracking_numbers.go.rl:61
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:77
 fg.sum = 0 
//line tracking_numbers.go.rl:75
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:94
 ups.sum = 0 
//line tracking_numbers.go.rl:110
 uss128.sum = 0 
//line tracking_numbers.go.rl:111
 uss128.sum += 3*(int(data[p]) - '0') 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line tracking_numbers.go:886
		if data[p] == 90 {
			goto st25
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr5
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr28
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr29
		}
		goto st0
tr28:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st26
tr29:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line tracking_numbers.go:934
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr30
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr31
		}
		goto st0
tr30:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st27
tr31:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line tracking_numbers.go:961
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr32
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr33
		}
		goto st0
tr32:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st28
tr33:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line tracking_numbers.go:988
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr34
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr35
		}
		goto st0
tr34:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st29
tr35:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line tracking_numbers.go:1015
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr36
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr37
		}
		goto st0
tr36:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st30
tr37:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line tracking_numbers.go:1042
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr38
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr39
		}
		goto st0
tr38:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st31
tr39:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line tracking_numbers.go:1069
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr41
		}
		goto st0
tr40:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st32
tr41:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line tracking_numbers.go:1096
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr43
		}
		goto st0
tr42:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st33
tr43:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line tracking_numbers.go:1123
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr44
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr45
		}
		goto st0
tr44:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st34
tr45:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line tracking_numbers.go:1150
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr47
		}
		goto st0
tr46:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st35
tr47:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line tracking_numbers.go:1177
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr49
		}
		goto st0
tr48:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st36
tr49:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line tracking_numbers.go:1204
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr51
		}
		goto st0
tr50:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st37
tr51:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line tracking_numbers.go:1231
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr53
		}
		goto st0
tr52:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st38
tr53:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line tracking_numbers.go:1258
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr54
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr55
		}
		goto st0
tr54:
//line tracking_numbers.go.rl:92
 ups.sum += 2*(int(data[p]) - '0') 
	goto st39
tr55:
//line tracking_numbers.go.rl:93
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line tracking_numbers.go:1285
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr56
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto tr57
		}
		goto st0
tr56:
//line tracking_numbers.go.rl:90
 ups.sum += 1*(int(data[p]) - '0') 
	goto st40
tr57:
//line tracking_numbers.go.rl:91
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line tracking_numbers.go:1312
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr58
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr3:
//line tracking_numbers.go.rl:174
 wordStart = p 
//line tracking_numbers.go.rl:123
 uss39.sum10 = 0; uss39.sum11 = 0; uss39.sumOK = false 
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line tracking_numbers.go:1337
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st42
		}
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr60
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr60:
//line tracking_numbers.go.rl:124
 uss39.sum11 += 8*(int(data[p]) - '0'); uss39.sum10 += 1*(int(data[p]) - '0') 
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line tracking_numbers.go:1378
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr61
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr61:
//line tracking_numbers.go.rl:125
 uss39.sum11 += 6*(int(data[p]) - '0'); uss39.sum10 += 3*(int(data[p]) - '0') 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line tracking_numbers.go:1401
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr62
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr62:
//line tracking_numbers.go.rl:126
 uss39.sum11 += 4*(int(data[p]) - '0'); uss39.sum10 += 1*(int(data[p]) - '0') 
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line tracking_numbers.go:1424
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr63
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr63:
//line tracking_numbers.go.rl:127
 uss39.sum11 += 2*(int(data[p]) - '0'); uss39.sum10 += 3*(int(data[p]) - '0') 
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line tracking_numbers.go:1447
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr64
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr64:
//line tracking_numbers.go.rl:128
 uss39.sum11 += 3*(int(data[p]) - '0'); uss39.sum10 += 1*(int(data[p]) - '0') 
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line tracking_numbers.go:1470
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr65
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr65:
//line tracking_numbers.go.rl:129
 uss39.sum11 += 5*(int(data[p]) - '0'); uss39.sum10 += 3*(int(data[p]) - '0') 
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line tracking_numbers.go:1493
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr66
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr66:
//line tracking_numbers.go.rl:130
 uss39.sum11 += 9*(int(data[p]) - '0'); uss39.sum10 += 1*(int(data[p]) - '0') 
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line tracking_numbers.go:1516
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr67
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr67:
//line tracking_numbers.go.rl:131
 uss39.sum11 += 7*(int(data[p]) - '0'); uss39.sum10 += 3*(int(data[p]) - '0') 
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line tracking_numbers.go:1539
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr68
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
tr68:
//line tracking_numbers.go.rl:132

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

				uss39.sumOK = (checkDigit10 == (int(data[p]) - '0')) || (checkDigit11 == (int(data[p]) - '0'))
			}
		
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line tracking_numbers.go:1579
		if data[p] == 85 {
			goto st52
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
		}
		goto st0
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
		if data[p] == 83 {
			goto tr70
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st23
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st23
			}
		default:
			goto st23
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
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 12, 15, 20, 22:
//line tracking_numbers.go.rl:175

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
		
//line tracking_numbers.go:1694
		}
	}

	}

//line tracking_numbers.go.rl:203


	log.Printf("%q => %v", data, found)

	return found
}
