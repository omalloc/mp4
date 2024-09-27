package box

import (
	"fmt"
	"io"
	"log"
)

// Noop pDecoded Box (undefined box - optional)
//
// Status: not decoded

type NoopBox struct {
	Name       string
	Version    byte
	Flags      [3]byte
	notDecoded []byte
}

func DecodeAnyBox(name string) func(io.Reader) (Box, error) {
	return func(r io.Reader) (Box, error) {
		data, err := io.ReadAll(r)
		if err != nil {
			return nil, err
		}

		log.Printf("Decoding %s box (size: %d)", name, len(data))

		return &NoopBox{
			Name:       name,
			Version:    data[0],
			Flags:      [3]byte{data[1], data[2], data[3]},
			notDecoded: data[4:],
		}, nil
	}
}

func DecodedNoopBox(r io.Reader) (Box, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &NoopBox{
		Name:       "noop",
		Version:    data[0],
		Flags:      [3]byte{data[1], data[2], data[3]},
		notDecoded: data[4:],
	}, nil
}

func (b *NoopBox) Type() string {
	return b.Name
}

func (b *NoopBox) Size() int {
	return BoxHeaderSize + 4 + len(b.notDecoded)
}

func (b *NoopBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	buf := makebuf(b)
	buf[0] = b.Version
	buf[1], buf[2], buf[3] = b.Flags[0], b.Flags[1], b.Flags[2]
	copy(buf[4:], b.notDecoded)
	_, err = w.Write(buf)
	return err
}

func (b *NoopBox) Dump() {
	fmt.Printf("Box %s box (size: %d)\n", b.Name, len(b.notDecoded))
}
