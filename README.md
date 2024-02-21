# Design Pattern Implementations in Go

Welcome to my collection of design pattern implementations in Go! This repository contains various design patterns that I've implemented for learning and practice.

## Table of Contents

1. [Introduction](#introduction)
2. [How to Use](#how-to-use)
3. [Design Patterns](#design-patterns)
    - [Factory Method](#factory-method)
    - [Abstract Factory](#abstract-factory)
4. [Contributing](#contributing)
5. [License](#license)

## Introduction

Design patterns are essential in software development as they provide proven solutions to recurring design problems. This repository aims to showcase various design patterns implemented in Go.

## How to Use

To use any of the design patterns in your project, follow these general steps:

1. Clone this repository:

    ```bash
    git clone https://github.com/rafaelcmd/design-patterns-go.git
    ```

2. Navigate to the specific pattern directory.

3. Import the necessary files into your project or use the provided code directly.

## Design Patterns

### Factory Method

1. [Factory Method Implementation](factory-method/main.go)

The Factory Method pattern is a creational pattern that provides an interface for creating objects in a superclass, but allows subclasses to alter the type of objects that will be created. This pattern is useful when a class cannot anticipate the type of objects it needs to create.

### Abstract Factory

1. [Abstract Factory Implementation](abstract-factory/main.go)

The Abstract Factory pattern is a creational pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes. This pattern is useful when a system needs to be independent of how its objects are created, composed, and represented.

## Contributing

Contributions are welcome! If you have additional design patterns to add, find bugs, or want to improve existing implementations, please feel free to open a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
