# go-sample-module

Shows how to put a Go module on github and use it in another module e.g. the test harness  
This repo contains both the module and the test harness:

- cfssamplegomodule - the module code
- cfssamplegomodule-test-harness - the test harness, which references the above module


## TL;DR - How it works

- The module is in directory: "cfssamplegomodule"
- The module's go.mod includes the name of the module as:
```
module github.com/connorsadler/go-sample-module/cfssamplegomodule
```
- That module name will be used to find the module on the internet.  
You can point your browser at:
```
http://github.com/connorsadler/go-sample-module/
```
and look in the "cfssamplegomodule" for the module code and other files.

- In a project when wants to use this module, you can use:
```
go get github.com/connorsadler/go-sample-module/cfssamplegomodule
```

See the test harness for an example of such a project.  
All pretty handy!


## Test Harness - worked example:

Here are some steps to follow to use the test harness, to run code from the "cfssamplegomodule" module.

1. Clone the repo

2. checkout the "start_here" branch
```
git checkout start_here
```

3. cd to the cfssamplegomodule-test-harness directory

```
    cd cfssamplegomodule-test-harness
```

4. Try to run the test harness with:

```    
    go run .
```

It should fail with a message like this:
```
main.go:6:2: no required module provides package github.com/connorsadler/go-sample-module/cfssamplegomodule; to add it:
    go get github.com/connorsadler/go-sample-module/cfssamplegomodule
```

5. Get the dependency with the following command:
```
go get github.com/connorsadler/go-sample-module/cfssamplegomodule
```

You should see something like this:

```
go: downloading github.com/connorsadler/go-sample-module v0.0.0-20241003124043-d99ce668fcad
go: downloading github.com/connorsadler/go-sample-module/cfssamplegomodule v0.0.0-20241003124043-d99ce668fcad
go: added github.com/connorsadler/go-sample-module/cfssamplegomodule v0.0.0-20241003124043-d99ce668fcad
```

Observe that this command has now changed some files:
- the test harness go.mod to include a 'require'
- the test harness go.sum to include references to the module found  

Note: The module references include the git hash of the 'main' branch of the repo

6. Run main.go again with:
```
    go run .
```

This time it should print something like this:
```
cfssamplegomodule-test-harness - main running
resultFromModule: This is from SampleGoModuleFunction [main] at 20241001_155610
cfssamplegomodule-test-harness - main done
```


## Test Harness - further example:

TODO: Show a further example where we reference a different BRANCH of the module

## References

- https://go.dev/doc/modules/managing-source
- https://go.dev/ref/mod#go-get


## Gotchas:

### Gotcha 1:

go get caches modules to speed things up, but a lot of times you don't want this when developing - you want to force it to 
get the **latest version** from github, for example.  
To achieve this, you may need these commands to stop the 'go' command caching things:

- go clean -modcache
- export GOPROXY="direct"

I just run both commands in a terminal and enjoy cache-free goodness.  
The "go clean" command you may need to run more than once.

### Gotcha 2:

Error message "invalid version: unknown revision"  
You may see an error like this:

```
go: downloading github.com/connorsadler/go-sample-module/cfssamplegomodule v0.0.0-20241001163936-f26e55ee2fa1
go: cfssamplegomodule-test-harness imports
github.com/connorsadler/go-sample-module/cfssamplegomodule: github.com/connorsadler/go-sample-module/cfssamplegomodule@v0.0.0-20241001163936-f26e55ee2fa1: invalid version: unknown revision f26e55ee2fa1
```

**Cause:**  

go.mod and go.sum include the git commit hash of the module, and that hash is probably no longer valid. This happens if I rebase main to tidy it up!  

**Solution:**  

- remove the "require" from go.mod
- remove the 2 lines from go.sum which related to module: github.com/connorsadler/go-sample-module/cfssamplegomodule
- run both commands in "Gotcha 1" to ensure go isn't caching anything
- rerun the "go get" command:
```
go get github.com/connorsadler/go-sample-module/cfssamplegomodule
```

