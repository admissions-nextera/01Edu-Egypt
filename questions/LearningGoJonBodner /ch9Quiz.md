# 📘 Learning Go — Chapter 9 Quiz
## Modules, Packages, and Imports

**Questions:** 18 | **Time:** 25 minutes | **Passing Score:** 14/18 (78%)

---

### Q1: What is the difference between a module and a package?

**A)** They are synonyms  
**B)** A module is a collection of packages with a shared `go.mod`; a package is a collection of `.go` files in one directory that share a namespace  
**C)** A package is larger than a module  
**D)** Modules are only for open-source projects  

<details><summary>💡 Answer</summary>

**B) Module = collection of packages under a `go.mod`; Package = one directory**

```
mymodule/               ← module root (has go.mod)
├── go.mod
├── main.go             ← package main
├── handlers/           
│   └── handlers.go     ← package handlers
└── models/
    └── models.go       ← package models
```

One module, three packages. You `go get` modules; you `import` packages.

</details>

---

### Q2: What naming convention does the book recommend for packages?

**A)** CamelCase names like `UserService`  
**B)** Short, lowercase, single-word names that don't stutter — `http`, `json`, `io`, not `httpClient` or `userservice`  
**C)** Prefix with the module name  
**D)** Match the directory name with underscores for multi-word names  

<details><summary>💡 Answer</summary>

**B) Short, lowercase, single word — no stuttering**

"Stuttering" means `user.UserService` — the package name and type name repeat. Instead: `user.Service`. The package name is part of the API: `http.Client`, `json.Decoder`, `io.Reader`. Short names + descriptive type names is the Go convention.

</details>

---

### Q3: What is the package name for `main` and why is it special?

**A)** Any package name works for the entry point  
**B)** `package main` — the only package that can contain the `main()` function; Go only compiles it as a standalone executable (not an importable library)  
**C)** Only the first file needs `package main`  
**D)** `main` is a reserved keyword for all packages  

<details><summary>💡 Answer</summary>

**B) `package main` + `func main()` = executable entry point**

Every `.go` file must declare its package. Files with `package main` that contain `func main()` become executables. Packages with any other name become importable libraries. A project can have only one `package main` per binary.

</details>

---

### Q4: What does it mean for a function/variable to be "exported" from a package?

**A)** Stored in a special export file  
**B)** Its name starts with an uppercase letter — accessible from other packages  
**C)** It has a `public` keyword  
**D)** It appears in the package documentation  

<details><summary>💡 Answer</summary>

**B) Uppercase first letter = exported (public)**

```go
package user

type User struct {                // exported
    Name     string              // exported field
    password string              // unexported — not accessible outside package
}

func New(name string) User { ... }  // exported
func validate(u User) bool { ... }  // unexported
```

This is Go's entire visibility system — one rule, no `public`/`private`/`protected` keywords.

</details>

---

### Q5: What is the `init` function and what happens if a package has multiple of them?

**A)** Must be the first function in a file; only one allowed  
**B)** An optional function that runs once when the package is imported; a package can have multiple `init` functions (even in the same file) and they all run in order  
**C)** Only runs in `package main`  
**D)** Must be exported  

<details><summary>💡 Answer</summary>

**B) Runs on import; multiple allowed per package**

```go
// File 1:
func init() { registerDrivers() }

// File 2 in same package:
func init() { initMetrics() }
```

Both run. The book warns: use `init` sparingly — it runs invisibly and can cause surprising behavior. Side effects in `init` are hard to test.

</details>

---

### Q6: What does the blank import `import _ "package/path"` do?

**A)** Imports the package without using any of its symbols — triggers its `init` function(s)  
**B)** Imports all public symbols  
**C)** Disables the package  
**D)** Creates a compile error  

<details><summary>💡 Answer</summary>

**A) Triggers `init` without using any symbols**

```go
import _ "github.com/lib/pq"  // registers postgres driver in database/sql via init()
```

This is the standard pattern for SQL drivers, image decoders, and other registerable plugins. The package's `init` registers itself with a central registry. The blank import says: "I need this side effect but not its API."

</details>

