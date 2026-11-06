# Bundle Update Log
* **Creation**: Added the [Go Packages](/01-getting-started/go-packages/) topic: [What package main Means](/01-getting-started/go-packages/package-main.md) (executable vs reusable packages, startup order) and [Package vs Module vs Workspace](/01-getting-started/go-packages/package-module-workspace.md) (with a `go.work` demo).
* **Update**: Moved [Modules and Packages](/01-getting-started/go-packages/modules-and-packages.md) into the Go Packages topic and updated every link to it.
* **Creation**: Added the [Go CLI](/01-getting-started/go-cli/) topic, with one lesson each for `go run`, `go build`, `go install`, `go fmt`, `go test` and `go get`, all practiced on the new [`cli-lab`](/assets/code/ch01/cli-lab/main.go) project.
* **Update**: Rewrote the introduction of [A Tour of the Toolchain](/01-getting-started/toolchain-tour.md) as the finale of the Go CLI topic, and added executable package, reusable package and workspace to the [Glossary](/06-reference/glossary.md).

## 2026-11-06
* **Creation**: Added [Import Statements](/01-getting-started/go-packages/import-statements.md) to the Go Packages topic: what `import "fmt"` gives you, exported vs unexported names, the four import forms, how import paths resolve, and where to find package documentation (golang.org/pkg, now pkg.go.dev/std, plus `go doc`). It comes with a new runnable example, [`ch01/imports`](/assets/code/ch01/imports/main.go).
* **Creation**: Published the first edition of the book: six chapters, runnable examples in [`assets/code`](/assets/code/go.mod), and diagrams in [`assets/images`](/assets/images/).
* **Creation**: Wrote [Chapter 1 · Getting Started](/01-getting-started/), including an [Environment Setup](/01-getting-started/environment-setup/) section on installing Go, installing VS Code, and configuring VS Code with the Go extension.
* **Update**: Added a step-by-step "Install the Go analysis tools from the Command Palette" section (`Go: Install/Update Tools`) to [Configuring VS Code for Go](/01-getting-started/environment-setup/configuring-vscode-go.md).
* **Creation**: Wrote [Chapter 2 · Language Basics](/02-language-basics/) and [Chapter 3 · Types and Abstraction](/03-types-and-abstraction/).
* **Creation**: Wrote [Chapter 4 · Concurrency](/04-concurrency/) and [Chapter 5 · Production Go](/05-production-go/).
* **Creation**: Wrote [Chapter 6 · Reference](/06-reference/): a cheatsheet, a gotchas page, and a glossary.
* **Initialization**: Set up the bundle as an Open Knowledge Format v0.2 bundle, with a root [index](/index.md) and this log.
