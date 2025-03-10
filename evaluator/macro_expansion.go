package evaluator

import (
	"github.com/claudio4/monkey-interpreter/ast"
	"github.com/claudio4/monkey-interpreter/object"
)

func ExpandMacros(program ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(program, func(node ast.Node) ast.Node {
		callExpression, ok := node.(*ast.CallExpression)
		if !ok {
			return node
		}

		macro := isMacroCall(callExpression, env)
		if macro == nil {
			return node
		}

		args := quoteArgs(callExpression)
		evalEnv := extendMacroEnv(macro, args)

		evaluated := Eval(macro.Body, evalEnv)

		quote, ok := evaluated.(*object.Quote)
		if !ok {
			panic("we only support returning AST-nodes from macros")
		}

		return quote.Node
	})
}

func DefineMacros(program *ast.Program, env *object.Environment) {
	definitions := []int{}

	for i, stmt := range program.Statements {
		if tryAddMacro(stmt, env) {
			definitions = append(definitions, i)
		}
	}

	for i := len(definitions) - 1; i >= 0; i-- {
		defIdx := definitions[i]
		program.Statements = append(program.Statements[:defIdx], program.Statements[defIdx+1:]...)
	}
}

func isMacroCall(exp *ast.CallExpression, env *object.Environment) *object.Macro {
	ident, ok := exp.Function.(*ast.Identifier)
	if !ok {
		return nil
	}

	obj, ok := env.Get(ident.Value)
	if !ok {
		return nil
	}

	macro, ok := obj.(*object.Macro)
	if !ok {
		return nil
	}

	return macro
}

func quoteArgs(exp *ast.CallExpression) []*object.Quote {
	args := []*object.Quote{}

	for _, a := range exp.Arguments {
		args = append(args, &object.Quote{Node: a})
	}

	return args
}

func extendMacroEnv(macro *object.Macro, args []*object.Quote) *object.Environment {
	extended := object.NewEnclosedEnvironment(macro.Env)
	for i, param := range macro.Parameters {
		extended.Set(param.Value, args[i])
	}

	return extended
}

func tryAddMacro(stmt ast.Statement, env *object.Environment) bool {
	letStatement, ok := stmt.(*ast.LetStatement)
	if !ok {
		return false
	}
	macroLiteral, ok := letStatement.Value.(*ast.MacroLiteral)
	if !ok {
		return false
	}

	macro := &object.Macro{
		Parameters: macroLiteral.Parameters,
		Env:        env,
		Body:       macroLiteral.Body,
	}

	env.Set(letStatement.Name.Value, macro)
	return true
}
