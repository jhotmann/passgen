package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

const defaultSpecials string = "!@#$%^&*"

func passgenV2(opts opts) string {
	charset := createCharset(opts)
	var out string

	for i := 0; true; i++ {
		if len(out) >= opts.length { // password is long enough
			if containsNecessaryCharacters(out[0:opts.length], opts) { // password contains all necessary character classes
				break
			} else { // password is missing character classes, keep second half and keep generating if necessary
				out = out[int(len(out)/2):]
			}
		}

		passphrasePart := strings.Repeat(opts.passphrase, i+1)
		hash := sha256.Sum256([]byte(passphrasePart + opts.salt))

		// convert hash array of bytes into bigint
		num := new(big.Int)
		num.SetString(hex.EncodeToString(hash[:]), 16)
		// encode bigint in custom character set
		encoded := encodeToCustomCharset(num, charset)

		out = fmt.Sprintf("%s%s", out, encoded)
	}

	return out[0:opts.length]
}

func createCharset(opts opts) string {
	charset := ""
	if !opts.noUppers {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}
	charset += "abcdefghijklmnopqrstuvwxyz"
	if !opts.noNumbers {
		charset += "0123456789"
	}
	if !opts.noSpecials {
		if opts.customSpecials != "" {
			charset += opts.customSpecials
		} else {
			charset += defaultSpecials
		}
	}

	return charset
}

func containsNecessaryCharacters(s string, opts opts) bool {
	if !opts.noNumbers {
		matches, _ := regexp.MatchString(".*[0-9].*", s)
		if !matches {
			return false
		}
	}

	if !opts.noUppers {
		matches, _ := regexp.MatchString(".*[A-Z].*", s)
		if !matches {
			return false
		}
	}

	if !opts.noSpecials {
		customSpecials := defaultSpecials
		if opts.customSpecials != "" {
			customSpecials = opts.customSpecials
		}

		escapeChars := []string{"-", "[", "]", "^"}
		for _, e := range escapeChars {
			customSpecials = strings.ReplaceAll(customSpecials, e, "\\"+e)
		}

		matches, _ := regexp.MatchString(".*["+customSpecials+"].*", s)
		if !matches {
			return false
		}
	}

	return true
}

func encodeToCustomCharset(v *big.Int, charset string) string {
	charsetLength := big.NewInt(int64(len(charset)))
	var ret string
	for {
		if v.Cmp(big.NewInt(0)) == 0 {
			break
		}
		remainder := new(big.Int).Mod(v, charsetLength)
		v = new(big.Int).Div(v, charsetLength)

		ret = fmt.Sprintf("%c%s", charset[remainder.Uint64()], ret)
	}
	return ret
}
