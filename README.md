# Learn Go

## Available Build tags
- `playingcard`
- `tradingcard`
- `feature`


## Usage

Based on the build tags, it will either run the generics
or non-generics version.

```sh
# Build
make build
```

```sh
# Runs the build binaries (this will build again)
make run # Default
```

```sh
# Run the dev 
make dev
```

```sh
# Clean binaries
make clean
```

Automatically uses either `*Generics` version 
or the non-generics version depending on the Go version.

- Support added upto go2 version