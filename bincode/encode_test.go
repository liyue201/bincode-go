package bincode

import (
	"bytes"
	"encoding/hex"
	"testing"
)

type packetSend struct {
	dstEid     uint32
	receiver   [32]byte
	message    []byte
	options    []byte
	nativeFee  uint64
	lzTokenFee uint64
}

func TestEncode(t *testing.T) {
	p := packetSend{}

	dataHex := "e87500000000000000000000000000005d3a1ff2b6bab83b63cd9ad0787074081a52ef3428000000000000000000000000000000096a5455c4538aed284615078fd3b52e6f7a79b200000000178816d1160000000003010011010000000000000000000000000000fde8d6640500000000000000000000000000"

	data, _ := hex.DecodeString(dataHex)

	var buf bytes.Buffer
	buf.Write(data)

	p.dstEid, _ = DecodeU32(&buf)
	DecodeByteArray(&buf, p.receiver[:])
	p.message, _ = DecodeSlice(&buf)
	p.options, _ = DecodeSlice(&buf)
	p.nativeFee, _ = DecodeU64(&buf)
	p.lzTokenFee, _ = DecodeU64(&buf)

	t.Logf("%+v", p)

	encoded := MustEncode(p)

	t.Logf("encoded: %x", encoded)

	if dataHex != hex.EncodeToString(encoded) {
		t.Fatalf("encoding error")
	}
}
