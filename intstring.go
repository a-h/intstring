package intstring

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type IntString int64

func (d IntString) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(d))
}

func (d *IntString) UnmarshalJSON(b []byte) (err error) {
	if bytes.HasPrefix(b, []byte{'"'}) {
		// It's a string.
		var s string
		err = json.Unmarshal(b, &s)
		if err != nil {
			return fmt.Errorf("IntString: failed to unmarshal as string: %w", err)
		}
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return fmt.Errorf("IntString: failed to parse string: %w", err)
		}
		*d = IntString(i)
		return nil
	}
	// It's an integer.
	var i int64
	err = json.Unmarshal(b, &i)
	if err != nil {
		return fmt.Errorf("IntString: failed to unmarshal int64: %w", err)
	}
	*d = IntString(i)
	return nil
}
