package types

import "strings"

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewMaintainer(value string) Maintainer {

	var maintainer Maintainer

	return maintainer

}

func ToMaintainer(value string) Maintainer {

	var maintainer Maintainer

	if strings.Contains(value, " <") && strings.Contains(value, "@") && strings.HasSuffix(value, ">") {

		tmp := strings.Split(value[:len(value)-1], " <")
		name := strings.TrimSpace(tmp[0])
		email := strings.TrimSpace(tmp[1])

		if name != "" {
			maintainer.Name = name
		}

		if email != "" {
			maintainer.Email = email
		}

	} else if strings.HasPrefix(value, "<") && strings.Contains(value, "@") && strings.HasSuffix(value, ">") {

		email := strings.TrimSpace(value[1 : len(value)-2])

		if email != "" {
			maintainer.Email = email
		}

	} else if strings.Contains(value, "@") && !strings.Contains(value, " ") {

		email := strings.TrimSpace(value)

		if email != "" {
			maintainer.Email = email
		}

	} else {

		name := strings.TrimSpace(value)

		if name != "" {
			maintainer.Name = name
		}

	}

	return maintainer

}

func (maintainer *Maintainer) IsIdentical(value Maintainer) bool {

	var result bool = false

	if maintainer.Name == value.Name {
		result = true
	}

	return result

}

func (maintainer *Maintainer) IsValid() bool {

	if maintainer.Name != "" {

		if maintainer.Email != "" && strings.Contains(maintainer.Email, "@") {
			return true
		}

	}

	return false

}
