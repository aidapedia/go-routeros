package types

import "fmt"

type AiInt int

func (a AiInt) String() string {
	if a == -1 {
		return Unlimited
	}
	return fmt.Sprintf("%d", a)
}
