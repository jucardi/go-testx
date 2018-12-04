package testx

import (
	"fmt"
	"gopkg.in/jucardi/go-terminal-colors.v1"
	"testing"
)

func Convey(description string, t *testing.T, f func()) {
	ctx := newCtx(t)
	contextPile = append(contextPile, ctx)
	ctx.Println("\n").PrintIndent(description+" ", descriptionColor()...)
	f()
	if len(contextPile) <= 2 {
		ctx.Println(fmt.Sprintf("\n\n%d total assertions", ctx.assertions), assertionsColor()...)
	}
	contextPile = contextPile[:len(contextPile)-1]
	currentCtx().assertions += ctx.assertions
}

func descriptionColor() []fmtc.Color {
	if len(contextPile) > 2 {
		return []fmtc.Color{fmtc.White}
	}
	return []fmtc.Color{fmtc.Yellow, fmtc.Bold}
}

func assertionsColor() []fmtc.Color {
	if len(contextPile) > 2 {
		return []fmtc.Color{fmtc.Green}
	}
	return []fmtc.Color{fmtc.Green}
}
