#!/bin/bash

RUN_PORT=${PORT:-8030}

PORT=$RUN_PORT /usr/local/go/bin/go run server.go