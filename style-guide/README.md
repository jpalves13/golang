# Style guide
Compiled from some references to codify the current best approaches.

## Table of contents
- [Naming](#naming)
    - [Function and method names](#function-and-method-names)
    - [Avoid global variables](#avoid-global-variables)
    - [Consistent header naming](#consistent-header-naming)
    - [Use meaningful variable names](#use-meaningful-variable-names)
- [Imports](#imports)
    - [Protos e stubs](#protos-and-stubs)
    - [Divide imports](#divide-imports)
    - [Use goimports](#use-goimports)
- [Errors](#errors)
    - [Handle Type Assertion Failures](#handle-type-assertion-failures)
    - [Add context to errors](#add-context-to-errors)
- [Avoid Mutable Globals](#)
- [Keep the happy path left](#keep-the-happy-path-left)
- [Function Grouping and Ordering](#function-and-method-names)
- [Don't over-interface](#dont-over-interface)
- [Main first](#main-first)
- [Structs](#structs)
	- [Use named structs](#use-named-structs)
	- [Avoid new keyword](#avoid-new-keyword)
- [Avoid magic numbers](#avoid-magic-numbers)
- [Testing](#testing)
	- [Use an assert library](#use-an-assert-libary)
- [Linting]

## Naming
## Functions and method
**Bad:**
```go
package yamlconfig

func ParseYAMLConfig(input string) (*Config, error)
```

**Good:**
```go
package yamlconfig

func Parse(input string) (*Config, error)
```

**Bad:**
```go
func (c *Config) WriteConfigTo(w io.Writer) (int64, error)
```

**Good:**
```go
func (c *Config) WriteTo(w io.Writer) (int64, error)
```

## Function and method names

**Bad:**
```go
package yamlconfig

func ParseYAMLConfig(input string) (*Config, error)
```

**Bad:**
```go
package yamlconfig

func Parse(input string) (*Config, error)
```


## Avoid global variables

**Bad:**
```go
var db *sql.DB

func main() {
	db = // ...
	http.HandleFunc("/drop", DropHandler)
	// ...
}

func DropHandler(w http.ResponseWriter, r *http.Request) {
	db.Exec("DROP DATABASE prod")
}
```

Global variables make testing and readability hard and every method has access
to them (even those, that don't need it).

**Good:**
```go
func main() {
	db := // ...
	handlers := Handlers{DB: db}
	http.HandleFunc("/drop", handlers.DropHandler)
	// ...
}

type Handlers struct {
	DB *sql.DB
}

func (h *Handlers) DropHandler(w http.ResponseWriter, r *http.Request) {
	h.DB.Exec("DROP DATABASE prod")
}
```
Use structs to encapsulate the variables and make them available only to those functions that actually need them by making them methods implemented for that struct.

Alternatively, higher-order functions can be used to inject dependencies via closures.
```go
func main() {
	db := // ...
	http.HandleFunc("/drop", DropHandler(db))
	// ...
}

func DropHandler(db *sql.DB) http.HandleFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		db.Exec("DROP DATABASE prod")
	}
}
```

If you really need global variables or constants, e.g., for defining errors or string constants, put them at the top of your file.

**Bad:**
```go
import "xyz"

func someFunc() {
	//...
}

const route = "/some-route"

func someOtherFunc() {
	// usage of route
}

var NotFoundErr = errors.New("not found")

func yetAnotherFunc() {
	// usage of NotFoundErr
}
```

**Good:**
```go
import "xyz"

const route = "/some-route"

var NotFoundErr = errors.New("not found")

func someFunc() {
	//...
}

func someOtherFunc() {
	// usage of route
}

func yetAnotherFunc() {
	// usage of NotFoundErr
}
```

## Consistent header naming
**Bad:**
```go
r.Header.Get("authorization")
w.Header.Set("Content-type")
w.Header.Set("content-type")
w.Header.Set("content-Type")
```

**Good:**
```go
r.Header.Get("Authorization")
w.Header.Set("Content-Type")
```

## Use meaningful variable names
Avoid single-letter variable names. They may seem more readable to you at the moment of writing but they make the code hard to understand for your colleagues and your future self.

**Bad:**
```go
func findMax(l []int) int {
	m := l[0]
	for _, n := range l {
		if n > m {
			m = n
		}
	}
	return m
}
```

**Good:**
```go
func findMax(inputs []int) int {
	max := inputs[0]
	for _, value := range inputs {
		if value > max {
			max = value
		}
	}
	return max
}
```
Single-letter variable names are fine in the following cases.
* They are absolute standard like ...
	* `t` in tests
	* `r` and `w` in http request handlers
	* `i` for the index in a loop
* They name the receiver of a method, e.g., `func (s *someStruct) myFunction(){}`

Of course also too long variables names like `createInstanceOfMyStructFromString` should be avoided.



## Imports

## Protos and stubs
**Bad:**
```go
import (
    foo_service_go_proto "path/to/package/foo_service_go_proto"
    foo_service_go_grpc "path/to/package/foo_service_go_grpc"
)
```

 **Good:**
```go
import (
    fspb "path/to/package/foo_service_go_proto"
    fsgrpc "path/to/package/foo_service_go_grpc"
)
```

## Divide imports

**Bad:**
```go
import (
	"encoding/json"
	"github.com/some/external/pkg"
	"fmt"
	"github.com/this-project/pkg/some-lib"
	"os"
)
```

**Good:**
```go
import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/bahlo/this-project/pkg/some-lib"

	"github.com/bahlo/another-project/pkg/some-lib"
	"github.com/bahlo/yet-another-project/pkg/some-lib"

	"github.com/some/external/pkg"
	"github.com/some-other/external/pkg"
)
```
## Use goimports

Only commit gofmt'd files. Use `goimports` for this to format/update the import statements as well.

[Linting](#linting)

## Errors

## Handle Type Assertion Failures
The single return value form of a type assertion will panic on an incorrect type. Therefore, always use the "comma ok" idiom.
**Bad:**
```go
t := i.(string)
```

**Good:**
```go
t, ok := i.(string)
if !ok {
  // handle the error gracefully
}
```


## Add context to errors

**Bad:**
```go
file, err := os.Open("foo.txt")
if err != nil {
	return err
}
```

Using the approach above can lead to unclear error messages because of missing
context.

**Good:**
```go
file, err := os.Open("foo.txt")
if err != nil {
	return fmt.Errorf("open foo.txt failed: %w", err)
}
```

## Avoid Mutable Globals
Avoid mutating global variables, instead opting for dependency injection. This applies to function pointers as well as other kinds of values.

**Bad:**
```go
var timeNow = time.Now

func sign(msg string) string {
  now := timeNow()
  return signWithTime(msg, now)
}
```

**Good:**
```go
type signer struct {
  now func() time.Time
}

func newSigner() *signer {
  return &signer{
    now: time.Now,
  }
}

func (s *signer) Sign(msg string) string {
  now := s.now()
  return signWithTime(msg, now)
}
```

## Keep the happy path left

**Bad:**
```go
if item, ok := someMap[someKey]; ok {
	return item
}
return ErrKeyNotFound
```

**Good:**
```go
item, ok := someMap[someKey]
if !ok {
	return ErrKeyNotFound
}
return item
```

This helps to keep your code clear and readable. Not doing it accumulates in 
larger functions and leads to the happy path being buried in a lot of if/for/... 
statements.

## Function Grouping and Ordering

**Bad:**
```go
func (s *something) Cost() {
  return calcCost(s.weights)
}

type something struct{ ... }

func calcCost(n []int) int {...}

func (s *something) Stop() {...}

func newSomething() *something {
    return &something{}
}
```

**Good:**
```go
type something struct{ ... }

func newSomething() *something {
    return &something{}
}

func (s *something) Cost() {
  return calcCost(s.weights)
}

func (s *something) Stop() {...}

func calcCost(n []int) int {...}
```

## Don't over-interface

**Bad:**
```go
type Server interface {
	Serve() error
	Some() int
	Fields() float64
	That() string
	Are([]byte) error
	Not() []string
	Necessary() error
}

func debug(srv Server) {
	fmt.Println(srv.String())
}

func run(srv Server) {
	srv.Serve()
}
```

**Good:**
```go
type Server interface {
	Serve() error
}

func debug(v fmt.Stringer) {
	fmt.Println(v.String())
}

func run(srv Server) {
	srv.Serve()
}
```

Favour small interfaces and only expect the interfaces you need in your funcs.

## Main first

**Bad:**
```go
package main // import "github.com/me/my-project"

func someHelper() int {
	// ...
}

func someOtherHelper() string {
	// ...
}

func Handler(w http.ResponseWriter, r *http.Reqeust) {
	// ...
}

func main() {
	// ...
}
```

**Good:**
```go
package main // import "github.com/me/my-project"

func main() {
	// ...
}

func Handler(w http.ResponseWriter, r *http.Reqeust) {
	// ...
}

func someHelper() int {
	// ...
}

func someOtherHelper() string {
	// ...
}
```

Putting `main()` first makes reading the file a lot easier. Only the
`init()` function should be above it.

Divide imports into four groups sorted from internal to external for readability:
1. Standard library
2. Project internal packages
3. Company internal packages
4. External packages



## Structs
### Use named structs
If a struct has more than one field, include field names when instantiating it.

**Bad:**
```go
params := myStruct{
	1, 
	true,
}
```

**Good:**
```go
params := myStruct{
	Foo: 1,
	Bar: true,
}
```

### Avoid new keyword
Using the normal syntax instead of the `new` keyword makes it more clear what is happening: a new instance of the struct is created `MyStruct{}` and we get the pointer for it with `&`.

**Bad:**
```go
s := new(MyStruct)
```

**Good:**
```go
s := &MyStruct{}
```

## Avoid magic numbers
A number without a name and any context is just a random value. It tells us nothing, so avoid them in your code (the exception might be the number 0, for example when creating loops). 

**Bad:**
```go
func IsStrongPassword(password string) bool {
	return len(password) >= 8
}
```

**Good:**
```go
const minPasswordLength = 8

func IsStrongPassword(password string) bool {
	return len(password) >= minPasswordLength
}
```


## Testing

### Use an assert library

**Bad:**
```go
func TestAdd(t *testing.T) {
	actual := 2 + 2
	expected := 4
	if (actual != expected) {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
```

**Good:**
```go
import "github.com/stretchr/testify/assert"

func TestAdd(t *testing.T) {
	actual := 2 + 2
	expected := 4
	assert.Equal(t, expected, actual)
}
```

Using assert libraries makes your tests more readable, requires less code and
provides consistent error output.

### Use sub-tests to structure functional tests
**Bad:**
```go
func TestSomeFunctionSuccess(t *testing.T) {
	// ...
}

func TestSomeFunctionWrongInput(t *testing.T) {
	// ...
}
```

**Good:**
```go
func TestSomeFunction(t *testing.T) {
	t.Run("success", func(t *testing.T){
		//...
	})

	t.Run("wrong input", func(t *testing.T){
		//...
	})
}
```

## Linting
- [errcheck](https://github.com/kisielk/errcheck) to ensure that errors are handled
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports) to format code and manage imports
- [golint](https://github.com/golang/lint) to point out common style mistakes
- [govet](https://pkg.go.dev/cmd/vet) to analyze code for common mistakes
- [staticcheck](https://staticcheck.dev) to do various static analysis checks
- [golangci-lint](https://github.com/golangci/golangci-lint) to lint your projects before committing


## Referências
[bahlo](https://github.com/bahlo/go-styleguide)

[google](https://google.github.io/styleguide/go/index)

[uber](https://github.com/uber-go/guide/blob/master/style.md)

[awesome-go](https://awesome-go.com/style-guides/)