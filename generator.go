package go_oauth2_server

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var _ IGenerator = &Generator{}

type IGenerator interface {
	GetSource() rand.Source
	Generate() string
}

type Generator struct {
	Len int
}

func (g *Generator) Generate() string {
	b := make([]byte, g.Len)
	for i, cache, remain := g.Len-1, g.GetSource().Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = g.GetSource().Int63(), letterIdxMax
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

func (g *Generator) GetSource() rand.Source {
	return rand.NewSource(time.Now().UnixNano())
}