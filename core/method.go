package core

const (
	GET  = "GET"
	SET  = "SET"
	DEL  = "DEL"
	KEYS = "KEYS"
)

var MethodNameByteMap = map[string]byte{
	GET:  0x01,
	SET:  0x02,
	DEL:  0x03,
	KEYS: 0x04,
}
