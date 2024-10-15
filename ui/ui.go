package ui

import (
	"fmt"

	"github.com/massarakhsh/lik/likdom"
	"github.com/massarakhsh/lik/stack"
)

func GenPage(do *stack.ItStack) likdom.Domer {
	html := likdom.BuildPageHtml()
	if head, _ := html.GetDataTag("head"); head != nil {
		head.BuildString("<script type='text/javascript' src='/js/jquery.js'></script>")
		head.BuildString("<link rel='stylesheet' href='/js/styles.css'/>")
		head.BuildString("<script type='text/javascript' src='/js/scripts.js'></script>")
	}
	if body, _ := html.GetDataTag("body"); body != nil {
		script := fmt.Sprintf("start_run(%d);", 0)
		body.SetAttr("onload", script)
		body.AppendItem(showPage(do))
	}
	return html
}

func ShowTextLink(text string, path string) likdom.Domer {
	data := likdom.BuildItem("a", "href", path)
	if text != "" {
		data.BuildString(text)
	} else {
		data.BuildString(path)
	}
	return data
}
