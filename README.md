# factorial

## Repository overview

### [`/cmd`](/cmd)
This directory contains main.go file, which contains the main function that serves as the starting point of the application.
### [`/pkg`](/pkg)
Package pkg provides a concurrent and sequential methods for calculating the factorial of a given number.
### [`/internal`](/internal)
It initializes the http server, configures routes, and starts the server.

## Tests

Tests are performed on 12-CPU machine.

| Factorial of | Sequential execution | Concurrent execution |
|--------------|----------------------|----------------------|
| 10000        | 10.6412ms            | 3ms                  |
| 100000       | 1.2393814s           | 100ms                |
| 1000000      | 157s                 | 6s                   |