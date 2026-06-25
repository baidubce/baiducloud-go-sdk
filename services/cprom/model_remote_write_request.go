package cprom

import "io"

type RemoteWriteRequest struct {
	RemoteWriteUrl  *string       `json:"-"`
	ContentType     *string       `json:"-"`
	ContentEncoding *string       `json:"-"`
	InstanceId      *string       `json:"-"`
	Authorization   *string       `json:"-"`
	Body            io.ReadCloser `json:"-"`
}
