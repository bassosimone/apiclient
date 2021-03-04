# PoC OONI client API

This repository contains a PoC where we automatically generate
Go clients for using the OONI API.

Most code is generated from annotated data structures and from
definitions of all the supported APIs.

We generate a Swagger 2.0 definition of the API recognized by
the client. Tests check whether the client's definition of the
API is in sync with the server's own definition.

This code will eventually be merged into ooni/probe-cli.

**Update**: merged in probe-cli here: https://github.com/ooni/probe-cli/pull/234
