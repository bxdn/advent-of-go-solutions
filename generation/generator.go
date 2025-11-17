package generation

import (
	"advent-of-go/utils"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"os"
)

var fileTemplate = utils.Unpack(template.ParseFiles("generation/solution_template.txt"))
var registryTemplate = utils.Unpack(template.ParseFiles("generation/registry_template.txt"))

func Generate(year, day int) error {
	yearDirExists, e := doesYearDirExist(year)
	if e != nil {
		return fmt.Errorf("error checking if year directory exists: %w", e)
	}
	if e := createSolutions(year, day); e != nil {
		return fmt.Errorf("error creating solutions: %w", e)
	}
	if !yearDirExists {
		if e := addToMainSolutionSet(year); e != nil {
			return fmt.Errorf("error registering solution set: %w", e)
		}
		if e := addSolutionSet(year); e != nil {
			return fmt.Errorf("error creating solution set: %w", e)
		}
	}
	if e := addToSolutionSet(year, day); e != nil {
		return fmt.Errorf("error registering solutions: %w", e)
	}
	return nil
}

func addSolutionSet(year int) error {
	fileName := fmt.Sprintf("solutions/%d/solutions.go", year)
	file, e := os.Create(fileName)
	if e != nil {
		return fmt.Errorf("error creating solutions file: %w", e)
	}
	defer file.Close()
	if e := registryTemplate.Execute(file, year); e != nil {
		return fmt.Errorf("error executing template for registry: %w", e)
	}
	return nil
}

func addToMainSolutionSet(year int) error {
	fset := token.NewFileSet()
	fileName := "solutions/solutions.go"
	node, e := parser.ParseFile(fset, fileName, nil, 0)
	if e != nil {
		return fmt.Errorf("error parsing file: %w", e)
	}
	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			addYearToFunction(d, year)
		case *ast.GenDecl:
			addYearImport(d, year)
		}
	}
	file, e := os.Create(fileName)
	if e != nil {
		return fmt.Errorf("error creating file: %w", e)
	}
	defer file.Close()
	if e := printer.Fprint(file, fset, node); e != nil {
		return fmt.Errorf("error dumping new ast: %w", e)
	}
	return nil
}

func addToSolutionSet(year, day int) error {
	fset := token.NewFileSet()
	fileName := fmt.Sprintf("solutions/%d/solutions.go", year)
	node, e := parser.ParseFile(fset, fileName, nil, 0)
	if e != nil {
		return fmt.Errorf("error parsing file: %w", e)
	}
	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			addToFunction(d, day)
		case *ast.GenDecl:
			addImport(d, year, day)
		}
	}
	file, e := os.Create(fileName)
	if e != nil {
		return fmt.Errorf("error creating file: %w", e)
	}
	defer file.Close()
	if e := printer.Fprint(file, fset, node); e != nil {
		return fmt.Errorf("error dumping new ast: %w", e)
	}
	return nil
}

func addYearToFunction(d *ast.FuncDecl, year int) {
	callExpr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: fmt.Sprintf("y%d", year)},
			Sel: &ast.Ident{Name: "Solutions"},
		},
	}
	args := d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CallExpr).Args
	d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CallExpr).Args = append(args, callExpr)
}

func addYearImport(d *ast.GenDecl, year int) {
	d.Specs = append(d.Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%s"`, fmt.Sprintf("advent-of-go/solutions/%d", year)),
		},
		Name: &ast.Ident{Name: fmt.Sprintf("y%d", year)},
	})
}

func addImport(d *ast.GenDecl, year, day int) {
	d.Specs = append(d.Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: fmt.Sprintf(`"%s"`, fmt.Sprintf("advent-of-go/solutions/%d/day%d", year, day)),
		},
	})
}

func addToFunction(d *ast.FuncDecl, day int) {
	callExpr1 := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: fmt.Sprintf("day%d", day)},
			Sel: &ast.Ident{Name: "Pt1"},
		},
	}
	callExpr2 := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: fmt.Sprintf("day%d", day)},
			Sel: &ast.Ident{Name: "Pt2"},
		},
	}
	els := d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CompositeLit).Elts
	d.Body.List[0].(*ast.ReturnStmt).Results[0].(*ast.CompositeLit).Elts = append(els, callExpr1, callExpr2)
}

func doesYearDirExist(year int) (bool, error) {
	_, e := os.Stat(fmt.Sprintf("solutions/%d", year))
	if e != nil && !os.IsNotExist(e) {
		return false, fmt.Errorf("error checking if year directory exists: %w", e)
	}
	return !os.IsNotExist(e), nil
}

func createSolutions(year, day int) error {
	dirName := fmt.Sprintf("solutions/%d/day%d", year, day)
	if _, e := os.Stat(dirName); e == nil {
		return fmt.Errorf("solution already exists for year %d day %d", year, day)
	} else if !os.IsNotExist(e) {
		return fmt.Errorf("error checking if solution directory exists: %w", e)
	}
	if e := os.MkdirAll(dirName, 0777); e != nil {
		return fmt.Errorf("error creating solutions directory: %w", e)
	}
	if e := createSolutionFile(year, day, 1); e != nil {
		return fmt.Errorf("error creating part 1 solution file: %w", e)
	}
	if e := createSolutionFile(year, day, 2); e != nil {
		return fmt.Errorf("error creating part 2 solution file: %w", e)
	}
	return nil
}

func createSolutionFile(year, day, part int) error {
	pt1File, e := os.Create(fmt.Sprintf("solutions/%d/day%d/pt%d.go", year, day, part))
	if e != nil {
		return fmt.Errorf("error creating pt1 file: %w", e)
	}
	defer pt1File.Close()
	pt1Sol := utils.Solution{Year: year, Day: day, Part: part}
	if e := fileTemplate.Execute(pt1File, pt1Sol); e != nil {
		return fmt.Errorf("error executing template for part 1: %w", e)
	}
	return nil
}
