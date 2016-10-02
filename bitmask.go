package ecs

// bitCounter this is very fast func to counts the number of 1 bits in uint32
// instead of doing o(32) we are doing only o(12)
func bitCounter(val uint32) uint32 {
	var count uint32
	count = val - ((val >> 1) & 033333333333) - ((val >> 2) & 011111111111)
	return ((count + (count >> 3)) & 030707070707) % 63
}

// calcPosition simply counts the number of ones from right to left.
// as an example, `combined` is `uint32` number which being AND with mask.
// mask is a uint32 power of 2 number which tries to count the number of
// one bits in combined varibale.
// so, let's say
// combined = 01101101
//     mask = 00001000
// -------------------
//                   3
// this is useful when you have an array of object and you want to find
// a specific one. rather than going through the array every time, you can
// find the index of that specific item in constant time.
func calcBitIndex(all, target uint32) uint32 {
	if all&target == 0 {
		return 0
	}

	var targetAllRightBits uint32

	//this is a trick that only works on power of 2 numbers.
	//for example 32 representation is 000100000. now in order to make only the right side
	//of one ones is double the target and subtracting by 1. so if you do that, 64 - 1 => 000011111
	//now when you all this with rightBits, you will get the current index.
	targetAllRightBits = 2*target - 1
	targetAllRightBits &= all

	return bitCounter(targetAllRightBits)
}
