package runes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*

/// DESCRIPTION

Julius Caesar protected his confidential information by encrypting it using a cipher.
Caesar's cipher shifts each letter by a number of letters. For example, in the case of
a rotation by 3, w, x, y and z would map to z, a, b and c.

Original alphabet:      abcdefghijklmnopqrstuvwxyz
Alphabet rotated +3:    defghijklmnopqrstuvwxyzabc

/// REQUIREMENTS.

1. If the shift takes you past the end of the alphabet, just rotate back to the front of the alphabet.
2. If the shift number is a negative number, log an error. X
3. The cipher only encrypts letters; all other characters remain unencrypted.
4. Build a function CaesarCipher(s string, n int) (string, error) to handle your implementation.
5. Create unit tests to validate previous requirements.

/// EXAMPLE.

s := "uno-square *es* %una gran empresa."
n := 3

The alphabet is rotated by 3, matching the mapping above.

The encrypted string is: "xqr-vtxduh *hv* %xqd judq hpsuhvd."
><
*/

func addToRune(r rune, addVal int) (rune, error) {
	if addVal < 0 {
		return r, fmt.Errorf("invalid addVal")
	}
	if r >= LC_A && r <= LC_Z {
		return (r-LC_A+rune(addVal))%LC_MOD + LC_A, nil
	} else if r >= UC_A && r <= UC_Z {
		return (r-UC_A+rune(addVal))%UC_MOD + UC_A, nil
	}
	return r, nil
}

func CaesarCipher(s string, n int) (string, error) {
	var res string
	for _, r := range s {
		rot, err := addToRune(r, n)
		if err != nil {
			return res, err
		}
		res += fmt.Sprintf("%c", rot)
	}
	return res, nil
}

func Test_Runes_Excersise(t *testing.T) {
	c, _ := CaesarCipher("uno-square *es* %una gran empresa.", 3)
	fmt.Println(c)

	assert.Equal(t, "xqr-vtxduh *hv* %xqd judq hpsuhvd.", c)
}
