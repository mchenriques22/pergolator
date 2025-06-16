module github.com/mchenriques22/pergolator

go 1.24.2

require (
	github.com/dave/jennifer v1.7.1
	github.com/fatih/structtag v1.2.0
	github.com/iancoleman/strcase v0.3.0
	github.com/mchenriques22/pergolator/modifiers v0.0.0-20250421081237-028940c20d93
	github.com/mchenriques22/pergolator/tree v0.0.1
	github.com/stretchr/testify v1.10.0
	github.com/urfave/cli/v3 v3.1.1
	golang.org/x/tools v0.32.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	golang.org/x/mod v0.24.0 // indirect
	golang.org/x/sync v0.13.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/mchenriques22/pergolator/tree => ./tree

replace github.com/mchenriques22/pergolator/modifiers => ./modifiers
