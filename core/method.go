package core

const (
	GET  = "get"
	SET  = "set"
	DEL  = "DEL"
	KEYS = "KEYS"
)

var METHOD_NAME_BYTE_MAP = map[string]byte{
	GET:  0x01,
	SET:  0x02,
	DEL:  0x03,
	KEYS: 0x04,
}
