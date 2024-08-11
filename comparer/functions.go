package comparer

import (
	"fmt"
	"github.com/nsecho/fgcomparer/parser"
	"sync"
)

func classCount(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Classes())
	nlen := len(newP.Classes())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedClasses(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, newCls := range newP.Classes() {
		found = false
		for _, oldCls := range oldP.Classes() {
			if newCls.Name == oldCls.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", newCls.Name)
		}
	}
	logMessage(name, msg)
}

func deletedClasses(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, oldCls := range oldP.Classes() {
		found = false
		for _, newCls := range newP.Classes() {
			if newCls.Name == oldCls.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", oldCls.Name)
		}
	}
	logMessage(name, msg)
}

func functionCount(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Functions())
	nlen := len(newP.Functions())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedFunctions(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, newFn := range newP.Functions() {
		found = false
		for _, oldFn := range oldP.Functions() {
			if newFn.Name == oldFn.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", newFn.Name)
		}
	}
	logMessage(name, msg)
}

func deletedFunctions(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, oldFn := range oldP.Functions() {
		found = false
		for _, newFn := range newP.Functions() {
			if newFn.Name == oldFn.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", oldFn.Name)
		}
	}
	logMessage(name, msg)
}

func enumerationCount(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Enumerations())
	nlen := len(newP.Enumerations())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedEnumerations(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, newEn := range newP.Enumerations() {
		found = false
		for _, oldEn := range oldP.Enumerations() {
			if newEn.Name == oldEn.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", newEn.Name)
		}
	}
	logMessage(name, msg)
}

func deletedEnumerations(name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := ""
	var found bool
	for _, oldEn := range oldP.Enumerations() {
		found = false
		for _, newEn := range newP.Enumerations() {
			if newEn.Name == oldEn.Name {
				found = true
				break
			}
		}
		if !found {
			msg += fmt.Sprintf("\n\t%s", oldEn.Name)
		}
	}
	logMessage(name, msg)
}
