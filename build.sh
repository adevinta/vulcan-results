#!/bin/bash

set -e

# Ensure the current version of gogen is installed

go install github.com/goadesign/goa/goagen


# Autogenerate content
goagen bootstrap -d github.com/adevinta/vulcan-results/design

# Compile and install
go install ./...
