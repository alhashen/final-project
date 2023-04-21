#!/bin/sh

url=http://localhost:2400/users

echo -e "Log In with valid user data"
curl -s -X POST $url/login \
  -H 'Content-Type: application/json' \
  -d '{"user_email":"another_email@address.com", "user_password":"pass123"}' \
  | json_pp

echo -e "\nLog In with false user data"
curl -s -X POST $url/login \
  -H 'Content-Type: application/json' \
  -d '{"user_email":"my_email1232@address.com", "user_password":"pass123"}' \
  | json_pp

