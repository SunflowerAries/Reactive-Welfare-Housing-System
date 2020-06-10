#!/bin/sh
for i in $(ls $RWHS_HOME/json | grep '.*\.json');
do
echo $i
done

echo "Add new houses"
curl -X POST -H "Content-Type: application/json" -d @$RWHS_HOME/json/houses\.json http://localhost:3000/houses

echo "Check houses"
curl -X POST -H "Content-Type: application/json" -d @$RWHS_HOME/json/check\.json http://localhost:3000/houses/check

echo "Tenants checkout houses"
curl -X POST -H "Content-Type: application/json" -d @$RWHS_HOME/json/checkout\.json http://localhost:4000/users/checkout

echo "Tenants apply new houses"
curl -X POST -H "Content-Type: application/json" -d @$RWHS_HOME/json/request\.json http://localhost:4000/users
