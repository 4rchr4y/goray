package ason

import (
	"go/ast"
	"go/token"
	"os"
)

// TODO List
//
// *ast.Field
// *ast.FieldList
// *ast.BadExpr
// *ast.Ellipsis
// *ast.FuncLit
// *ast.CompositeLit
// *ast.ParenExpr
// *ast.SelectorExpr
// *ast.IndexExpr
// *ast.IndexListExpr
// *ast.SliceExpr
// *ast.TypeAssertExpr
// *ast.CallExpr:
// *ast.StarExpr:
// *ast.UnaryExpr:
// *ast.BinaryExpr
// *ast.KeyValueExpr
// *ast.ArrayType
// *ast.StructType
// *ast.FuncType
// *ast.InterfaceType
// *ast.MapType
// *ast.ChanType
// *ast.BadStmt
// *ast.DeclStmt
// *ast.EmptyStmt
// *ast.LabeledStmt
// *ast.ExprStmt
// *ast.SendStmt
// *ast.IncDecStmt
// *ast.AssignStmt
// *ast.GoStmt
// *ast.DeferStmt
// *ast.ReturnStmt
// *ast.BranchStmt
// *ast.BlockStmt
// *ast.IfStmt
// *ast.CaseClause
// *ast.SwitchStmt
// *ast.TypeSwitchStmt
// *ast.CommClause
// *ast.SelectStmt
// *ast.ForStmt
// *ast.RangeStmt
// *ast.ImportSpec
// *ast.TypeSpec
// *ast.BadDecl
// *ast.GenDecl
// *ast.FuncDecl
// *ast.File
// *ast.Package

type SerConfig struct {
	// RefCounterEnable flag must be used when you carry out some manual
	// manipulations with the source AST tree. For example, you duplicate nodes,
	// which can create nodes that have the same references to the original object in memory.
	//
	// Use this flag when duplicating nodes containing many fields.
	RefCounterEnable bool

	// Standard Position structure contains fields such as `Filename`, `Offset`, `Line` and `Column`.
	// Usually, all this information is not required for, but only the
	// `Line` in the source code file and `Filename` is required.
	//
	// Use this flag when you do not need support for backward compatibility with the original AST
	// and you do not require the fields that the standard position structure contains
	PosCompress bool
}

type SerPass struct {
	fset     *token.FileSet
	refmap   map[ast.Node]Ason
	refcount uint
	conf     *SerConfig
}

// SerPassOptFn is a functional option type that allows us to configure the SerPass.
type SerPassOptFn func(*SerPass)

func NewSerPass(fset *token.FileSet, options ...SerPassOptFn) *SerPass {
	pass := &SerPass{
		fset: fset,
		conf: new(SerConfig),
	}

	for _, opt := range options {
		opt(pass)
	}

	return pass
}

func WithRefCounter() SerPassOptFn {
	return func(pass *SerPass) {
		pass.refmap = make(map[ast.Node]Ason)
		pass.conf.RefCounterEnable = true
	}
}

func WithPosCompression() SerPassOptFn {
	return func(pass *SerPass) {
		pass.conf.PosCompress = true
	}
}

type SerFn[I ast.Node, R Ason] func(*SerPass, I) R

func WithRefLookup[I ast.Node, R Ason](pass *SerPass, input I, ser SerFn[I, R]) R {
	if node, exists := pass.refmap[input]; exists {
		return node.(R)
	}

	node := ser(pass, input)

	pass.refcount++
	pass.refmap[input] = node

	return node
}

func SerProcessList[I ast.Node, R Ason](pass *SerPass, inputList []I, ser SerFn[I, R]) []R {
	if inputList == nil {
		return nil
	}

	result := make([]R, len(inputList))
	for i := 0; i < len(inputList); i++ {
		result[i] = ser(pass, inputList[i])
	}

	return result
}

func SerProcessNode(node ast.Node) Node {
	return Node{
		Type: SerProcessNodeType(node),
	}
}

func SerProcessNodeWithRef(node ast.Node, ref uint) Node {
	return Node{
		Ref:  ref,
		Type: SerProcessNodeType(node),
	}
}

func SerializeFile(pass *SerPass, input *ast.File) *File {
	return &File{
		Name:      SerializeIdent(pass, input.Name),
		Decls:     SerProcessList[ast.Decl, Decl](pass, input.Decls, SerProcessDecl),
		Doc:       SerProcessCommentGroup(pass, input.Doc),
		Package:   SerProcessPos(pass, input.Package),
		Loc:       SerProcessLoc(pass, input.FileStart, input.FileEnd),
		Size:      SerProcessFileSize(pass, input),
		Comments:  SerProcessList[*ast.CommentGroup, *CommentGroup](pass, input.Comments, SerProcessCommentGroup),
		GoVersion: input.GoVersion,
		Node:      SerProcessNode(input),
	}
}

func SerProcessFileSize(pass *SerPass, input *ast.File) int {
	position := pass.fset.PositionFor(input.Name.NamePos, false)
	content, err := os.ReadFile(position.Filename)
	if err != nil {
		return 1<<_GOARCH() - 2
	}

	return len(content)
}

func SerializePos(pass *SerPass, pos token.Pos) Pos {
	position := pass.fset.PositionFor(pos, false)

	if pass.conf.PosCompress {
		return &PosCompressed{
			Filename: position.Filename,
			Line:     position.Line,
		}
	}

	return &Position{
		Filename: position.Filename,
		Offset:   position.Offset,
		Line:     position.Line,
		Column:   position.Column,
	}
}

