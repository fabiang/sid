package sid

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const (
	MASK_8_BIT  = 0xff
	MASK_32_BIT = 0xffffffff
	MASK_48_BIT = 0xffffffffffff
)

func ConvertToString(input []byte) (string, error) {
	length := len(input)

	if len(input) < 8 {
		return "", fmt.Errorf("Byte string given must have at least 8 bytes, but got only %d bytes", length)
	}

	revision := input[0] & MASK_8_BIT
	numberOfSubAuthorityParts := int(input[1] & MASK_8_BIT)

	check := 8 + numberOfSubAuthorityParts*4
	if length != check {
		return "", fmt.Errorf("According to byte 1 of the SID it total length should be %d bytes, however its actual length is %d bytes", check, length)
	}

	buf := bytes.NewReader(input)

	var authority uint64
	err := binary.Read(buf, binary.BigEndian, &authority)
	if err != nil {
		return "", err
	}

	var subAuthority = make([]uint32, numberOfSubAuthorityParts)
	for i := 0; i < numberOfSubAuthorityParts; i++ {
		err = binary.Read(buf, binary.LittleEndian, &subAuthority[i])
		if err != nil {
			return "", err
		}
	}

	stringValue := fmt.Sprintf("S-%d-%d", revision, authority&MASK_48_BIT)
	for j := 0; j < numberOfSubAuthorityParts; j++ {
		stringValue += fmt.Sprintf("-%d", subAuthority[j]&MASK_32_BIT)
	}
	return stringValue, nil
}
