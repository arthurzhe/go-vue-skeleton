# go-vue-skeleton

## background

This project serves to be a skeleton for a SPA that utilises Vue.js
for the FE and a Go API server underneath.

## layout

Several base concepts are at play in this skeleton, prepared with
each of the accompanying commands to allow for a straight-to-dev
process.

When first using this skeleton, a full rename and text replace of
``go-vue-skeleton`` to ``your-project-name`` kick-starts the process.
With this completed, the binary that will inevitably be built will
originate from the **main** package in ``./your-project-name``.

The ``data`` directory contains the database interactions, including
any migrations for managing the structure.

Vue components and structure files sit within ``src``, including all
other javascript. In this skeleton case, ``api.js`` is a file we use
to make calls to our API server, with remaining files being a part
of the Vue bootstrapping and usage process. To join with this, the
``public/index.html`` file is used when building assets.

In the root directory then, dependency and project configuration
files are found. This means that all commands necessary to build
and manage application components are run from the root directory.

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

## notes

For this skeleton project, there are no containers in use. This
means that locally there are a few dependencies to be installed.

1. Go
2. Node.js (with Yarn)
3. SQLite (assuming the database is kept)

Of course, the database can be swapped out with library references
replaced within the code.

Intended outcome from this project is a single binary, preventing
the need to have many dependent components installed on the target
host. Using containers would be recommended for teams wishing to
take this skeleton and collaborate. Many headaches are alleviated
by doing so (or use linux as your development machine OS).

During development, when running ``yarn serve`` alongside the Go
binary, the ``./src/api.js`` axios client will require the
``baseURL:`` value (as they will run on separate ports). Remove
this ``baseURL:`` when releasing this or keep it and update if
there will only be a single, known, domain used to serve from.
The default value for this will be the current hostname
including the port.
