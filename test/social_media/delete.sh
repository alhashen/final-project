#!/bin/sh

url=http://localhost:2400/social_medias
valid_id=15
invalid_id=100
dvalid_id=17

echo -e "Deletes valid social media"
curl -s -X DELETE $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nDeletes invalid social media"
curl -s -X DELETE $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nDeletes unauthorized social_media"
curl -s -X DELETE $url/$dvalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp
