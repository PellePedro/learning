
[Text Templaytes](https://github.com/RangelReale/gotexttemplate/tree/master/util/goconv/cmd/goconv)

[mobile](https://github.com/guyezi/mobile/blob/3c8601c510d0503ac84d1e5cb8e24de550201dea/cmd/gobind/main.go#L69)
[MObile Internal] (https://github.com/guyezi/mobile/blob/3c8601c510d0503ac84d1e5cb8e24de550201dea/internal/importers/ast.go)

```
	testDeps := flag.Bool("t", false, "Include test dependencies")
	std := flag.Bool("std", false, "Include standard library dependencies")
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedImports | packages.NeedDeps,
		Tests: *testDeps,
	}
	pkgs, err := packages.Load(cfg, pkg)
	if err != nil {
		log.Fatalln("Error loading packages:", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	depSet := make(map[string]struct{})
	pre := func(pkg *packages.Package) bool {
		depSet[pkg.PkgPath] = struct{}{}
		return true
	}
	packages.Visit(pkgs, pre, nil)
	if !*std {
		cfg.Tests = false
		stdPkgs, err := packages.Load(cfg, "std")
		if err != nil {
			log.Fatalln("Error discovering packages:", err)
		}
		if packages.PrintErrors(pkgs) > 0 {
			os.Exit(1)
		}

		stdSet := make(map[string]struct{})
		pre := func(pkg *packages.Package) bool {
			stdSet[pkg.PkgPath] = struct{}{}
			return true
		}
		packages.Visit(stdPkgs, pre, nil)
		for pkg := range stdSet {
			delete(depSet, pkg)
		}
	}
	for _, pkg := range pkgs {
		delete(depSet, pkg.PkgPath)
	}
	var deps []string
	for dep := range depSet {
		deps = append(deps, dep)
	}
	sort.Strings(deps)
	for _, dep := range deps {
		fmt.Println(dep)
	}
  
  
```




```
func ExtractSymbols(f *ast.File) []string {
	symbols := []string{}
	for _, decl := range f.Decls {
		names := []string{}
		switch v := decl.(type) {
		case *ast.FuncDecl:
			names = append(names, v.Name.Name)
		case *ast.GenDecl:
			switch v.Tok {
			case token.TYPE:
				s := v.Specs[0].(*ast.TypeSpec)
				names = append(names, s.Name.Name)
			case token.CONST, token.VAR:
				s := v.Specs[0].(*ast.ValueSpec)
				for _, n := range s.Names {
					names = append(names, n.Name)
				}
			}
		}
		for _, name := range names {
			if ast.IsExported(name) {
				symbols = append(symbols, name)
			}
		}
	}
	return symbols
}

var _ = It("ensures complete coverage of the core dsl", func() {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, "../", nil, 0)
	Ω(err).ShouldNot(HaveOccurred())
	expectedSymbols := []string{}
	for fn, file := range pkgs["ginkgo"].Files {
		if fn == "../deprecated_dsl.go" {
			continue
		}
		expectedSymbols = append(expectedSymbols, ExtractSymbols(file)...)
	}

	actualSymbols := []string{}
	for _, pkg := range []string{"core", "reporting", "decorators", "table"} {
		pkgs, err := parser.ParseDir(fset, "./"+pkg, nil, 0)
		Ω(err).ShouldNot(HaveOccurred())
		for _, file := range pkgs[pkg].Files {
			actualSymbols = append(actualSymbols, ExtractSymbols(file)...)
		}
	}

	Ω(actualSymbols).Should(ConsistOf(expectedSymbols))
})

```


[Code See]()
```
package parser

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/mod/modfile"
)

// recursiveModulePath takes in the root of the project and a directory within
// that root, and it will search all directories starting with dir and ending
// with root to find a go.mod file. It returns the module path (which is
// retrieved from the go.mod), and the directory where the go.mod was found,
// which is the module root.
func recursiveModulePath(root, dir string) (string, string, error) {
	modFilePath := dir + "/go.mod"
	_, err := os.Stat(modFilePath)
	if err != nil && !os.IsNotExist(err) {
		return "", "", errors.WithStack(err)
	}

	if err == nil {
		// A go.mod file exists in this directory.
		mod, err := ioutil.ReadFile(modFilePath)
		if err != nil {
			return "", "", errors.WithStack(err)
		}
		return modfile.ModulePath(mod), dir, nil
	}

	if dir == root {
		// This means that we didn't find a go.mod file anywhere in the
		// directory tree, so this project might not be using Go modules.
		// Behavior without a go.mod is not fully tested. We could either throw
		// an error or just try to run it and see what happens. Sometimes it
		// does work (e.g. with the golang/go repo).
		return "", "", nil
	}

	// If we didn't find a go.mod in this directory, and we're not at the root
	// yet, go up one directory and look for a go.mod file there.
	return recursiveModulePath(root, filepath.Dir(dir))
}





package parser

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/pkg/errors"
)

type ParsedDir struct {
	// FileSet is the token.FileSet that was used to parse the directory. This
	// is used to get filenames from an AST node.
	FileSet *token.FileSet
	// ModulePath is the name of the module that is defined in a go.mod file.
	// This is used to resolve imports within the same module.
	ModulePath string
	ModuleRoot string
	// Packages is the return value of parser.ParseDir, where the map key is the
	// package name and the map value is the AST of the whole package (which is
	// a directory in Go).
	Packages map[string]*ast.Package
}

type Parser struct {
	root  string
	cache map[string]*ParsedDir
}

func New(root string) *Parser {
	return &Parser{
		root:  root,
		cache: map[string]*ParsedDir{},
	}
}

func (p *Parser) Parse(dir string) (*ParsedDir, error) {
	// First, we check the cache to see if we've already parsed this file, and
	// if we have, return the cached version instead.
	if parsedDir, ok := p.cache[dir]; ok {
		return parsedDir, nil
	}

	modulePath, moduleRoot, err := recursiveModulePath(p.root, dir)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, 0)
	if err != nil {
		// If we encounter an error when parsing, then it's probably not a
		// valid Go file, so we just skip it.
		p.cache[dir] = nil
		return nil, nil
	}

	p.cache[dir] = &ParsedDir{
		FileSet:    fset,
		ModulePath: modulePath,
		ModuleRoot: moduleRoot,
		Packages:   pkgs,
	}
	return p.cache[dir], nil
}


func TestParser_Parse(t *testing.T) {
	t.Run("parses a Go file", func(tt *testing.T) {
		root := "../testdata/simple-repo"
		dir := "../testdata/simple-repo/cmd/api"
		p := New(root)

		parsedDir, err := p.Parse(dir)
		require.NoError(tt, err)

		require.NotNil(tt, parsedDir)
		assert.NotNil(tt, parsedDir.FileSet)
		assert.NotNil(tt, parsedDir.Packages)
		assert.Len(tt, parsedDir.Packages, 1)
		assert.NotNil(tt, parsedDir.Packages["main"])
	})

	t.Run("returns a cached version if we've parsed the file before", func(tt *testing.T) {
		root := "../testdata/simple-repo"
		dir := "../testdata/simple-repo/cmd/api"
		p := New(root)

		firstParsedDir, err := p.Parse(dir)
		require.NoError(tt, err)
		require.NotNil(tt, firstParsedDir)

		secondParsedDir, err := p.Parse(dir)
		require.NoError(tt, err)
		require.NotNil(tt, firstParsedDir)

		// This asserts that the pointers are the same.
		assert.Equal(tt, firstParsedDir, secondParsedDir)
	})

	t.Run("returns nil for an invalid file", func(tt *testing.T) {
		root := "../testdata/simple-repo"
		dir := "../testdata/simple-repo/pkg/invalid"
		p := New(root)

		parsedDir, err := p.Parse(dir)
		require.NoError(tt, err)

		assert.Nil(tt, parsedDir)
	})
}
```
