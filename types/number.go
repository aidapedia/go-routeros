package types

import "fmt"

// AiInt is a custom type for int
type AiInt int

func (a *AiInt) String() string {
	if a == nil {
		return Unlimited
	}
	return fmt.Sprintf("%v", *a)
}

// AiInt64 is a custom type for int64
type AiInt64 int64

func (a *AiInt64) String() string {
	if a == nil {
		return Unlimited
	}
	return fmt.Sprintf("%v", *a)
}
