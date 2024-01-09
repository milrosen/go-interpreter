package ast

import (
	"go_interpreter/token"
	"testing"
)

func TestAstString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DefStatement{
				Token: token.Token{Type: token.DEF, Literal: "def"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "fireworks"},
					Value: "fireworks",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "skyRocketsInFlight"},
					Value: "skyRocketsInFlight",
				},
			},
		},
	}

	if program.String() != "def fireworks := skyRocketsInFlight" {
		t.Errorf("program.String() returned %s, not def fireworks := skyRocketsInFlight", program.String())
	}
}
