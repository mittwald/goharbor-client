package harbor

import "testing"

func TestValidatePassword(t *testing.T) {
	validPasswords := []string{
		"Test1234",
		"F§f3fedwf",
		"D§if3m2mßf3m,f0g4mfg9",
	}

	invalidPasswords := []string{
		"",
		"t",
		"Aa1",
		"aA1",
		"1aA",
		"1Aa",
		"12345678",
		"aaaaaaaa",
		"bbbbbbbbbbbb",
		"test1234",
	}

	t.Run("valid passwords", func(t *testing.T) {
		for _, pw := range validPasswords {
			err := ValidatePassword(pw)
			if err != nil {
				t.Errorf("%s", err.Error())
			}
		}
	})

	t.Run("invalid passwords", func(t *testing.T) {
		for _, pw := range invalidPasswords {
			err := ValidatePassword(pw)
			if err != ErrPasswordMalformed {
				t.Errorf("ValidatePassword does not return with ErrPasswordMalformed")
			}
		}
	})
}
