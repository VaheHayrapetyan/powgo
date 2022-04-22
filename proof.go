package powgo

import "encoding/base64"

// Proof structure
type Proof struct {
	buf []byte
}

func (proof Proof) marshalText() []byte {
	return []byte(base64.RawStdEncoding.EncodeToString(proof.buf))
}

func (proof *Proof) unmarshalText(buf []byte) error {
	var err error
	proof.buf, err = base64.RawStdEncoding.DecodeString(string(buf))
	return err
}

// Verify whether the proof is ok
func (proof *Proof) verify(challenge Challenge, data []byte) bool {
	switch challenge.Alg {
	case Sha111:
		return verifySha111(proof.buf, challenge.Nonce, data, challenge.Difficulty)
	default:
		panic("No such algorithm")
	}
}
