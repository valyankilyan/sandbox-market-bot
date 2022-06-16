#!/bin/bash

docker run \
--name mypostgres \
-p 5432:5432 \
-d mypostgres

# -v /data:/var/lib/postgresql/data