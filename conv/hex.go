package conv

import "encoding/hex"

func ToHexString(value string) string {
	return hex.EncodeToString([]byte(value))
}

func FromHexString(value string) (string, error) {
	bytes, err := hex.DecodeString(value)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}
