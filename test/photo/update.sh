#!/bin/sh

url=http://localhost:2400/photos
valid_id=20
invalid_id=1000

echo -e "Photo before being updated"
curl -s -X GET $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  | json_pp

echo -e "\nUpdate valid photo"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title":"Nice sunset!", "photo_caption":"with friends!", "photo_url":"http://imgurl.com"}' \
  | json_pp

echo -e "\nUpdate photo with invalid ID"
curl -s -X PUT $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title":"Not Lonely", "photo_caption":"Not lonely, yo!"}' \
  | json_pp

echo -e "\nUpdate invalid photo value"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title":"i", "photo_url":"a"}' \
  | json_pp 

echo -e "\nUpdate unauthorized photo"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"photo_title":"Helllo", "photo_caption":"Not lonely, yo!"}' \
  | json_pp 





