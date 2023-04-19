package ips2crash

import (
	"bufio"
	"bytes"
)

func splitLines(fileBytes []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
	const maxCapacity = 1024 * 1024 * 3 // 3MB buffer
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