func SerProcessPos(pass *SerPass, pos token.Pos) Pos {
	if pos != token.NoPos {
		return SerializePos(pass, pos)
	}

	return new(NoPos)
}

func SerProcessLoc(pass *SerPass, start, end token.Pos) *Loc {
	loc := new(Loc)

	if start != token.NoPos {
		loc.Start = SerializePos(pass, start)
	}

	if end != token.NoPos {
		loc.End = SerializePos(pass, end)
	}

	return loc
}

func SerializeComment(pass *SerPass, input *ast.Comment) *Comment {
	return &Comment{
		Node:  SerProcessNode(input),
		Slash: SerProcessPos(pass, input.Slash),
		Text:  input.Text,
	}
}

func SerializeCommentGroup(pass *SerPass, input *ast.CommentGroup) *CommentGroup {
	return &CommentGroup{
		Node: SerProcessNode(input),
		List: SerProcessList[*ast.Comment, *Comment](pass, input.List, SerializeComment),
	}
}

func SerProcessCommentGroup(pass *SerPass, group *ast.CommentGroup) *CommentGroup {
	if group != nil {
		return SerializeCommentGroup(pass, group)
	}

	return nil
}

func SerializeIdent(pass *SerPass, input *ast.Ident) *Ident {
	if !pass.conf.RefCounterEnable {
		return serializeIdent(pass, input)
	}

	return WithRefLookup[*ast.Ident, *Ident](pass, input, serializeIdent)
}

func serializeIdent(pass *SerPass, input *ast.Ident) *Ident {
	return &Ident{
		Loc:     SerProcessLoc(pass, input.Pos(), input.End()),
		NamePos: SerProcessPos(pass, input.NamePos),
		Name:    input.Name,
		Node:    SerProcessNodeWithRef(input, pass.refcount),
	}

}

func SerializeBasicLit(pass *SerPass, input *ast.BasicLit) *BasicLit {
	if !pass.conf.RefCounterEnable {
		return serializeBasicLit(pass, input)
	}

	return WithRefLookup[*ast.BasicLit, *BasicLit](pass, input, serializeBasicLit)
}

func serializeBasicLit(pass *SerPass, input *ast.BasicLit) *BasicLit {
	return &BasicLit{
		Loc:      SerProcessLoc(pass, input.Pos(), input.End()),
		ValuePos: SerProcessPos(pass, input.ValuePos),
		Kind:     input.Kind.String(),
		Value:    input.Value,
		Node:     SerProcessNodeWithRef(input, pass.refcount),
	}
}

func SerializeValueSpec(pass *SerPass, input *ast.ValueSpec) *ValueSpec {
	if !pass.conf.RefCounterEnable {
		return serializeValueSpec(pass, input)
	}

	return WithRefLookup[*ast.ValueSpec, *ValueSpec](pass, input, serializeValueSpec)
}

func serializeValueSpec(pass *SerPass, input *ast.ValueSpec) *ValueSpec {
	return &ValueSpec{
		Loc:    SerProcessLoc(pass, input.Pos(), input.End()),
		Names:  SerProcessList[*ast.Ident, *Ident](pass, input.Names, SerializeIdent),
		Values: SerProcessList[ast.Expr, Expr](pass, input.Values, SerProcessExpr),
		Node:   SerProcessNode(input),
	}
}

func SerializeGenDecl(pass *SerPass, input *ast.GenDecl) *GenDecl {
	if !pass.conf.RefCounterEnable {
		return serializeGenDecl(pass, input)
	}

	return WithRefLookup[*ast.GenDecl, *GenDecl](pass, input, serializeGenDecl)
}

func serializeGenDecl(pass *SerPass, input *ast.GenDecl) *GenDecl {
	return &GenDecl{
		Loc:      SerProcessLoc(pass, input.Pos(), input.End()),
		TokenPos: SerProcessPos(pass, input.TokPos),
		Lparen:   SerProcessPos(pass, input.Lparen),
		Rparen:   SerProcessPos(pass, input.Rparen),
		Tok:      input.Tok.String(),
		Specs:    SerProcessList[ast.Spec, Spec](pass, input.Specs, SerProcessSpec),
		Node:     SerProcessNode(input),
	}
}

func SerProcessExpr(pass *SerPass, expr ast.Expr) Expr {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return SerializeBasicLit(pass, e)
	default:
		return nil
	}
}

func SerProcessSpec(pass *SerPass, spec ast.Spec) Spec {
	switch s := spec.(type) {
	case *ast.ValueSpec:
		return SerializeValueSpec(pass, s)
	default:
		return nil
	}
}

func SerProcessDecl(pass *SerPass, decl ast.Decl) Decl {
	switch d := decl.(type) {
	case *ast.GenDecl:
		return SerializeGenDecl(pass, d)
	default:
		return nil
	}
}

// SerProcessNodeType takes any type that can represents AST node and returns its type as a string.
// It matches the node type with predefined NodeType constants.
// If the node type is not recognized, NodeTypeInvalid is returned.
func SerProcessNodeType(node ast.Node) string {
	// n :=
	switch node.(type) {
	case *ast.File:
		return NodeTypeFile
	case *ast.Comment:
		return NodeTypeComment
	case *ast.CommentGroup:
		return NodeTypeCommentGroup
	case *ast.Ident:
		return NodeTypeIdent
	case *ast.BasicLit:
		return NodeTypeBasicLit
	case *ast.ValueSpec:
		return NodeTypeValueSpec
	case *ast.GenDecl:
		return NodeTypeGenDecl
	default:
		return NodeTypeInvalid
	}
}