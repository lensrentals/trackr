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

		{"FedEx 12-digit", "012345678983", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"mangled FedEx 12-digit", "012345678980", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678981", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678982", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678984", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678985", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678986", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678987", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678988", []TrackingNumber{}},
		{"mangled FedEx 12-digit", "012345678989", []TrackingNumber{}},
		{"FedEx 12-digit with extra number", "0123456789831", []TrackingNumber{}},
		{"FedEx 12-digit with spaces", "  012345678983  ", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx 12-digit with words", "this 012345678983 thing", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx 12-digit with punctuation", "here:012345678983!!!", []TrackingNumber{{"FedEx", "012345678983"}}},
		{"FedEx 12-digit with numbers", "123543 012345678983.763", []TrackingNumber{{"FedEx", "012345678983"}}},

		{"FedEx 15-digit", "753840910347261", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"mangled FedEx 15-digit", "753840910347260", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347262", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347263", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347264", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347265", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347266", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347267", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347268", []TrackingNumber{}},
		{"mangled FedEx 15-digit", "753840910347269", []TrackingNumber{}},
		{"FedEx 15-digit with extra number", "7538409103472611", []TrackingNumber{}},
		{"FedEx 15-digit with spaces", "  753840910347261  ", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"FedEx 15-digit with words", "this 753840910347261 thing", []TrackingNumber{{"FedEx", "753840910347261"}}},
		{"FedEx 15-digit with punctuation", "here:753840910347261!!!", []TrackingNumber{{"FedEx", "753840910347261"}}},

		{"FedEx 9-digit", "555121207", []TrackingNumber{}},
		{"FedEx 18-digit", "90175491001234560", []TrackingNumber{}},
	}
	for _, tt := range tests {
		if got := findTrackingNumbers(tt.message); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. findTrackingNumbers() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
