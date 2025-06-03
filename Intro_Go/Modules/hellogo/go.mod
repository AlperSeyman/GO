module github.com/AlperSeyman/hellogo

go 1.24.2

// This works for local-only development, bad.
replace github.com/AlperSeyman/mystrings v0.0.0 => ../mystrings

require (
    github.com/AlperSeyman/mystrings v0.0.0
) 