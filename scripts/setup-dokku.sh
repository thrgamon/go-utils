#!/bin/bash
set -Eeuo pipefail

ssh dokku apps:create ${APP_NAME}
git remote add dokku dokku@${DOKKU_REMOTE}:${APP_NAME}
ssh dokku domains:add ${APP_NAME} ${APP_NAME}.${DOMAIN_NAME}
ssh dokku letsencrypt:enable ${APP_NAME}
