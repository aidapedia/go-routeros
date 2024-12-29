package types

import "fmt"

type AiInt int

func (a *AiInt) String() string {
	if a == nil {
		return Unlimited
	}
	return fmt.Sprintf("%v", *a)
}
