#!/bin/bash

dirs=$(find .files/src -type d -print0 | xargs -0)

go-bindata -prefix .files/src -pkg proxyfs -o embedded_test.go $dirs
