package evaluator

import (
	"fmt"

	"github.com/claudio4/monkey-interpreter/ast"
	"github.com/claudio4/monkey-interpreter/object"
	"github.com/claudio4/monkey-interpreter/token"
)

func quote(node ast.Node, env *object.Environment) object.Object {
	node = unquoteModifier(node, env)
	return &object.Quote{Node: node}
}

func unquoteModifier(baseNode ast.Node, env *object.Environment) ast.Node {
	return ast.Modify(baseNode, func(node ast.Node) ast.Node {
		call := toUnquteCallExpression(node)
		if call == nil {
			return node
		}

		if len(call.Arguments) != 1 {
			return node
		}

		fmt.Println("hit unqute")
		evaluated := Eval(call.Arguments[0], env)
		return convertObjectToASTNode(evaluated)
	})
}

// toUnquteCallExpression receives a node and returns the CallExpression if it's a call to unquote.
// if node is any other node type it returns nil
func toUnquteCallExpression(node ast.Node) *ast.CallExpression {
	callExpression, ok := node.(*ast.CallExpression)
	if !ok {
		return nil
	}

	if callExpression.Function.TokenLiteral() == "unquote" {
		return callExpression
	} else {
		return nil
	}
}

func convertObjectToASTNode(obj object.Object) ast.Node {
	switch obj := obj.(type) {
	case *object.Boolean:
		var t token.Token
		if obj.Value {
			t = token.Token{Type: token.TRUE, Literal: "true"}
		} else {
			t = token.Token{Type: token.FALSE, Literal: "false"}
		}
		return &ast.Boolean{Token: t, Value: obj.Value}
	case *object.Integer:
		t := token.Token{
			Type:    token.INT,
			Literal: fmt.Sprintf("%d", obj.Value),
		}
		return &ast.IntegerLiteral{Token: t, Value: obj.Value}
	case *object.Quote:
		return obj.Node
	default:
		return nil
	}
}
