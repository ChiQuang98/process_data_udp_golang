package sys_log

import (
	"bytes"
	"regexp"
)

func ParseSyslog(message []byte) map[string]string {
	result := make(map[string]string)

	// Extract the priority, timestamp, and hostname
	message = bytes.TrimPrefix(message, []byte("<"))
	parts := bytes.SplitN(message, []byte(">"), 2)
	result["priority"] = string(parts[0])
	message = parts[1]

	parts = bytes.SplitN(message, []byte(" "), 2)
	result["level"] = string(parts[0])
	message = parts[1]

	parts = bytes.SplitN(message, []byte(" "), 2)
	result["timestamp"] = string(parts[0])

	message = parts[1]
	// Extract the remaining information from the message
	re := regexp.MustCompile(`([\w-]+="[\w.:/]+")`)

	matches := re.FindAll(message, -1)
	for _, match := range matches {
		keyValue := bytes.SplitN(match, []byte("="), 2)
		key := string(bytes.Trim(keyValue[0], `"`))
		value := string(bytes.Trim(keyValue[1], `"`))
		result[key] = value
	}
	return result
}
