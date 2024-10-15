package ui

import (
	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/lik/stack"
)

func showPage(do *stack.ItStack) likdom.Domer {
	data := likdom.BuildDiv()
	data.BuildString("Hello!")
	return data
}
