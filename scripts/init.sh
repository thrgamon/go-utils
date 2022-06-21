#!/bin/bash
set -Eeuo pipefail

current_dir=${PWD##*/}

echo "Setting up go application to be deployed to dokku"
echo "What is the name of the app? ${current_dir}"
read app_name

if [ -z "$app_name" ]
then
  app_name=${current_dir}
fi

echo "Setting \$APP_NAME to ${app_name}"
echo "APP_NAME=${app_name}" >> .env

echo "Setting \$DB_NAME to ${app_name}_db"
echo "DB_NAME=${app_name}_db" >> .env

if [ -z "$DOMAIN_NAME" ]
then
  echo "What is the domain?"
  read domain

  echo "Setting \$DOMAIN_NAME to ${domain}"
  echo "DOMAIN_NAME=${domain}" >> .env
else
  echo "DOMAIN_NAME is already set: $DOMAIN_NAME"
fi

if [ -z "$DOKKU_REMOTE" ]
then
  echo "What is the domain?"
  read dokku_remote

  echo "Setting \$DOKKU_REMOTE to ${dokku_remote}"
  echo "DOKKU_REMOTE=${dokku_remote}" >> .env
else
  echo "DOKKU_REMOTE is already set: $DOKKU_REMOTE"
fi
