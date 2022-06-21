#!/bin/bash
set -Eeuo pipefail

ssh dokku postgres:create ${DB_NAME}
ssh dokku postgres:link ${DB_NAME} ${APP_NAME}
ssh dokku postgres:expose ${DB_NAME}
shh dokku postgres:info ${DB_NAME}
