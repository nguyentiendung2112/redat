package core

func encode(key string, value string) []uint8 {
	var key_size = len(key)
	var value_size = len(value)
	var result = make([]uint8, 16+key_size+value_size)
	result[0] = uint8(key_size)
	result[1] = uint8(value_size)

	return result

}
