  # minimal-go-nix-project

  Minimal Go starter with a Nix flake + direnv dev shell and
  GitHub Actions CI.

  ## Prerequisites

  - [Nix](https://nixos.org/download/) with flakes enabled
  - [direnv](https://direnv.net/)

  ## Setup

  ```bash
  direnv allow
  direnv exec . zsh
  # or without direnv
  nix develop --command zsh
  ```

  That loads the flake shell: Go, gopls, delve, golangci-lint,
  and buf.

  ## Rename for your project

  1. Change the module path in `go.mod` (`module example` →
  your path).
  2. Rename `cmd/example` to match your binary name.

  ## Develop

  After entering direnv/nix shell

  ```bash
  go run ./cmd/example
  go test ./...
  ```

  CI builds and tests on pushes to `main` and on pull
  requests.

