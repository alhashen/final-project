#!/bin/sh

url=http://localhost:2400/comments/photo
valid_id=20
invalid_id=100

echo -e "Adding valid comment"
curl -s -X POST $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":"LOL, okay"}' \
  | json_pp

echo -e "\nAdding comment on invalid photo"
curl -s -X POST $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":"LOL"}' \
  | json_pp

echo -e "\nAdding invalid comment value"
curl -s -X POST $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"comment_message":""}' \
  | json_pp

