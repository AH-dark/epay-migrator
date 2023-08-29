package conn

import (
	"gorm.io/gorm/schema"
	"strings"
)

type replacer struct {
}

func (r *replacer) Replace(name string) string {
	if strings.Contains(name, "reg_code") {
		return strings.ReplaceAll(name, "reg_code", "regcode")
	}

	return name
}

var _ schema.Replacer = (*replacer)(nil)

func NewReplacer() schema.Replacer {
	return &replacer{}
}
