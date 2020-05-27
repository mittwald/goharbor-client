package harbor

import (
	"errors"
	"regexp"
)

var hasLower, hasUpper, hasNumber *regexp.Regexp
var ErrPasswordMalformed = errors.New(
	"password does not meet the requirements (at least 8 chars, 1 uppercase letter, 1 lowercase letter and 1 number)",
)

func init() {
	hasLower = regexp.MustCompile(`[a-z]`)
	hasUpper = regexp.MustCompile(`[A-Z]`)
	hasNumber = regexp.MustCompile(`[0-9]`)
}

// ValidatePassword takes the pw string and validates it to have at least
// 8 chars, 1 uppercase letter, 1 lowercase letter and 1 number.
// see https://github.com/goharbor/harbor/blob/fdbd59daf4502de6013771ee3df12877a5e2dace/src/core/api/user.go#L613
func ValidatePassword(pw string) error {
	if len(pw) >= 8 && hasLower.MatchString(pw) && hasUpper.MatchString(pw) && hasNumber.MatchString(pw) {
		return nil
	}
	return ErrPasswordMalformed
}
