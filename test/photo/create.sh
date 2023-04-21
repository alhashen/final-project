#!/bin/sh

url=http://localhost:2400/photos

echo -e "Unauthorized photo upload"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title":"Lonely", "photo_caption":"lonely, yo!", "photo_url": "http://url.com"}' \
  | json_pp 

echo -e "\nCreating valid photo"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title":"My Pic 5", "photo_caption":"okay", "photo_url": "http://image-url.com"}' \
  | json_pp 

echo -e "\nCreating invalid photo"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"photo_title": ""}' \
  | json_pp 