---

### Q7: What is an alias import and when is it used?

**A)** `import alias "path"` — access the package as `alias.Function()`  
**B)** It's not supported in Go  
**C)** Required for all imports  
**D)** Only for standard library packages  

<details><summary>💡 Answer</summary>

**A) `import alias "path"` — rename the package for local use**

```go
import (
    gofmt "go/format"            // avoid conflict with local `format` variable
    rand "math/rand"             // if two packages are both named `rand`
)
```

Common uses: resolve naming conflicts, shorten a long package name, clarity when working with multiple packages that have the same last path element.

</details>

---

### Q8: What is the `internal` package mechanism?

**A)** Packages named `internal` are only accessible within the module  
**B)** Packages at or below an `internal` directory can only be imported by code in the parent of the `internal` directory — not by external users of your module  
**C)** `internal` packages are never tested  
**D)** `internal` is just a naming convention with no enforcement  

<details><summary>💡 Answer</summary>

**B) Enforced visibility boundary**

```
mymodule/
├── internal/
│   └── helpers/helpers.go   ← only importable by mymodule code
├── api/api.go               ← can import internal/helpers
└── go.mod
```

External modules cannot import `internal` packages. This lets you share code within your module without exposing it as public API. The compiler enforces this — import attempts from outside the boundary fail.

</details>

---

### Q9: What does `go get github.com/user/repo@v1.2.3` do?

**A)** Downloads the source code to `$GOPATH/src`  
**B)** Adds the dependency to `go.mod` at the specified version and downloads it to the module cache  
**C)** Installs the package as a CLI tool  
**D)** Creates a fork of the repository  

<details><summary>💡 Answer</summary>

**B) Updates `go.mod` with the dependency at the specified version**

```
go get github.com/user/repo@v1.2.3
go get github.com/user/repo@latest
go get github.com/user/repo@main
```

The downloaded module is stored in the module cache (`$GOPATH/pkg/mod`). `go.mod` records the minimum version. `go.sum` records the hash. Run `go mod tidy` afterward to clean up unused dependencies.

</details>

---

### Q10: What does `go mod tidy` do?

**A)** Formats the `go.mod` file  
**B)** Adds missing dependencies and removes unused ones from `go.mod` and `go.sum`  
**C)** Upgrades all dependencies to latest versions  
**D)** Downloads all dependencies  

<details><summary>💡 Answer</summary>

**B) Adds missing; removes unused — synchronizes `go.mod` with actual imports**

Run `go mod tidy` after: adding new imports, removing imports, or changing versions. It's safe to run repeatedly. The book recommends running it before committing. CI should verify `go.mod` and `go.sum` are tidy.

</details>

---

### Q11: What is semantic versioning (semver) and how does Go use it?

**A)** A timestamp-based versioning scheme  
**B)** `vMAJOR.MINOR.PATCH` — Go's module system uses it; breaking changes require a major version bump  
**C)** Only used for major releases  
**D)** Go doesn't use semver  

<details><summary>💡 Answer</summary>

**B) `vMAJOR.MINOR.PATCH` — breaking changes = major bump**

- `MAJOR` bump: breaking API change
- `MINOR` bump: new backward-compatible features  
- `PATCH` bump: backward-compatible bug fixes

In Go modules, v2+ modules must have `/v2` at the end of their module path: `github.com/user/repo/v2`. This allows different major versions to coexist in one program.

</details>

---

### Q12: All files in a directory must have the same package name (with one exception). What is the exception?

**A)** `main.go` can have a different package name  
**B)** Test files (`_test.go`) can use either the package name OR `packagename_test` for black-box testing  
**C)** Files starting with `_` are excluded  
**D)** There is no exception  

<details><summary>💡 Answer</summary>

**B) Test files can use `packagename_test` for external test packages**

```go
// handlers_test.go
package handlers_test  // external test package — can only use exported API

// OR:
package handlers       // internal test — has access to unexported identifiers
```

`package handlers_test` forces you to test the public API only, which often produces better tests. Both styles are valid and coexist in the same directory.

</details>

---

### Q13: What is the minimum version selection (MVS) algorithm?

