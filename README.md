# todo-list-service

curl -X POST -H "Content-Type: application/json" -d '{                                                                                           
  "task": "develop go list service",
  "duedate": "2020-03-20",
  "labels": "Dev",
  "comments": "Initiated"
}' http://127.0.0.1:8080/todo/

curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/
curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/{TID}
curl -X DELETE -H "Content-Type: application/json"  http://127.0.0.1:8080/todo/{TID}

curl -X PUT -H "Content-Type: application/json" -d '{
  "task": "develop go list service",
  "duedate": "2020-03-25",
  "labels": "Dev",
  "comments": "Initiated"
}' http://127.0.0.1:8080/todo/5e66104b8bacfbc0da3a7165