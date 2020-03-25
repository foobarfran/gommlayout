# Go Multi Module Project Layout

## How to use it
Configure your libraries and apps modules in the `APP_MODULES` and `LIB_MODULES` vars in the Makefile 

## Make basic rules 
  - deps: executes go mod tidy for all modules
  - test: tests all modules
  - build: build all modules, generating the apps binaries in the `target` folder

Each rule has the previous one as a prerequisite. For example, if you call `make build`, `deps` and `test` are executed prior to the build.  

## Make additional rules
 - clean: rule that just deletes the `target` folder
 - release: creates the git tags for all modules. A flag `RELEASE_VERSION` in the way `1.2.3` is required

### Some Examples

```
# Build all modules
make clean build

# Release a new version
make release RELEASE_VERSION=0.0.1
```
