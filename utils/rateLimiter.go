package utils

import (
	"fmt"
	"sync/atomic"
)

const EXSUASTED_TOKENS = 0

type RateLimiter struct {
	Tokens uint64
}

func (r *RateLimiter) Allow() bool {
	atomic.AddUint64(&r.Tokens, ^uint64(0))
	fmt.Println(r.Tokens)
	return r.Tokens != EXSUASTED_TOKENS
}

func (r *RateLimiter) ResetTokens(tokenSize uint64) {
	atomic.SwapUint64(&r.Tokens, tokenSize)
}
