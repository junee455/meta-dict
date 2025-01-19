#!/usr/bin/bash
docker compose --env-file ./local.env -f ./compose.yaml up --build
