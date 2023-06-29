# Clusterflux Contributing Guide

Welcome to the contributing guide for the Clusterflux repository! We appreciate your interest in contributing to our project. By following these guidelines, you can help us maintain a collaborative and productive environment.

## Table of Contents

- [Clusterflux Contributing Guide](#clusterflux-contributing-guide)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Setting up the Development Environment](#setting-up-the-development-environment)
  - [Contributing](#contributing)
    - [Creating an Issue](#creating-an-issue)
      - [Bug Report](#bug-report)
      - [Enhancement Request](#enhancement-request)
      - [Feature Request](#feature-request)
    - [Submitting a Pull Request](#submitting-a-pull-request)

## Introduction

Welcome to the contributing guide for the ClusterFlux repository! We appreciate your interest in contributing to our project. ClusterFlux is a powerful tool designed to simplify the configuration and management of Kubernetes clusters using GitOps principles.

ClusterFlux builds upon the [Flux Cluster Template][flux_cluster_template_url] project and leverages the capabilities of GitOps to automate the deployment of Kubernetes applications and manage the lifecycle of Kubernetes resources. By using ClusterFlux, developers and operators can streamline their workflows, improve consistency, and ensure a reliable deployment process for their Kubernetes clusters.

In this guide, you'll find instructions on how to contribute to the ClusterFlux project, including creating and resolving issues, submitting pull requests, and adhering to our code guidelines. Your contributions can help enhance the functionality, improve usability, fix bugs, add new features, and address any issues that users may encounter.

We value collaboration, open communication, and a friendly environment for all contributors. Together, let's build an exceptional tool that empowers developers and operators in their Kubernetes journey.

Thank you for your interest in contributing to ClusterFlux. We appreciate your support and look forward to your contributions!

## Getting Started

### Prerequisites

Before you start contributing to the ClusterFlux repository, ensure that you have the following prerequisites in place:

- Go Programming Language
- Git
- Kubernetes Cluster
- kubectl
- Text Editor/IDE
  - I personally use *Visual Studio Code*

These prerequisites should be sufficient to get you started with contributing to the ClusterFlux repository. Make sure to refer to the repository's documentation or README file for any additional prerequisites or specific instructions related to the project.

### Setting up the Development Environment

To contribute to the ClusterFlux Go application and work with the code in Visual Studio Code (VS Code), follow these steps to configure your development environment:

1. **Go Programming Language**: Install Go on your system if you haven't already. You can download and install Go from the [official website][go_dl_url] and follow the installation instructions for your operating system. Make sure you have Go version 1.16 or above.

2. **Git**: Install Git on your system if you haven't already. You can download and install Git from the [official website][git_dl_url] and follow the installation instructions for your operating system.

3. **Visual Studio Code**: Download and install [Visual Studio Code][vscode_dl_url] for your operating system.

4. **Go Extension**: Launch VS Code and open the Extensions view by clicking on the square icon on the left sidebar or pressing `Ctrl+Shift+X` (`Cmd+Shift+X` on macOS). Search for the "Go" extension by Microsoft, and click the "Install" button to install it. This extension provides rich language support, code completion, and debugging capabilities for Go development.

5. **ClusterFlux Repository**: Clone the ClusterFlux repository to your local machine using the following command:

    ```shell
    git clone https://github.com/snoopy82481/clusterflux.git
    ```

6. **Open Project in VS Code**: Open VS Code and select "Open Folder" from the "File" menu. Navigate to the directory where you cloned the ClusterFlux repository and select it to open the project in VS Code.

7. **Go Tools**: VS Code will prompt you to install recommended Go tools and extensions when you open a Go project for the first time. Follow the instructions to install the required tools and extensions.

8. **Configure Go Tools in VS Code**: Open the VS Code settings by selecting "Preferences" → "Settings" or pressing `Ctrl+,` (`Cmd+,` on macOS). Search for "Go: Gopath" in the search bar, and set the appropriate Go workspace path to the ClusterFlux repository. This will ensure that VS Code uses the correct workspace for Go development.

9. **ClusterFlux Dependencies**: The ClusterFlux repository may have specific Go dependencies required to build and run the application. Typically, Go projects use a dependency management tool such as Go Modules. You can initialize Go Modules in the project's root directory by running the following command:

    ```shell
    go mod init github.com/snoopy82481/clusterflux
    ```

    This will initialize a new Go module and create a go.mod file in the project's root directory. To install the project's dependencies, you can use the go get command followed by the import path of the dependency.

10. **Build and Run**: Use the VS Code integrated terminal by selecting "View" → "Terminal" or pressing Ctrl+\` (Cmd+\` on macOS). In the terminal, navigate to the ClusterFlux repository's root directory and use the following command to build and run the ClusterFlux application:

    ```shell
    go run main.go
    ```

    Ensure that the application starts successfully without any errors.

Once you have set up the development environment, you can start exploring the ClusterFlux Go code, making changes, and contributing to the project using Visual Studio Code.

## Contributing

To contribute to the ClusterFlux repository, we require that every pull request is associated with a corresponding issue. This ensures proper tracking, discussion, and context for the changes being proposed.

### Creating an Issue

When creating an issue for the ClusterFlux repository, please follow the guidelines below to provide clear and actionable information. By following these guidelines, you help maintain an organized and efficient issue tracking process.

#### Bug Report

If you encounter a bug or unexpected behavior in the ClusterFlux application, use the bug report template to provide the necessary details. Follow these steps:

1. **Title**: Provide a descriptive title summarizing the issue concisely.
2. **Description**: Describe the issue in detail, including steps to reproduce it, the expected behavior, and the actual behavior you observed.
3. **Environment**: Mention the specific environment in which the bug occurred, including the operating system, Go version, and any other relevant details.
4. **Screenshots**: If applicable, include screenshots or error messages that help illustrate the issue.

#### Enhancement Request

If you have an idea to improve an existing feature or suggest a new capability for ClusterFlux, use the enhancement request template. Follow these steps:

1. **Title**: Provide a descriptive title summarizing the enhancement or new feature.
2. **Description**: Describe the enhancement or feature request in detail, including the problem it solves or the value it adds.
3. **Proposal**: Present a clear proposal for the enhancement, outlining the expected behavior and any implementation details if available.
4. **Alternatives**: If applicable, suggest alternative solutions or approaches for achieving the desired outcome.

#### Feature Request

If you have an idea for a new feature in the ClusterFlux application, use the feature request template. Follow these steps:

1. **Title**: Provide a descriptive title summarizing the new feature.
2. **Description**: Describe the feature request in detail, including its purpose, benefits, and any specific requirements or constraints.
3. **Use Cases**: Explain the use cases or scenarios in which this feature would be valuable.
4. **Implementation**: If you have any suggestions or ideas regarding the implementation of the feature, provide them in this section.

When creating an issue, make sure to search the existing issues to avoid duplicates. Provide as much relevant information as possible to help other contributors and maintainers understand and address the issue effectively.

### Submitting a Pull Request

1. **Create an Issue**: First, create a new issue in the repository by clicking on the "Issues" tab and then the "New Issue" button. Describe the problem you're addressing, the proposed changes, or the new feature you intend to add. If an issue already exists that corresponds to your contribution, you can proceed to the next step.
2. **Branch Naming Convention**: Create a new branch for your changes, using a naming convention that includes the issue number and a brief description of the changes. For example, if the issue number is 42 and you're adding a new feature, you can name your branch as `issue-42-new-feature`.
3. **Work on the Issue**: Make the necessary changes in your branch to address the issue or implement the proposed feature. Ensure that your changes adhere to the project's code guidelines and best practices.
4. **Commit and Reference the Issue**: As you make commits, please include a reference to the corresponding issue in your commit messages. This can be done by using the issue number followed by a brief description of the commit. For example, Fixes #42: Added validation for input fields.
5. **Create the Pull Request**: Once you have made the necessary changes and committed them with the appropriate references, it's time to create the pull request. Go to the repository's "Pull Requests" tab and click on the "New Pull Request" button. Select your branch as the "compare" branch and provide a descriptive title and detailed description of your changes.
6. **Link the Issue**: In the pull request description, link the corresponding issue by using the issue number and a brief description. For example, `This pull request addresses issue #42 by adding a new feature that improves user authentication`.

By following these guidelines, we can ensure that each pull request has an associated issue, facilitating better collaboration, tracking, and review processes.

[flux_cluster_template_url]: https://github.com/onedr0p/flux-cluster-template
[go_dl_url]: https://go.dev/dl/
[git_dl_url]: https://git-scm.com/downloads
[vscode_dl_url]: https://code.visualstudio.com/
