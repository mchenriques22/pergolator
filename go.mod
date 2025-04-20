module github.com/antoninferrand/pergolator

go 1.24.2

require (
	github.com/antoninferrand/pergolator/tree v0.0.0-00000000000000-000000000000
	github.com/dave/jennifer v1.7.1
	github.com/fatih/structtag v1.2.0
	github.com/stretchr/testify v1.10.0
	github.com/urfave/cli/v3 v3.1.1
	golang.org/x/tools v0.32.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/antoninferrand/pergolator/tree => ./tree

replace github.com/antoninferrand/pergolator/modifiers => ./modifiers
