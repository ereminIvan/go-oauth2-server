package go_oauth2_server

import (
	"math/rand"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var _ IGenerator = &generator{}

//IGenerator
type IGenerator interface {
	Generate() string
}

//Generator
type generator struct {
	len int
}

//Generate random string with Len
func (g *generator) Generate() string {
	b := make([]byte, g.len)
	for i, cache, remain := g.len-1, g.getSource().Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = g.getSource().Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

//GetSource
func (g *generator) getSource() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}

//NewGenerator create new generator
func NewGenerator(len int) string {
	return &generator{len: len}
}
