// Code generated by 'goexports go/doc'. DO NOT EDIT.

// +build go1.12,!go1.13

package stdlib

import (
	"go/doc"
	"reflect"
)

func init() {
	Symbols["go/doc"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"AllDecls":        reflect.ValueOf(doc.AllDecls),
		"AllMethods":      reflect.ValueOf(doc.AllMethods),
		"Examples":        reflect.ValueOf(doc.Examples),
		"IllegalPrefixes": reflect.ValueOf(&doc.IllegalPrefixes).Elem(),
		"IsPredeclared":   reflect.ValueOf(doc.IsPredeclared),
		"New":             reflect.ValueOf(doc.New),
		"PreserveAST":     reflect.ValueOf(doc.PreserveAST),
		"Synopsis":        reflect.ValueOf(doc.Synopsis),
		"ToHTML":          reflect.ValueOf(doc.ToHTML),
		"ToText":          reflect.ValueOf(doc.ToText),

		// type definitions
		"Example": reflect.ValueOf((*doc.Example)(nil)),
		"Filter":  reflect.ValueOf((*doc.Filter)(nil)),
		"Func":    reflect.ValueOf((*doc.Func)(nil)),
		"Mode":    reflect.ValueOf((*doc.Mode)(nil)),
		"Note":    reflect.ValueOf((*doc.Note)(nil)),
		"Package": reflect.ValueOf((*doc.Package)(nil)),
		"Type":    reflect.ValueOf((*doc.Type)(nil)),
		"Value":   reflect.ValueOf((*doc.Value)(nil)),
	}
}
