package sys_log

import (
	"bytes"
	"regexp"
	"time"
)

func ParseSyslog(message []byte) string {
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
	//Two type of time in syslog_maybe more type
	//2023-02-13T11:51:29.555Z
	//2023-02-13T18:51:29.509+07:00
	// layout1 := "2006-01-02T15:04:05.999-07:00"
	// layout2 := "2006-01-02T15:04:05.999Z"

	t1, err := time.Parse(time.RFC3339Nano, result["timestamp"])
	msg := result["timestamp"] + "," + result["source-address"] + "," + result["source-port"] + "," + result["nat-destination-address"] + "," + result["nat-destination-port"]
	if err == nil {
		msg = t1.Format("20060102150405") + "," + result["source-address"] + "," + result["source-port"] + "," + result["nat-destination-address"] + "," + result["nat-destination-port"]

	}
	t2, err := time.Parse(time.RFC3339Nano, result["timestamp"])
	if err == nil {
		msg = t2.Format("20060102150405") + "," + result["source-address"] + "," + result["source-port"] + "," + result["nat-destination-address"] + "," + result["nat-destination-port"]
	}
	// println(msg)
	return msg
}
