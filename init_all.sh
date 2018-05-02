#!/bin/sh

if [ "$INITALL" = "true" ]
then
   echo "INITALL is true"
   resetShyftGeth_docker.sh && initShyftGeth_docker.sh && startShyftGeth_docker.sh
else
  if [ "$FIRSTINIT" = "true" ]
  then
     echo "FIRSTINIT is true"
     initShyftGeth_docker.sh && startShyftGeth_docker.sh
  else
     echo "FIRSTINIT is false"
     startShyftGeth_docker.sh
  fi
fi