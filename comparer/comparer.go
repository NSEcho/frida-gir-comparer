package comparer

import (
	"fmt"
	"github.com/nsecho/fgcomparer/parser"
	"sync"
)

type compareFn func(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup)

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
	src  *Source
}

func NewComparer(oldP, newP *parser.Parser) *Comparer {
	return &Comparer{
		oldP: oldP,
		newP: newP,
		src:  &Source{},
	}
}

func (c *Comparer) Compare() {
	var wg sync.WaitGroup
	wg.Add(len(functions))
	for name, fn := range functions {
		fn(c.src, name, c.oldP, c.newP, &wg)
	}
	wg.Wait()
}

func (c *Comparer) String() string {
	return c.src.String()
}

func logMessage(name, msg string) {
	if msg != "" {
		fmt.Printf("[*] [%s] %s\n", name, msg)
	}
}
