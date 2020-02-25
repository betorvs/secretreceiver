#!/usr/bin/env bash
# MAC OS tested only


METHOD="$1"
LOCALSERVER="http://localhost:9090/secretreceiver/v1/secret"

case $METHOD in
  "POST")

  NAME="$2"
  VALUE="$3"
  NAMESPACE="$4"

  CHECKSUM=$( echo -n "$VALUE" | shasum -a 512 | awk '{print $1}')

  curl -v --header "Content-Type: application/json" \
    --request POST \
    --data "{\"name\":\"${NAME}\",\"namespace\":\"${NAMESPACE}\", \"checksum\": \"${CHECKSUM}\", \"data\": { \"token\": \"${VALUE}\" } }" \
    ${LOCALSERVER}

;; 
"GET")
  NAME="$2"
  NAMESPACE="$3"
  curl -v --header "Content-Type: application/json" "${LOCALSERVER}/${NAMESPACE}/${NAME}"


;;
"PUT")

  NAME="$2"
  VALUE="$3"
  NAMESPACE="$4"

  CHECKSUM=$( echo -n "$VALUE" | shasum -a 512 | awk '{print $1}')

  curl -v --header "Content-Type: application/json" \
    --request PUT \
    --data "{\"name\":\"${NAME}\",\"namespace\":\"${NAMESPACE}\", \"checksum\": \"${CHECKSUM}\", \"data\": { \"token\": \"${VALUE}\" } }" \
    ${LOCALSERVER}

;; 
"DELETE")
  NAME="$2"
  NAMESPACE="$3"
  curl -v --request DELETE --header "Content-Type: application/json" "${LOCALSERVER}/${NAMESPACE}/${NAME}"


;;
"CHECK")
  kubectx docker-desktop
  kubectl get secret -n ${NAMESPACE}
  ;;
esac