package comparer

import (
	"fmt"
	"github.com/nsecho/fgcomparer/parser"
	"sync"
)

func classCount(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Classes())
	nlen := len(newP.Classes())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedClasses(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
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
			src.NewClasses = append(src.NewClasses, newCls)
		}
	}
}

func deletedClasses(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
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
			src.RemovedClasses = append(src.RemovedClasses, oldCls.Name)
		}
	}
}

func functionCount(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Functions())
	nlen := len(newP.Functions())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedFunctions(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
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

func deletedFunctions(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
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

func enumerationCount(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
	olen := len(oldP.Enumerations())
	nlen := len(newP.Enumerations())
	if olen != nlen {
		logMessage(name, fmt.Sprintf("old count: %d; new count: %d", olen, nlen))
	}
}

func addedEnumerations(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
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
			src.NewEnumerations = append(src.NewEnumerations, newEn)
		}
	}
}

func deletedEnumerations(src *Source, name string, oldP, newP *parser.Parser, wg *sync.WaitGroup) {
	defer wg.Done()
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
			src.RemovedEnumerations = append(src.RemovedEnumerations, oldEn.Name)
		}
	}
}
