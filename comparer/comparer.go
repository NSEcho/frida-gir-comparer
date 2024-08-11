package comparer

import (
	"fmt"
	"github.com/nsecho/fgcomparer/parser"
	"sync"
)

type compareFn func(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup)

var functions = map[string]compareFn{
	"ClassCount":          classCount,
	"AddedClasses":        addedClasses,
	"DeletedClasses":      deletedClasses,
	"FunctionCount":       functionCount,
	"AddedFunctions":      addedFunctions,
	"DeletedFunctions":    deletedFunctions,
	"EnumerationCount":    enumerationCount,
	"AddedEnumerations":   addedEnumerations,
	"DeletedEnumerations": deletedEnumerations,
}

type Comparer struct {
	oldP *parser.Parser
	newP *parser.Parser
}

func NewComparer(oldP, newP *parser.Parser) *Comparer {
	return &Comparer{
		oldP: oldP,
		newP: newP,
	}
}

func (c *Comparer) Compare() {
	var wg sync.WaitGroup
	wg.Add(len(functions))
	for name, fn := range functions {
		fn(name, c.oldP, c.newP, &wg)
	}
	wg.Wait()
}

func logMessage(name, msg string) {
	if msg != "" {
		fmt.Printf("[*] [%s] %s\n", name, msg)
	}
}
