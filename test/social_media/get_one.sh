#!/bin/sh

url=http://localhost:2400/social_medias
valid_id=17
invalid_id=100

echo -e "Getting valid social media"
curl -s -X GET $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nGetting invalid social media"
curl -s -X GET $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp


