package evaluator

import (
	"fmt"
	"github.com/joeygibson/monkey/object"
	"time"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"now": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}

			now := time.Now()
			return &object.String{Value: now.String()}
		},
	},
	"nowInMillis": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("wrong number of arguments. got=%d, want=0", len(args))
			}

			millis := time.Now().UnixNano() / int64(time.Millisecond)
			return &object.Integer{Value: millis}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}

				return arg.Elements[0]
			default:
				return newError("argument to `first` not supported, got %s", args[0].Type())
			}
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) == 0 {
					return NULL
				}

				lastIndex := len(arg.Elements) - 1

				return arg.Elements[lastIndex]
			default:
				return newError("argument to `last` not supported, got %s", args[0].Type())
			}
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) < 2 {
					return NULL
				}

				length := len(arg.Elements)
				elems := make([]object.Object, length-1, length-1)
				copy(elems, arg.Elements[1:length])

				return &object.Array{Elements: elems}

			default:
				return newError("argument to `last` not supported, got %s", args[0].Type())
			}
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				if len(arg.Elements) < 2 {
					return NULL
				}

				length := len(arg.Elements)
				newElems := make([]object.Object, length+1, length+1)

				copy(newElems, arg.Elements)
				newElems[length] = args[1]

				return &object.Array{Elements: newElems}

			default:
				return newError("argument to `last` not supported, got %s", args[0].Type())
			}
		},
	},
	"puts": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return newError("wrong number of arguments. got=%d, want=> 0", len(args))
			}

			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
