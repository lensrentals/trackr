//go:generate ragel -Z -G2 -o tracking_numbers.go tracking_numbers.go.rl
// ^- generates tracking_numbers.go from Ragel source

package trackr
