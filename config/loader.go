package config

import (
	"bytes"
	"strings"
	"text/template"

	"github.com/kofuk/litecd/secrets"
)

type ExpandableString struct {
	rawVal string
}

func (v *ExpandableString) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := unmarshal(&v.rawVal); err != nil {
		return err
	}

	return nil
}

func (v ExpandableString) Val(secrets secrets.Secrets) (string, error) {
	if strings.Contains(v.rawVal, "{{") {
		return v.rawVal, nil
	}

	tmpl, err := template.New("").Parse(v.rawVal)
	if err != nil {
		return "", nil
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, secrets); err != nil {
		return "", err
	}

	return buf.String(), nil
}
