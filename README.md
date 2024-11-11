# Go-Git

This library provides functionalities to clone and pull Git repositories using the `go-git` package.

It supports setting authentication, branch selection, depth cloning, progress reporting, and destination directory customization.

## Installation

To install this library, you can use the following command:

```sh
go get github.com/weidongkl/git
```

## Usage

### Cloning a Repository

1. **Create a new clone instance**:
    ```go
    clone := NewClone("https://github.com/user/repo.git")
    ```

2. **Set optional parameters** (authentication, branch, depth, progress, destination directory):
    ```go
    clone.SetAuth("username", "password").
           SetBranch("main").
           SetDepth(1).
           SetProgress(os.Stdout).
           SetDstDir("/path/to/destination")
    ```

3. **Clone the repository**:
    ```go
    err := clone.Clone()
    if err != nil {
        log.Fatalf("Failed to clone repository: %v", err)
    }
    ```

### Pulling Changes from a Repository

1. **Create a new pull instance**:
    ```go
    pull := NewPull("/path/to/local/repo")
    ```

2. **Set optional parameters** (authentication, progress):
    ```go
    pull.SetAuth("username", "password").
         SetProgress(os.Stdout)
    ```

3. **Pull changes**:
    ```go
    err := pull.Pull()
    if err != nil {
        log.Fatalf("Failed to pull changes: %v", err)
    }
    ```

## Example Code

Here is an example of how to use the library to clone and pull a repository:

```go
package main

import (
	"log"
	"os"
	"github.com/weidongkl/git"
)

func main() {
	// Clone a repository
	clone := git.NewClone("https://github.com/user/repo.git").
		SetAuth("username", "password").
		SetBranch("main").
		SetDepth(1).
		SetProgress(os.Stdout).
		SetDstDir("/path/to/destination")
	err := clone.Clone()
	if err != nil {
		log.Fatalf("Failed to clone repository: %v", err)
	}

	// Pull changes from a repository
	pull := git.NewPull("/path/to/local/repo").
		SetAuth("username", "password").
		SetProgress(os.Stdout)
	err = pull.Pull()
	if err != nil {
		log.Fatalf("Failed to pull changes: %v", err)
	}
}
```

## Contributing

Contributions are welcome! Please fork the repository and create a pull request with your changes. Make sure to follow the existing code style and add tests for your new features or fixes.

## License

This project is licensed under the MIT License.

## Contact

For any questions or issues, please contact weidongkx@gmail.com