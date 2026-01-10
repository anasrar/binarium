package buffer

import "io"

// NOTE: not compatible with java WriteUTF
func WriteString(writer io.Writer, str string) error {
	b := []byte(str)
	_, err := writer.Write(b)
	return err
}

// NOTE: not compatible with java ReadUTF
func ReadString(reader io.Reader, length int) (string, error) {
	b := make([]byte, length)
	if _, err := reader.Read(b); err != nil {
		return "", err
	}

	return string(b), nil
}
