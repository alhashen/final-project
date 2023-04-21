#!/bin/sh

url=http://localhost:2400/social_medias
valid_id=17
invalid_id=1000

echo -e "Social media before being updated"
curl -s -X GET $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  | json_pp

echo -e "\nUpdate valid social media"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"name":"My Name", "social_media_url": "http://social_media.com"}' \
  | json_pp

echo -e "\nUpdate social media with invalid ID"
curl -s -X PUT $url/$invalid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"name":"My Name", "social_media_url": "http://social_media.com"}' \
  | json_pp

echo -e "\nUpdate invalid social media value"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTYsInVzZXJfZW1haWwiOiJteV9lbWFpbEBhZGRyZXNzLmNvbSJ9.6kfn1Q3IUbGWY71TMoPDkPA8FKIJvYMHxII01QYjGPY' \
  -d '{"name":"", "social_media_url": ""}' \
  | json_pp 

echo -e "\nUpdate unauthorized social media"
curl -s -X PUT $url/$valid_id\
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTEsInVzZXJfZW1haWwiOiJlbWFpbEBhZGRyZXNzLmNvbSJ9.IePgtiBIOCPOJCiZlut25WUGPehkW2yToB36ZuI5q-U' \
  -d '{"name":"My Name", "social_media_url": "http://social_media.com"}' \
  | json_pp 





