package runes

import "fmt"

const (
	// rune is alias of int32
	// 65 is the ASCII code of A
	// 90 is the ASCII code of Z
	// 97 is the ASCII code of a
	// 122 is the ASCII code of z
	LC_A   = 'a'
	LC_Z   = 'z'
	UC_A   = 'A'
	UC_Z   = 'Z'
	LC_MOD = LC_Z - LC_A
	UC_MOD = UC_Z - UC_A
)

type RuneOperator interface {
	AddToRune(r rune, addVal int) (rune, error)
	AddToRuneGuaranteed(r rune, addVal int) rune
	FmtRune(r rune)
}

type runeOperatorImpl struct {
}

func NewRunOperator() RuneOperator {
	return &runeOperatorImpl{}
}

func isValidRange(r rune) bool {
	return (r >= LC_A && r <= LC_Z) || (r >= UC_A && r <= UC_Z)
}

func (roi *runeOperatorImpl) AddToRune(r rune, addVal int) (rune, error) {
	res := rune(int(r) + addVal)
	if !isValidRange(res) {
		return res, fmt.Errorf("invalid range")
	}
	return res, nil
}

func (roi *runeOperatorImpl) AddToRuneGuaranteed(r rune, addVal int) (res rune) {
	if r >= LC_A && r <= LC_Z {
		// e.g. being 'c' = 99 + 3
		//      (99-97+3)=5 => (5%26)=5 + 97 = 'f'
		b := (r - LC_A + rune(addVal))
		m := b % LC_MOD
		res = m + LC_A
	} else if r >= UC_A && r <= UC_Z {
		b := (r - UC_A + rune(addVal))
		m := b % UC_MOD
		res = m + UC_A
	}
	// otherwise original rune is returned without any sum
	return res
}

func (roi *runeOperatorImpl) FmtRune(r rune) {
	fmt.Printf("%c %d %T\n", r, r, r)
}
