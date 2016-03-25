package trackr

import (
	"reflect"
	"testing"
)

func TestFindTrackingNumbers(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    []TrackingNumber
	}{
		{"empty string", "", []TrackingNumber{}},
		{"random word", "foo", []TrackingNumber{}},
		{"random words", "foo bar baz", []TrackingNumber{}},

		{"FedEx Express 12-digit", "012345678983", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"mangled FedEx Express 12-digit", "012345678980", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678981", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678982", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678984", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678985", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678986", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678987", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678988", []TrackingNumber{}},
		{"mangled FedEx Express 12-digit", "012345678989", []TrackingNumber{}},
		{"FedEx Express 12-digit with extra number", "0123456789831", []TrackingNumber{}},
		{"FedEx Express 12-digit with spaces", "  012345678983  ", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx Express 12-digit with words", "this 012345678983 thing", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx Express 12-digit with punctuation", "here:012345678983!!!", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx Express 12-digit with numbers", "123543 012345678983.763", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"multiple FedEx Express 12-digit", "012345678983 452520484968", []TrackingNumber{{"FedEx", "012345678983"}, {"FedEx", "452520484968"}}},

		{"FedEx Express 15-digit", "753840910347261", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"mangled FedEx Express 15-digit", "753840910347260", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347262", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347263", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347264", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347265", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347266", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347267", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347268", []TrackingNumber{}},
		{"mangled FedEx Express 15-digit", "753840910347269", []TrackingNumber{}},
		{"FedEx Express 15-digit with extra number", "7538409103472611", []TrackingNumber{}},
		{"FedEx Express 15-digit with spaces", "  753840910347261  ", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"FedEx Express 15-digit with words", "this 753840910347261 thing", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"FedEx Express 15-digit with punctuation", "here:753840910347261!!!", []TrackingNumber{{"FedEx", "753840910347261"}}},

		{"FedEx Express 9-digit", "555121207", []TrackingNumber{}},
		{"FedEx Express 18-digit", "90175491001234560", []TrackingNumber{}},

		{"FedEx Ground 15-digit", "987654312345672", []TrackingNumber{{"FedEx", "987654312345672"}}},
		{"FedEx Ground 15-digit", "366226074417244", []TrackingNumber{{"FedEx", "366226074417244"}}},
		{"mangled FedEx Ground 15-digit", "987654312345670", []TrackingNumber{}},
		// 987654312345671 is, in fact, a valid FedEx Express number
		{"mangled FedEx Ground 15-digit", "987654312345673", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345674", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345675", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345676", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345677", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345678", []TrackingNumber{}},
		{"mangled FedEx Ground 15-digit", "987654312345679", []TrackingNumber{}},
		{"FedEx Ground 15-digit with extra number", "3662260744172444", []TrackingNumber{}},
		{"FedEx Ground 15-digit with spaces", "  366226074417244  ", []TrackingNumber{{"FedEx", "366226074417244"}}},
		{"FedEx Ground 15-digit with words", "this 366226074417244 thing", []TrackingNumber{{"FedEx", "366226074417244"}}},
		{"FedEx Ground 15-digit with punctuation", "here:366226074417244!!!", []TrackingNumber{{"FedEx", "366226074417244"}}},
		{"multiple FedEx Ground 15-digits", "366226074417244 and 987654312345672?", []TrackingNumber{{"FedEx", "366226074417244"}, {"FedEx", "987654312345672"}}},

		{"UPS", "1Z0A19T50395201562", []TrackingNumber{{"UPS", "1Z0A19T50395201562"}}},
		{"mangled UPS", "1Z0A19T50395201560", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201561", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201563", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201564", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201565", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201566", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201567", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201568", []TrackingNumber{}},
		{"mangled UPS", "1Z0A19T50395201569", []TrackingNumber{}},
		{"UPS with extra number", "1Z0A19T503952015695", []TrackingNumber{}},
		{"UPS with extra letter", "1Z0A19T50395201569A", []TrackingNumber{}},
		{"UPS with spaces", "  1Z0A19T50395201562  ", []TrackingNumber{{"UPS", "1Z0A19T50395201562"}}},
		{"UPS with words", "this 1Z0A19T50395201562 thing", []TrackingNumber{{"UPS", "1Z0A19T50395201562"}}},
		{"UPS with punctuation", "here:1Z0A19T50395201562!!!", []TrackingNumber{{"UPS", "1Z0A19T50395201562"}}},
		{"multiple UPS", "1Z17RV550319270595 and 1ZW0W2150362177018?", []TrackingNumber{{"UPS", "1Z17RV550319270595"}, {"UPS", "1ZW0W2150362177018"}}},

		// from LR database
		{"USPS mod 10", "9505514863416085007661", []TrackingNumber{{"USPS", "9505514863416085007661"}}},
		{"USPS mod 10", "9500114129076084016042", []TrackingNumber{{"USPS", "9500114129076084016042"}}},
		{"USPS mod 10", "9114999944314135394161", []TrackingNumber{{"USPS", "9114999944314135394161"}}},
		{"USPS mod 10", "9405503699300135960087", []TrackingNumber{{"USPS", "9405503699300135960087"}}},
		{"USPS mod 10", "9400109699938317563628", []TrackingNumber{{"USPS", "9400109699938317563628"}}},
		{"USPS mod 10", "9405510200830034303190", []TrackingNumber{{"USPS", "9405510200830034303190"}}},
		{"USPS mod 10", "70150640000609285719", []TrackingNumber{{"USPS", "70150640000609285719"}}},
		{"USPS mod 10", "13152810000050524531", []TrackingNumber{{"USPS", "13152810000050524531"}}},

		// from https://github.com/franckverrot/activevalidators/blob/master/test/validations/tracking_number_test.rb
		{"USS39 mod 10", "EA123456784US", []TrackingNumber{{"USPS", "EA123456784US"}}},
		{"USS39 mod 11", "RB123456785US", []TrackingNumber{{"USPS", "RB123456785US"}}},
		{"20 character USS128 mod 10", "71123456789123456787", []TrackingNumber{{"USPS", "71123456789123456787"}}},
		{"20 character USS128 mod 10 ending in 0", "03110240000115809160", []TrackingNumber{{"USPS", "03110240000115809160"}}},
		{"22 character USS128 mod 10 ending in 0", "9171969010756003077385", []TrackingNumber{{"USPS", "9171969010756003077385"}}},

		{"USS39 tracking number with invalid check digit", "EA123456782US", []TrackingNumber{}},
		{"USS39 tracking number that is too short", "123456784US", []TrackingNumber{}},
		{"USS39 tracking number that is too long", "EAB123456784US", []TrackingNumber{}},
		{"USS39 tracking number with non-US product id", "EA123456784UT", []TrackingNumber{}},
		{"USS128 tracking number with invalid check-digit", "71123456789123456788", []TrackingNumber{}},
		{"USS128 tracking number that is too short", "7112345678912345678", []TrackingNumber{}},
		{"USS128 tracking number that is too long", "711234567891234567879287", []TrackingNumber{}},
		{"USS128 tracking number with invalid chars", "U11234567891234567879", []TrackingNumber{}},

		{"mangled USPS", "9114901496450568878985", []TrackingNumber{}},
	}
	for _, tt := range tests {
		if got := findTrackingNumbers(tt.message); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%s failed: findTrackingNumbers(%q) = %v, want %v", tt.name, tt.message, got, tt.want)
		}
	}
}
