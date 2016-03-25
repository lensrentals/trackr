
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

	
//line tracking_numbers.go:56
	{
	cs = tracking_number_start
	}

//line tracking_numbers.go:61
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
	}
	goto st_out
tr0:
//line tracking_numbers.go.rl:113

	goto st0
tr15:
//line tracking_numbers.go.rl:57

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
//line tracking_numbers.go.rl:113

	goto st0
tr19:
//line tracking_numbers.go.rl:57

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
//line tracking_numbers.go.rl:76

			if fg.end > fg.start && !(fe.end > fe.start) { // because we can overlap with FedEx Express
				found = append(found, TrackingNumber{"FedEx", data[fg.start:fg.end]})
			}
		
//line tracking_numbers.go.rl:113

	goto st0
tr53:
//line tracking_numbers.go.rl:98

			if ups.end > ups.start {
				found = append(found, TrackingNumber{"UPS", data[ups.start:ups.end]})
			}
		
//line tracking_numbers.go.rl:113

	goto st0
	st0:
		if p++; p == pe {
			goto _test_eof0
		}
	st_case_0:
//line tracking_numbers.go:184
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
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr1:
//line tracking_numbers.go.rl:51
 fe.start = p; fe.sum = 0 
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:70
 fg.start = p; fg.sum = 0 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line tracking_numbers.go:216
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
//line tracking_numbers.go.rl:48
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line tracking_numbers.go:241
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
//line tracking_numbers.go.rl:50
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line tracking_numbers.go:266
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
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line tracking_numbers.go:291
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
//line tracking_numbers.go.rl:48
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line tracking_numbers.go:316
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
//line tracking_numbers.go.rl:50
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line tracking_numbers.go:341
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
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line tracking_numbers.go:366
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
//line tracking_numbers.go.rl:48
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line tracking_numbers.go:391
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
//line tracking_numbers.go.rl:50
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line tracking_numbers.go:416
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
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line tracking_numbers.go:441
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
//line tracking_numbers.go.rl:48
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line tracking_numbers.go:466
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr14
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr14:
//line tracking_numbers.go.rl:50
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:52

			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line tracking_numbers.go:497
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
		goto tr15
tr16:
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line tracking_numbers.go:522
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
//line tracking_numbers.go.rl:48
 fe.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:69
 fg.sum += 3*(int(data[p]) - '0') 
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line tracking_numbers.go:547
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr18
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr18:
//line tracking_numbers.go.rl:50
 fe.sum += 7*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:52

			if ((fe.sum-7*(int(data[p]) - '0')) % 11) % 10 == (int(data[p]) - '0') {
				fe.end = p+1
			}
		
//line tracking_numbers.go.rl:71

			if 10-(fg.sum % 10) == (int(data[p]) - '0') {
				fg.end = p+1
			}
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line tracking_numbers.go:582
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
		goto tr19
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
tr2:
//line tracking_numbers.go.rl:51
 fe.start = p; fe.sum = 0 
//line tracking_numbers.go.rl:49
 fe.sum += 3*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:70
 fg.start = p; fg.sum = 0 
//line tracking_numbers.go.rl:68
 fg.sum += 1*(int(data[p]) - '0') 
//line tracking_numbers.go.rl:92
 ups.start = p; ups.sum = 0 
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line tracking_numbers.go:631
		if data[p] == 90 {
			goto st18
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr4
			}
		case data[p] > 89:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr23
		}
		goto tr0
tr22:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st19
tr23:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line tracking_numbers.go:679
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr25
		}
		goto tr0
tr24:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st20
tr25:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line tracking_numbers.go:706
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr27
		}
		goto tr0
tr26:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st21
tr27:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line tracking_numbers.go:733
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr28
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr29
		}
		goto tr0
tr28:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st22
tr29:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line tracking_numbers.go:760
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr30
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr31
		}
		goto tr0
tr30:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st23
tr31:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line tracking_numbers.go:787
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr32
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr33
		}
		goto tr0
tr32:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st24
tr33:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line tracking_numbers.go:814
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr34
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr35
		}
		goto tr0
tr34:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st25
tr35:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line tracking_numbers.go:841
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr36
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr37
		}
		goto tr0
tr36:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st26
tr37:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line tracking_numbers.go:868
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr38
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr39
		}
		goto tr0
tr38:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st27
tr39:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line tracking_numbers.go:895
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr40
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr41
		}
		goto tr0
tr40:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st28
tr41:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line tracking_numbers.go:922
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr43
		}
		goto tr0
tr42:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st29
tr43:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line tracking_numbers.go:949
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr44
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr45
		}
		goto tr0
tr44:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st30
tr45:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line tracking_numbers.go:976
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr46
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr47
		}
		goto tr0
tr46:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st31
tr47:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line tracking_numbers.go:1003
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr48
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr49
		}
		goto tr0
tr48:
//line tracking_numbers.go.rl:90
 ups.sum += 2*(int(data[p]) - '0') 
	goto st32
tr49:
//line tracking_numbers.go.rl:91
 ups.sum += 2*((int(data[p]) - '?') % 10) 
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line tracking_numbers.go:1030
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto tr51
		}
		goto tr0
tr50:
//line tracking_numbers.go.rl:88
 ups.sum += 1*(int(data[p]) - '0') 
	goto st33
tr51:
//line tracking_numbers.go.rl:89
 ups.sum += 1*((int(data[p]) - '?') % 10) 
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line tracking_numbers.go:1057
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr52
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st16
			}
		default:
			goto st16
		}
		goto tr0
tr52:
//line tracking_numbers.go.rl:93

			if 10-(ups.sum % 10) == (int(data[p]) - '0') {
				ups.end = p+1
			}
		
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line tracking_numbers.go:1084
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
		goto tr53
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

	_test_eof: {}
	if p == eof {
		switch cs {
		case 12:
//line tracking_numbers.go.rl:57

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
		case 34:
//line tracking_numbers.go.rl:98

			if ups.end > ups.start {
				found = append(found, TrackingNumber{"UPS", data[ups.start:ups.end]})
			}
		
		case 0:
//line tracking_numbers.go.rl:113

		case 15:
//line tracking_numbers.go.rl:57

			if fe.end > fe.start {
				found = append(found, TrackingNumber{"FedEx", data[fe.start:fe.end]})
			}
		
//line tracking_numbers.go.rl:76

			if fg.end > fg.start && !(fe.end > fe.start) { // because we can overlap with FedEx Express
				found = append(found, TrackingNumber{"FedEx", data[fg.start:fg.end]})
			}
		
//line tracking_numbers.go:1168
		}
	}

	}

//line tracking_numbers.go.rl:122


	log.Printf("%q => %v", data, found)

	return found
}
