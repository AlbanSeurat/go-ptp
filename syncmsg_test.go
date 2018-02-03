package ptp

import (
	"bytes"
	"testing"
	"time"
)

func TestMarshalSync(t *testing.T) {

	var tests = []struct {
		desc string
		m    *SyncMsg
		b    []byte
		err  error
	}{
		{
			desc: "Correct structure",
			m: &SyncMsg{
				Header: Header{
					MessageType:      SyncMsgType,
					CorrectionNs:     0,
					CorrectionSubNs:  0,
					ClockIdentity:    0x000af7fffe42a753,
					PortID:           2,
					SequenceID:       55330,
					LogMessagePeriod: -4,
				},
				OriginTimestamp: time.Unix(500, 200),
			},
			b: append([]byte{0x0, 0x2, 0x0, 0x2c, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0,
				0x0, 0x0, 0x0, 0x0,
				0x0, 0xa, 0xf7, 0xff, 0xfe, 0x42, 0xa7, 0x53, 0x0, 0x2, 0xd8, 0x22, 0x0, 0xfc,
				0x0, 0x0, 0x0, 0x0, 0x1, 0xf4, 0x0, 0x0, 0x0, 0xc8}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			b, err := tt.m.MarshalBinary()
			if err != nil {
				if want, got := tt.err, err; want != got {
					t.Fatalf("unexpected error: %v != %v", want, got)
				}

				return
			}

			if want, got := tt.b, b; !bytes.Equal(want, got) {
				t.Fatalf("unexpected Frame bytes:\n- want: %#v\n-  got: %#v", want, got)
			}
		})
	}
}

func TestUnmarshalSync(t *testing.T) {
}
