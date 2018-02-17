#!/bin/bash

while true; do
  docker stats --no-stream --no-trunc --format '{{json .}}'
  sleep 10
done
