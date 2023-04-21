#!/bin/sh

url=http://localhost:2400/users

echo "Register valid user"
curl -s -X POST $url/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"dbtest12", "user_email":"another_email@address.com", "user_password":"pass123", "user_age":15}' \
  | json_pp

echo -e "\nRegister duplicate user data"
curl -s -X POST $url/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"user1232", "user_email":"email@address.com", "user_password":"pass123", "user_age":15}' \
  | json_pp

echo -e "\nRegister invalid user"
curl -s -X POST $url/register \
  -H 'Content-Type: application/json' \
  -d '{"username":"user1232", "user_email":"email-address.com", "user_password":"pass123", "user_age":10}' \
  | json_pp
