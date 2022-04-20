package pow

// Verify the proof of work is Solved
func Verify(challenge []byte, proof []byte, data []byte) (ok bool, err error) {
	var ch Challenge
	var prf Proof
	err = ch.unmarshalText(challenge)
	if err != nil {
		return false, err
	}
	err = prf.unmarshalText(proof)
	if err != nil {
		return false, err
	}
	return prf.verify(ch, data), nil
}
