package pow

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

// Challenge structure.
type Challenge struct {
	Alg        Algorithm // The requested algorithm
	Difficulty uint32    // The requested difficulty
	Nonce      []byte    // Nonce to diversify the challenge
}

// NewChallenge with algorithm Sha111: from Provider to Requester.
func NewChallenge(difficulty uint32, nonce []byte) (challenge []byte) {
	ch := Challenge{
		Difficulty: difficulty,
		Nonce:      nonce,
		Alg:        Sha111,
	}

	return ch.marshalText()
}

func (challenge Challenge) marshalText() []byte {
	return []byte(
		fmt.Sprintf("%s-%d-%s",
			challenge.Alg,
			challenge.Difficulty,
			base64.RawStdEncoding.EncodeToString(challenge.Nonce),
		),
	)
}

func (challenge *Challenge) unmarshalText(buf []byte) error {
	bits := strings.SplitN(string(buf), "-", 3)
	if len(bits) != 3 {
		return fmt.Errorf("there should be two dashes in a PoW request")
	}
	alg := Algorithm(bits[0])
	if alg != Sha111 {
		return fmt.Errorf("%s: unsupported algorithm", bits[0])
	}
	challenge.Alg = alg
	diff, err := strconv.Atoi(bits[1])
	if err != nil {
		return err
	}
	challenge.Difficulty = uint32(diff)
	challenge.Nonce, err = base64.RawStdEncoding.DecodeString(bits[2])
	if err != nil {
		return err
	}
	return nil
}

// Solve the challenge.
func (challenge *Challenge) solve(data []byte) Proof {
	switch challenge.Alg {
	case Sha111:
		return Proof{solveSha111(challenge.Nonce, challenge.Difficulty, data)}
	default:
		panic("No such algorithm")
	}
}
