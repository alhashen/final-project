#!/bin/sh

url=http://localhost:2400/photos
valid_id=22
invalid_id=100
dvalid_id=21

echo -e "Deletes valid photo"
curl -s -X DELETE $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nDeletes invalid photo"
curl -s -X DELETE $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nDeletes unauthorized photo"
curl -s -X DELETE $url/$dvalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  | json_pp
