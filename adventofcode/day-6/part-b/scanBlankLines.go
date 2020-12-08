package main

import "bytes"

// scanBlankLines splits the scanner input when a line is empty
func scanBlankLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	searchBytes := []byte("\n\n")

	if i := bytes.Index(data, searchBytes); i >= 0 {
		return i + 2, data[0:i], nil
	}

	if atEOF {
		return len(data), data, nil
	}

	return 0, nil, nil
}
