# Overview

This package, `go_cli_scripts`, offers a collection of utility functions written in Go, designed to streamline various common operations such as array manipulation, asynchronous processing, directory navigation, file handling, and more. Here's a brief rundown of its capabilities:

- **Array Utilities**: Functions to remove elements, check for existence, filter, and convert between types.
- **Asynchronous Utilities**: A method to manage asynchronous tasks in batches.
- **Directory Utilities**: Functions to navigate and manipulate filesystem directories.
- **File Utilities**: A suite of functions for file copying, reading, writing, and more.
- **Docker Utilities**: Determine if the current environment is within a Docker container.
- **Git Utilities**: Facilitate operations like sparse cloning of repositories.
- **GitHub Utilities**: Interact with GitHub to fetch latest release information and assets.
- **JSON Utilities**: Assist in filtering and formatting JSON data.
- **Map Utilities**: Provide methods for map manipulation and conversion.
- **Menu Display Utilities**: Facilitate the creation of interactive menus in the terminal.
- **Standard I/O Utilities**: Streamline the process of capturing input and executing commands.
- **String Utilities**: Offer various string manipulation capabilities.
- **TestNG Utilities**: Assist in configuring and executing TestNG tests.
- **VSCode Settings Utilities**: Facilitate the retrieval and handling of VSCode settings.

This package is structured to support a wide range of applications, from simple file operations to more complex tasks like handling asynchronous processes or interacting with external APIs.


# Installation

To integrate the `go_cli_scripts` package into your Go project, follow these steps:

1. Ensure you have Go installed on your system. You can download and install Go from [the official Go website](https://golang.org/dl/).

2. Navigate to your project directory where you want to include the `go_cli_scripts` package.

3. Use the `go get` command to fetch the package from its repository. Since this is a local package, you would typically copy it into your project's directory or reference it locally. If it were hosted, you would use a command like:

   ```sh
   go get github.com/yourusername/go_cli_scripts
   ```

   Adjust the command based on the actual location of the `go_cli_scripts` package.

4. Import the package in your Go files where you need to use its functionality:

   ```go
   import (
       "path/to/go_cli_scripts/utils"
   )
   ```

   Replace `path/to/go_cli_scripts` with the actual path to the `go_cli_scripts` package in your project.

5. Once the package is imported, you can use its functions by referencing them with the `utils` prefix, like `utils.ArrayContainsAny(...)`.

By following these steps, you can leverage the various utility functions provided by the `go_cli_scripts` package in your Go projects to enhance your applications' capabilities and streamline your development process.
