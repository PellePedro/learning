
[Text Templaytes](https://github.com/RangelReale/gotexttemplate/tree/master/util/goconv/cmd/goconv)

[mobile](https://github.com/guyezi/mobile/blob/3c8601c510d0503ac84d1e5cb8e24de550201dea/cmd/gobind/main.go#L69)<br/>
[Mobile Internal] (https://github.com/guyezi/mobile/blob/3c8601c510d0503ac84d1e5cb8e24de550201dea/internal/importers/ast.go)<br/>
[GBB](https://github.com/voidint/gbb/tree/master)<br/>
[Dispel](https://github.com/vincent-petithory/dispel)<br/>
[Code See](https://github.com/Codesee-io/codesee-deps-go)

<details>
  <summary>P1</summary>
	
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
	
</details>


<details>
  <summary>P1</summary>

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
</details>
	
	
[Code See](https://github.com/Codesee-io/codesee-deps-go)

<details>
  <summary>Code See</summary>
	
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
```
</details>




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


func determineGoDirectories(root string) ([]string, error) {
	dirSet := map[string]struct{}{}

	err := godirwalk.Walk(root, &godirwalk.Options{
		Callback: func(path string, de *godirwalk.Dirent) error {
			segments := strings.Split(path, "/")
			for _, segment := range segments {
				// Skip over the entire .git directory to speed up the walking.
				if segment == ".git" {
					return godirwalk.SkipThis
				}
				// Skip over any vendored dependencies since we don't care about
				// external dependencies.
				if segment == "vendor" {
					return godirwalk.SkipThis
				}
			}

			// We're looking for Go files, so if this is a directory, skip over
			// it.
			if de.IsDir() {
				return nil
			}
			// If this file isn't a Go file, skip it.
			if !strings.HasSuffix(path, ".go") {
				return nil
			}

			// We've found a Go file, so we should add its directory to the set
			// of directories.
			dir := filepath.Dir(path)
			dirSet[dir] = struct{}{}

			return nil
		},
		Unsorted: true,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Convert the set to a slice.
	dirs := make([]string, 0, len(dirSet))
	for dir := range dirSet {
		dirs = append(dirs, dir)
	}

	return dirs, nil
}

func TestDetermineGoDirectories(t *testing.T) {
	t.Run("handles a simple repo", func(tt *testing.T) {
		root := filepath.Clean("../testdata/simple-repo")

		dirs, err := determineGoDirectories(root)
		require.NoError(tt, err)

		// Sort the slice since its order isn't deterministic.
		sort.Slice(dirs, func(i, j int) bool {
			return dirs[i] < dirs[j]
		})
		assert.Equal(tt, []string{
			"../testdata/simple-repo/cmd/api",
			"../testdata/simple-repo/pkg/handlers",
			"../testdata/simple-repo/pkg/invalid",
			"../testdata/simple-repo/pkg/server",
			"../testdata/simple-repo/pkg/signals",
		}, dirs)
	})
}
```


```
package links

import (
	"fmt"
	"go/ast"
	"path/filepath"
	"sort"
	"strings"

	"github.com/Codesee-io/codesee-deps-go/pkg/parser"
	"github.com/pkg/errors"
)

type Link struct {
	From string `json:"from"`
	To   string `json:"to"`
}

// This type aliases are only used to make some maps a bit more readable. They
// aren't actually necessary to work correctly.
type (
	// PackagePath e.g. github.com/Codesee-io/codesee-deps-go/pkg/parser
	PackagePath string
	// PackageName e.g. parser
	PackageName string
	// Identifier e.g. New or ParsedDir
	Identifier string
	// Filename e.g. /root/codesee-deps-go/pkg/parser/parser.go
	Filename string
)

// DetermineLinks takes in a root directory and generates all the links between
// the Go files in this directory, relative from this root directory. The order
// of links is not guaranteed to be deterministic to make it faster. If you're
// asserting equality for the links (e.g. in a test), make sure you sort it
// before your assertion.
func DetermineLinks(root string) ([]Link, error) {
	absRoot, err := filepath.Abs(root)
	dirs, err := determineGoDirectories(absRoot)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	linksSet := map[string]struct{}{}
	links := []Link{}

	p := parser.New(absRoot)

	// We determine the links for a project by making 2 passes over the
	// directories.

	// This is a mapping from package path to a package name. This is needed to
	// help generate the reverse mapping (name to path) for a specific file.
	// More details about why we need to do this can be found in the second
	// pass.
	pkgPathToPkgName := map[PackagePath]PackageName{}
	// This is a mapping from package path and object name in scope to the
	// filename that it's defined in. So this will look something like this:
	// {
	//   "github.com/Codesee-io/codesee-deps-go/pkg/parser": {
	//     "New": "/root/codesee-deps-go/pkg/parser/parser.go"
	//   }
	// }
	identifierToFilename := map[PackagePath]map[Identifier]Filename{}

	// This first pass populates pkgPathToPkgName and identifierToFilename.
	for _, dir := range dirs {
		parsedDir, err := p.Parse(dir)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if parsedDir == nil {
			// This package wasn't able to be parsed correctly, so we just skip
			// it.
			continue
		}

		// In most cases, there's only one package per directly, though this
		// isn't guaranteed.
		for pkgName, pkg := range parsedDir.Packages {
			// Add this package in our mapping from package path to package
			// name.
			pkgPath := PackagePath(strings.Replace(dir, parsedDir.ModuleRoot, parsedDir.ModulePath, -1))
			pkgPathToPkgName[pkgPath] = PackageName(pkgName)

			for _, file := range pkg.Files {
				pos := parsedDir.FileSet.Position(file.Pos())
				filename := Filename(pos.Filename)

				// For each file, go through all the objects that are in the
				// global scope (e.g. types, functions, const and var
				// declarations, etc.) and add them to our mapping from
				// identifier to filename.
				for name := range file.Scope.Objects {
					if _, ok := identifierToFilename[pkgPath]; !ok {
						identifierToFilename[pkgPath] = map[Identifier]Filename{}
					}
					identifierToFilename[pkgPath][Identifier(name)] = filename
				}
			}
		}
	}

	// This is a mapping from filename to the package path and object name that
	// is being used in that file. This is how we'll know exactly what is being
	// used in the imported package. This is necessary to be able to map back to
	// where a specific object is defined. So this will look something like
	// this:
	// {
	//   "/root/codesee-deps-go/pkg/links/links.go": {
	//     "github.com/Codesee-io/codesee-deps-go/pkg/parser": {
	//       "New": {}
	//     }
	//   }
	// }
	filenameToIdentifierUsed := map[Filename]map[PackagePath]map[Identifier]struct{}{}

	// This second pass populates filenameToIdentifierUsed.
	for _, dir := range dirs {
		parsedDir, err := p.Parse(dir)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		if parsedDir == nil {
			// This package wasn't able to be parsed correctly, so we just skip
			// it.
			continue
		}

		for _, pkg := range parsedDir.Packages {
			pkgPath := PackagePath(strings.Replace(dir, parsedDir.ModuleRoot, parsedDir.ModulePath, -1))

			for _, file := range pkg.Files {
				// This is on a per-file basis since each file can have an alias
				// for an import.
				pkgNameToPkgPath := map[PackageName]PackagePath{}

				pos := parsedDir.FileSet.Position(file.Pos())
				filename := Filename(pos.Filename)

				dotImports := []PackagePath{}

				// Go through all the imports in this file.
				for _, importSpec := range file.Imports {
					// The path value is wrapped in quotes, so we need to trim
					// them.
					importedPkgPath := strings.Trim(importSpec.Path.Value, "\"")

					if !strings.Contains(importedPkgPath, parsedDir.ModulePath) {
						continue
					}

					// The import spec's name is only defined if it's been
					// aliased to a different name, like this:
					// import (
					//   c "github.com/a/b"
					// )
					// In this example, the name would be "c". If there is no
					// alias, then name is nil.
					if importSpec.Name == nil {
						// If there isn't a custom package name, then we need to
						// use the package's assigned name. While this is
						// usually the final segment in the package path, this
						// isn't guaranteed. So that's why we need to use the
						// mapping from package path to package name that we
						// generated in the first pass to fill in the default
						// package name.
						if name, ok := pkgPathToPkgName[PackagePath(importedPkgPath)]; ok {
							// The imported package path is found in our
							// mapping, which means this is an internal import,
							// not an external dependency.
							pkgNameToPkgPath[name] = PackagePath(importedPkgPath)
						}
					} else {
						if importSpec.Name.String() == "." {
							// If the name is ".", then all of that package's
							// identifiers are accessible without needing to
							// qualify it with a package name. Here's an
							// example:
							// import (
							//   . "fmt"
							// )
							// With this, we can then use Println and Printf
							// instead of fmt.Println and fmt.Printf.
							dotImports = append(dotImports, PackagePath(importedPkgPath))
						} else {
							pkgNameToPkgPath[PackageName(importSpec.Name.String())] = PackagePath(importedPkgPath)
						}
					}
				}

				// These are all the packages that we should check for
				// unresolved identifiers. If an identifier is being used
				// without a package name, that means it's either defined in its
				// own package, or it was imported with a ".".
				packagePathsWithUnqualifiedIdentifiers := append([]PackagePath{pkgPath}, dotImports...)

				for _, ident := range file.Unresolved {
					for _, pkgPath := range packagePathsWithUnqualifiedIdentifiers {
						if toFilename, ok := identifierToFilename[pkgPath][Identifier(ident.String())]; ok {
							setKey := fmt.Sprintf("%s:%s", filename, toFilename)

							if _, ok := linksSet[setKey]; !ok {
								links = append(links, Link{
									From: strings.Replace(string(filename), absRoot+"/", "", -1),
									To:   strings.Replace(string(toFilename), absRoot+"/", "", -1),
								})
								linksSet[setKey] = struct{}{}
							}
						}
					}
				}

				// We've gotten all the intra-package resolutions, but we can't
				// rely on the Unresolved portion of the file AST for all of
				// them because it doesn't show the fully unresolved path e.g.
				// for parser.New, it will only tell us that parser is
				// unresolved, so we don't know what in parser was actually
				// used. To find those, we walk the AST to find all selector
				// expressions.
				ast.Inspect(file, func(n ast.Node) bool {
					// A selector expression is an expression in the format of
					// "X.Selector" (e.g. parser.New, p.Parse, parser.ParsedDir,
					// etc.). This is the main way that we'll determine how an
					// imported package is being used.
					selectorExpr, ok := n.(*ast.SelectorExpr)
					if !ok {
						return true
					}

					// X in our cause will be the package name.
					xIdent, ok := selectorExpr.X.(*ast.Ident)
					if !ok {
						return true
					}

					usedPkgName := PackageName(xIdent.String())
					// Sel in our cause is the identifier.
					usedIdentifier := Identifier(selectorExpr.Sel.String())

					usedPkgPath, ok := pkgNameToPkgPath[usedPkgName]
					if !ok {
						// This used package name is not found in our mapping,
						// which means this is not an internal import, but an
						// external dependency instead.
						return true
					}

					if _, ok := filenameToIdentifierUsed[filename]; !ok {
						filenameToIdentifierUsed[filename] = map[PackagePath]map[Identifier]struct{}{}
					}
					if _, ok := filenameToIdentifierUsed[filename][usedPkgPath]; !ok {
						filenameToIdentifierUsed[filename][usedPkgPath] = map[Identifier]struct{}{}
					}
					filenameToIdentifierUsed[filename][usedPkgPath][usedIdentifier] = struct{}{}

					return true
				})
			}
		}
	}

	// Now that we've pulled all the necessary data out of all the Go ASTs, we
	// can piece together a comprehensive list of links from file to file.
	for fromFilename, pkgPathsUsed := range filenameToIdentifierUsed {
		for pkgPath, identifiersUsed := range pkgPathsUsed {
			identifiersDefined, ok := identifierToFilename[pkgPath]
			if !ok {
				// We don't have the identifiers for this package path. This is
				// probably an external dependency.
				continue
			}

			for identifierUsed := range identifiersUsed {
				toFilename, ok := identifiersDefined[identifierUsed]
				if !ok {
					// We found an identifier being used by this package, but
					// that identifier isn't defined in this package. This could
					// be a Go file that wouldn't compile, or it could mean that
					// we missed adding it. Either way, we don't want it
					// interfering with all the other links, so we just skip it.
					continue
				}
				setKey := fmt.Sprintf("%s:%s", fromFilename, toFilename)

				if _, ok := linksSet[setKey]; !ok {
					links = append(links, Link{
						From: strings.Replace(string(fromFilename), absRoot+"/", "", -1),
						To:   strings.Replace(string(toFilename), absRoot+"/", "", -1),
					})
					linksSet[setKey] = struct{}{}
				}
			}
		}
	}

	// Sort the slice since its order isn't deterministic. While it doesn't need
	// to be sorted, it helps if it is. And it's probably faster to sort it here
	// than to do it downstream.
	sort.Slice(links, func(i, j int) bool {
		if links[i].From == links[j].From {
			return links[i].To < links[j].To
		}
		return links[i].From < links[j].From
	})

	return links, nil
}

func TestDetermineLinks(t *testing.T) {
	t.Run("determines links for a simple repo", func(tt *testing.T) {
		root := "../testdata/simple-repo"

		links, err := DetermineLinks(root)
		require.NoError(tt, err)

		assert.Equal(tt, []Link{
			{From: "cmd/api/main.go", To: "pkg/server/server.go"},
			{From: "cmd/api/main.go", To: "pkg/signals/signals.go"},
			{From: "pkg/server/server.go", To: "pkg/handlers/handlers.go"},
			{From: "pkg/signals/signals_test.go", To: "pkg/signals/signals.go"},
		}, links)
	})
}
```




[Doc extractor](https://github.com/joeshaw/doc-extract)




```
		docs := make(map[string]string)
		fset := token.NewFileSet() // positions are relative to fset
		d, err := parser.ParseDir(fset, "./handlers", nil, parser.ParseComments)
		if err != nil {
			log.Fatalln(err)
		}
		for _, f := range d {
			p := doc.New(f, "./", 2)
			// 获取所有func doc
			for _, f := range p.Funcs {
				if strings.HasSuffix(f.Name, "Perm") {
					funcName := fmt.Sprintf("diamond/handlers.%v", f.Name)
					docs[funcName] = f.Doc
				}
			}
		}
```

G

```

func (p *Parser) parseDir(dir string) error {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info os.FileInfo) bool {
		valid, name := true, info.Name()

		if p.ignoreTests {
			if strings.HasSuffix(name, testSuffix) {
				valid = false
			}
		}

		if len(p.ignore) != 0 {
			match, err := regexp.MatchString(p.ignore, dir+name)
			if err != nil {
				log.Fatal(err)
				return true
			}
			if match {
				valid = false
			}
		}

		return valid
	}, 0)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		for fn, f := range pkg.Files {
			ast.Walk(&treeVisitor{
				fileSet:     fset,
				packageName: pkg.Name,
				fileName:    fn,
				p:           p,
			}, f)
		}
	}

	return nil
}
```



```
package goconst

import (
	"go/ast"
	"go/token"
	"strings"
)

// treeVisitor carries the package name and file name
// for passing it to the imports map, and the fileSet for
// retrieving the token.Position.
type treeVisitor struct {
	p                     *Parser
	fileSet               *token.FileSet
	packageName, fileName string
}

// Visit browses the AST tree for strings that could be potentially
// replaced by constants.
// A map of existing constants is built as well (-match-constant).
func (v *treeVisitor) Visit(node ast.Node) ast.Visitor {
	if node == nil {
		return v
	}

	// A single case with "ast.BasicLit" would be much easier
	// but then we wouldn't be able to tell in which context
	// the string is defined (could be a constant definition).
	switch t := node.(type) {
	// Scan for constants in an attempt to match strings with existing constants
	case *ast.GenDecl:
		if !v.p.matchConstant {
			return v
		}
		if t.Tok != token.CONST {
			return v
		}

		for _, spec := range t.Specs {
			val := spec.(*ast.ValueSpec)
			for i, str := range val.Values {
				lit, ok := str.(*ast.BasicLit)
				if !ok || !v.isSupported(lit.Kind) {
					continue
				}

				v.addConst(val.Names[i].Name, lit.Value, val.Names[i].Pos())
			}
		}

	// foo := "moo"
	case *ast.AssignStmt:
		for _, rhs := range t.Rhs {
			lit, ok := rhs.(*ast.BasicLit)
			if !ok || !v.isSupported(lit.Kind) {
				continue
			}

			v.addString(lit.Value, rhs.(*ast.BasicLit).Pos())
		}

	// if foo == "moo"
	case *ast.BinaryExpr:
		if t.Op != token.EQL && t.Op != token.NEQ {
			return v
		}

		var lit *ast.BasicLit
		var ok bool

		lit, ok = t.X.(*ast.BasicLit)
		if ok && v.isSupported(lit.Kind) {
			v.addString(lit.Value, lit.Pos())
		}

		lit, ok = t.Y.(*ast.BasicLit)
		if ok && v.isSupported(lit.Kind) {
			v.addString(lit.Value, lit.Pos())
		}

	// case "foo":
	case *ast.CaseClause:
		for _, item := range t.List {
			lit, ok := item.(*ast.BasicLit)
			if ok && v.isSupported(lit.Kind) {
				v.addString(lit.Value, lit.Pos())
			}
		}

	// return "boo"
	case *ast.ReturnStmt:
		for _, item := range t.Results {
			lit, ok := item.(*ast.BasicLit)
			if ok && v.isSupported(lit.Kind) {
				v.addString(lit.Value, lit.Pos())
			}
		}
	}

	return v
}

// addString adds a string in the map along with its position in the tree.
func (v *treeVisitor) addString(str string, pos token.Pos) {
	str = strings.Replace(str, `"`, "", 2)

	// Ignore empty strings
	if len(str) == 0 {
		return
	}

	if len(str) < v.p.minLength {
		return
	}

	_, ok := v.p.strs[str]
	if !ok {
		v.p.strs[str] = make([]ExtendedPos, 0)
	}
	v.p.strs[str] = append(v.p.strs[str], ExtendedPos{
		packageName: v.packageName,
		Position:    v.fileSet.Position(pos),
	})
}

// addConst adds a const in the map along with its position in the tree.
func (v *treeVisitor) addConst(name string, val string, pos token.Pos) {
	val = strings.Replace(val, `"`, "", 2)
	v.p.consts[val] = ConstType{
		Name:        name,
		packageName: v.packageName,
		Position:    v.fileSet.Position(pos),
	}
}

func (v *treeVisitor) isSupported(tk token.Token) bool {
	for _, s := range v.p.supportedTokens {
		if tk == s {
			return true
		}
	}
	return false
}
```


