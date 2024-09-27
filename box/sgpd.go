package box

import (
	"fmt"
	"io"
)

type SgpdBox struct {
	Name       string
	Version    byte
	Flags      [3]byte
	notDecoded []byte
}

func DecodedSgpd(r io.Reader) (Box, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &SgpdBox{
		Name:       "sgpd",
		Version:    data[0],
		Flags:      [3]byte{data[1], data[2], data[3]},
		notDecoded: data[4:],
	}, nil
}

// Encode implements Box.
func (s *SgpdBox) Encode(w io.Writer) error {
	err := EncodeHeader(s, w)
	if err != nil {
		return err
	}
	buf := makebuf(s)
	buf[0] = s.Version
	buf[1], buf[2], buf[3] = s.Flags[0], s.Flags[1], s.Flags[2]
	copy(buf[4:], s.notDecoded)
	_, err = w.Write(buf)
	return err
}

// Size implements Box.
func (s *SgpdBox) Size() int {
	return BoxHeaderSize + 4 + len(s.notDecoded)
}

// Type implements Box.
func (s *SgpdBox) Type() string {
	return s.Name
}

func (s *SgpdBox) Dump() {
	fmt.Printf("Box %s box (size: %d)\n", s.Name, len(s.notDecoded))
}
