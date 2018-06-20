package helloworld

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func Popcount1(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func Popcount2(x uint64) int {
	var count int = 0
	for x != 0 {
		count += int(x & 0x1)
		x >>= 1
	}
	return count
}

func Popcount3(x uint64) int {
	count := 0
	for x != 0 {
		count++
		x &= (x - 1)
	}
	return count
}

/*Popcount4 HAKMEM Popcount
Consider a 3 bit number as being
      4a+2b+c
if we shift it right 1 bit, we have
      2a+b
subtracting this from the original gives
      2a+b+c
if we shift the original 2 bits right we get
      a
and so with another subtraction we have
      a+b+c
which is the number of bits in the original number.

Suitable masking  allows the sums of  the octal digits  in a 32 bit  number to
appear in  each octal digit.  This  isn't much help  unless we can get  all of
them summed together.   This can be done by modulo  arithmetic (sum the digits
in a number by  molulo the base of the number minus  one) the old "casting out
nines" trick  they taught  in school before  calculators were  invented.  Now,
using mod 7 wont help us, because our number will very likely have more than 7
bits set.   So add  the octal digits  together to  get base64 digits,  and use
modulo 63.   (Those of you  with 64  bit machines need  to add 3  octal digits
together to get base512 digits, and use mod 511.)

This is HACKMEM 169, as used in X11 sources.
Source: MIT AI Lab memo, late 1970's.
*/
func Popcount4(n uint32) int {
	var tmp uint32
	tmp = n - ((n >> 1) & 033333333333) - ((n >> 2) & 011111111111)
	return int(((tmp + (tmp >> 3)) & 030707070707) % 63)
}

func Popcount5(n uint64) int {
	return Popcount4(uint32(n)) + Popcount4(uint32(n>>32))
}
