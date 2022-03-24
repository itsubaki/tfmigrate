package history

import (
	"encoding/json"
	"fmt"
)

// FileHeader contains a meta data for file format.
type FileHeader struct {
	// Version is a file format version.
	Version int `json:"version"`
}

// ParseHistoryFile parses bytes and returns a History instance.
func ParseHistoryFile(b []byte) (*History, error) {
	version, err := detectHistoryFileVersion(b)
	if err != nil {
		return nil, err
	}

	switch version {
	case 1:
		return parseHistoryFileV1(b)

	default:
		return nil, fmt.Errorf("unknown history file version: %d", version)
	}
}

// detectHistoryFileVersion detects a file format version.
func detectHistoryFileVersion(b []byte) (int, error) {
	// peek a file header
	var header FileHeader
	err := json.Unmarshal(b, &header)
	if err != nil {
		return 0, err
	}

	return header.Version, nil
}
