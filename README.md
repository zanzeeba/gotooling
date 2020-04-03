# gotooling
this is the Linux Academy course System Tooling with Go

It implements a basic go application that asks for two items of input from the command line, lets you se what is required with a 'prompt' boolean and lets you see more details with a 'debug' bool.

it then writes this data to a file. In the original file it wrote the data to /etc/motd which meant altering several elements of the go path so i avoided that by writing to the current directory.

## Input Values

To prompt for the input values this version uses the Go flags

## cobra

cobra is a package for writing more complex and more manageable user cli inputs and ther is a version of this project from the same source at gotoolingcobra in my repos