**A)** Always uses the latest version of each dependency  
**B)** Selects the minimum version of each dependency that satisfies all requirements — predictable, reproducible builds  
**C)** Uses the newest stable version for security  
**D)** Lets each developer choose their preferred version  

<details><summary>💡 Answer</summary>

**B) Minimum version satisfying all requirements — reproducible builds**

If your module requires `lib@v1.2` and another dependency requires `lib@v1.3`, MVS selects `v1.3` (minimum that satisfies both). It never automatically upgrades to v1.4. This makes builds reproducible: the same `go.mod` always produces the same dependency graph. No "npm dependency hell."

</details>

---

### Q14: What is the `replace` directive in `go.mod`?

**A)** Renames a dependency  
**B)** Redirects an import to a different module or local path — useful for local development and forked dependencies  
**C)** Deprecates a dependency  
**D)** Replaces one package with another at compile time  

<details><summary>💡 Answer</summary>

**B) Redirects to local path or fork**

```go
// go.mod
replace github.com/original/lib => ../local-fork
replace github.com/original/lib => github.com/myfork/lib v1.0.0
```

Common uses: testing local changes to a dependency without publishing, using a fork with a critical fix. Remove `replace` before publishing your own module — it won't apply to your module's consumers.

</details>

---

### Q15: If two imported packages have the same name, what must you do?

**A)** Go automatically resolves the conflict  
**B)** Use import aliases to rename at least one of them  
**C)** You cannot import both — one must be removed  
**D)** Use the full path to distinguish them  

<details><summary>💡 Answer</summary>

**B) Import alias for at least one**

```go
import (
    mathrand "math/rand"
    cryptorand "crypto/rand"
)
```

Without aliases, both would be `rand` — a compile error. The alias applies only to the current file. Choose meaningful aliases that clarify which `rand` you're using.

</details>

---

### Q16: What does a `go.sum` file protect against?

**A)** Duplicate imports  
**B)** Tampered or modified module versions — stores cryptographic hashes of downloaded content; any change is detected  
**C)** Version conflicts  
**D)** Unauthorized network access  

<details><summary>💡 Answer</summary>

**B) Tamper detection via cryptographic hashes**

Even if a module author modifies what `v1.2.3` contains (by retagging a different commit), your `go.sum` will detect the mismatch and refuse to build. This is a supply chain security guarantee. Always commit `go.sum` and never manually edit it.

</details>

---

### Q17: What is a "vendor" directory and when would you use it?

**A)** A directory for compiled libraries  
**B)** A copy of all dependencies stored inside your module — created with `go mod vendor`; allows building without internet access or in air-gapped environments  
**C)** Where Go stores downloaded modules globally  
**D)** Required for all production modules  

<details><summary>💡 Answer</summary>

**B) Vendored dependencies — offline building**

```bash
go mod vendor        # copies all dependencies into ./vendor
go build -mod=vendor # uses vendor/ instead of module cache
```

Reasons to vendor: air-gapped CI, ensure exact reproducibility (no dependency on external servers), compliance requirements. The book mentions it but says most projects should use the module cache instead.

</details>

---

### Q18: What happens when you import a package but don't use any of its exported identifiers?

**A)** A warning is printed  
**B)** Compile error — unused imports are always errors in Go (except blank imports `_`)  
**C)** The import is silently ignored  
**D)** Works fine  

<details><summary>💡 Answer</summary>

**B) Compile error**

This is the same philosophy as unused variables — enforced by the compiler. The only exception: `import _ "pkg"` (blank import) explicitly signals "I need this for its side effects, not its API." `goimports` automatically manages imports, so this rarely causes problems in practice.

</details>

---

## 📊 Score

| Score | Result |
|---|---|
| 17–18 ✅ | **Excellent.** Modules and packages are clear. |
| 14–16 ✅ | **Ready.** Review `go mod tidy` and the `internal` package. |
| 10–13 ⚠️ | **Study `go.mod`/`go.sum` purpose and the visibility rules.** |
| Below 10 ❌ | **Reread Chapter 9 — modules affect every project you create.** |