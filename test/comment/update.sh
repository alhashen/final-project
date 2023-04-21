#!/bin/sh

url=http://localhost:2400/comments
valid_id=9
invalid_id=1000

echo -e "Comment before being updated"
curl -s -X GET $url/$valid_id\
  -H 'Content-type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  | json_pp

echo -e "\nUpdate valid comment"
curl -s -X PUT $url/$valid_id\
  -H 'Content-type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":"Nice!"}' \
  | json_pp

echo -e "\nUpdate comment with invalid id"
curl -s -X PUT $url/$invalid_id\
  -H 'Content-type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":"Nicee!"}' \
  | json_pp

echo -e "\nUpdate comment with invalid comment value"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":""}' \
  | json_pp 

echo -e "\nUpdate unauthorized comment"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"comment_message":"Nice!"}' \
  | json_pp 





