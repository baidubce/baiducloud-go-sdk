package pfs

import "encoding/json"

// MarshalJSON always emits mountPath. The server declares mountPath as a required field
// and generates "/pfs/<instanceId>" when it is empty, so an unset value is sent as "".
func (b MountTargetBindConfig) MarshalJSON() ([]byte, error) {
	type plain MountTargetBindConfig
	out := plain(b)
	if out.MountPath == nil {
		empty := ""
		out.MountPath = &empty
	}
	return json.Marshal(out)
}
