package main

import "fmt"
import "rsc.io/quote"

func main() {
  fmt.Println(quote.Go())
}

// go mod init example.com/greetings

// The module path is typically of the following form:
// 
// <prefix>/<descriptive-text>
// The prefix is typically a string that partially describes the module, such as a string that describes its origin. This might be:
// 
// The location of the repository where Go tools can find the module’s source code (required if you’re publishing the module).
// 
// For example, it might be github.com/<project-name>/.
// 
// Use this best practice if you think you might publish the module for others to use. For more about publishing, see Developing and publishing modules.
// 
// A name you control.
// 
// If you’re not using a repository name, be sure to choose a prefix that you’re confident won’t be used by others. A good choice is your company’s name. Avoid common terms such as widgets, utilities, or app.
// 
// For the descriptive text, a good choice would be a project name. Remember that package names carry most of the weight of describing functionality. The module path creates a namespace for those package names.
