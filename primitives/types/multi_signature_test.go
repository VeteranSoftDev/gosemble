package types

import (
	"bytes"
	"testing"

	sc "github.com/LimeChain/goscale"
	"github.com/stretchr/testify/assert"
)

func Test_EncodeMultiSignature(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       MultiSignature
		expectation []byte
	}{
		{
			label:       "Encode(MultiSignature(Ed25519))",
			input:       NewMultiSignatureEd25519(NewEd25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
			expectation: []byte{0x00, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64},
		},
		{
			label:       "Encode(MultiSignature(Sr25519))",
			input:       NewMultiSignatureSr25519(NewSr25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
			expectation: []byte{0x01, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64},
		},
		{
			label:       "Encode(MultiSignature(Ecdsa))",
			input:       NewMultiSignatureEcdsa(NewEcdsa(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x65}...)),
			expectation: []byte{0x02, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x65},
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}

			testExample.input.Encode(buffer)

			assert.Equal(t, testExample.expectation, buffer.Bytes())
		})
	}
}

func Test_DecodeMultiSignature(t *testing.T) {
	var testExamples = []struct {
		label       string
		input       []byte
		expectation MultiSignature
	}{
		{
			label:       "DecodeMultiSignature()",
			input:       []byte{0x00, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64},
			expectation: NewMultiSignatureEd25519(NewEd25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
		},
		{
			label:       "DecodeMultiSignature()",
			input:       []byte{0x01, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64},
			expectation: NewMultiSignatureSr25519(NewSr25519(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64}...)),
		},
		{
			label:       "DecodeMultiSignature()",
			input:       []byte{0x02, 0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x65},
			expectation: NewMultiSignatureEcdsa(NewEcdsa(sc.FixedSequence[sc.U8]{0x00, 0x62, 0x37, 0x61, 0x33, 0x63, 0x31, 0x32, 0x64, 0x63, 0x30, 0x63, 0x38, 0x63, 0x37, 0x34, 0x38, 0x61, 0x62, 0x30, 0x37, 0x35, 0x32, 0x35, 0x62, 0x37, 0x30, 0x31, 0x31, 0x32, 0x32, 0x62, 0x38, 0x38, 0x62, 0x64, 0x37, 0x38, 0x66, 0x36, 0x30, 0x30, 0x63, 0x37, 0x36, 0x33, 0x34, 0x32, 0x64, 0x32, 0x37, 0x66, 0x32, 0x35, 0x65, 0x35, 0x66, 0x39, 0x32, 0x34, 0x34, 0x34, 0x63, 0x64, 0x65}...)),
		},
	}

	for _, testExample := range testExamples {
		t.Run(testExample.label, func(t *testing.T) {
			buffer := &bytes.Buffer{}
			buffer.Write(testExample.input)

			result := DecodeMultiSignature(buffer)

			assert.Equal(t, testExample.expectation, result)
		})
	}
}
