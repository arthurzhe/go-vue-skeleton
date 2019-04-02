# go-vue-skeleton

## background

This project serves to be a skeleton for a SPA that utilises Vue.js
for the FE and a Go API server underneath.

## building

``go-bindata`` is required to load binary data into the Go context,
allowing for files to be deflated at runtime. To install this tool,
a make target is provided.

    make bin-prep

To compile the frontend components ready for binary compilation,
another make target is available.

    make build-fe

To build the bindata requirements, separate make targets are
utilised.

    make bin-migrations bin-dist

Once all bindata requirements are available, a series of binaries
can be constructed from 1 makefile command.

    make binaries

A series of binaries will then be available within the
``build/`` directory.

To build a local binary to run, a base target is available.

    make install

## development

Running tests will require the frontend to be built, and bindata
to be made available. These are captured in appropriate make
targets, and summed up into a single test.

    make test

For streamlined development, the frontend can be setup to make use
of a watch task.

    make dev-fe

A base target is available to compile the Go binary and execute it.

    make dev-be

