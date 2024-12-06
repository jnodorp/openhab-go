package openhab

import (
	"fmt"
	"strings"
	"time"
)

type OpenHABTime time.Time

const layout = "Jan 2, 2006, 3:04:05 PM"

// DeepCopyInto is a deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenHABTime) DeepCopyInto(out *OpenHABTime) {
	*out = *in
}

// DeepCopy is a deepcopy function, copying the receiver, creating a new OpenHABTime.
func (in *OpenHABTime) DeepCopy() *OpenHABTime {
	if in == nil {
		return nil
	}
	out := new(OpenHABTime)
	in.DeepCopyInto(out)
	return out
}

// UnmarshalJSON Parses the JSON string in the openHAB format
func (t *OpenHABTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	goT, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*t = OpenHABTime(goT)
	return nil
}

// MarshalJSON writes a quoted string in the openHAB format.
func (t OpenHABTime) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

// String returns the time in the openHAB format.
func (t *OpenHABTime) String() string {
	return fmt.Sprintf("%q", time.Time(*t).Format(layout))
}
