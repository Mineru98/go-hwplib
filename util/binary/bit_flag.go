package binary

/**
 * 이진 연산을 하는 패키지
 *
 * @author Mineru98
 */

// 특정 비트가 1인지 확인하는 함수
func Get(mask uint64, position int) bool {
	mask2 := uint64(1) << position
	return (mask & mask2) == mask2
}

// 특정 비트를 설정하는 함수
func Set(mask uint64, position int, flag bool) uint64 {
	if flag {
		mask = mask | (1 << position)
	} else {
		if (mask & (1 << position)) != 0 {
			mask = mask ^ (1 << position)
		}
	}
	return mask
}

// start부터 end까지의 비트를 반환하는 함수
func GetRange(mask uint64, start, end int) uint64 {
	ret := mask >> start
	temp := uint64(0)
	for i := 0; i < end-start+1; i++ {
		temp = (temp << 1) + 1
	}
	return ret & temp
}

// start부터 end까지의 비트를 설정하는 함수
func SetRange(mask uint64, start, end, value int) uint64 {
	for position := start; position <= end; position++ {
		flag := Get(uint64(value), position-start)
		mask = Set(mask, position, flag)
	}
	return mask
}
