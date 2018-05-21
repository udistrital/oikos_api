#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export OIKOS_API_DB_USER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/oikos_api/db/username --output text --query Parameter.Value)"
  export OIKOS_API_DB_PASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/oikos_api/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
