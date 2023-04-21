#!/bin/sh

url=http://localhost:2400/social_medias

echo -e "Create valid social media"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"name":"Protagonist", "social_media_url": "http://social_media.com"}' \
  | json_pp 

echo -e "\nCreate invalid social media"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"name": "", "social_media_url": ""}' \
  | json_pp 

echo -e "\nCreate duplicate social media"
curl -s -X POST $url/\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"Name":"Protagonist", "social_media_url": "http://social_media.com"}' \
  | json_pp 

