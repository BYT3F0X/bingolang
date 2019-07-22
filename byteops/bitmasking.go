package byteops
// QueryBit is a function to query the bit at position n inside a byte denoted by query
// this is accomplished by means of bitmasking so endianness plays a significant role in the function
// to help you determine if this function will behave as intended
// you may use the function endianness.DetermineEndianness()
// bits count from 0 to 7
func QueryBit(subject byte, query int) bool {
	if query > 7 {
		panic("trying to query a bit after the 8th")
	}
	var mask uint8
	switch query {
	case 0:
		mask = 1
	case 1:
		mask = 2
	case 2:
		mask = 4
	case 3:
		mask = 8
	case 4:
		mask = 16
	case 5:
		mask = 32
	case 6:
		mask = 64
	case 7:
		mask = 128
	default:
		mask = 0
	}
	if subject & mask != 0 {
		return true
	} else {
		return false
	}
}
