package powgo

// Solve the proof of work challenge
func Solve(challenge []byte, data []byte) (proof []byte, err error) {
	var ch Challenge
	err = ch.unmarshalText(challenge)
	if err != nil {
		return nil, err
	}

	prf := ch.solve(data)

	return prf.marshalText(), nil
}
