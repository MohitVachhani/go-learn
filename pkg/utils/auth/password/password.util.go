package password

import (
	"crypto/sha256"
	"encoding/base64"
)

//ConvertToEncryptedString is...
func ConvertToEncryptedString(input string) string {

	hasher := sha256.New()
	hasher.Write([]byte(input))
	encryptedString := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return encryptedString
}
