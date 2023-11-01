package redis

import "encoding"

type Cacher interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}
