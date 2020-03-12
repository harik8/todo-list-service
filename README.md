# todo-list-service

Todo List Service is a GO API service which serves Todo data.

# Requests 

- Add a Todo task
```
curl -X POST -H "Content-Type: application/json" -d '{                                                                                           
  "task": "Fix bug 130320",
  "duedate": "2020-03-20",
  "labels": "Dev",
  "comments": "In progress"
}' http://127.0.0.1:8080/todo/
```

- Update a Todo task
```
curl -X PUT -H "Content-Type: application/json" -d '{
  "task": "Fix bug 130320",
  "duedate": "2020-03-20",
  "labels": "QA",
  "comments": "Verification"
}' http://127.0.0.1:8080/todo/{TID}
```

- Read all Todo tasks
```
curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/
```

- Read a specific Todo task
```
curl -X GET -H "Content-Type: application/json"     http://127.0.0.1:8080/todo/{TID}
```

- Delete a specific Todo task
```
curl -X DELETE -H "Content-Type: application/json"  http://127.0.0.1:8080/todo/{TID}
```

# Docker

`docker pull harik8/todo:v1.0.0`

# Local Run

- Install [kind](https://github.com/kubernetes-sigs/kind)

- ```kind create cluster```
- ```kubectl create ns mongodb```
- ```kubectl create ns todo```
- ```helm install mongodb stable/mongodb -n mongodb```
- ```export MONGODB_ROOT_PASSWORD=$$(kubectl get secret --namespace mongodb mongodb -o jsonpath="{.data.mongodb-root-password}" | base64 --decode)```
- ```kubectl create secret generic mongo-cred --from-literal=MONGO_PASSWORD=${MONGODB_ROOT_PASSWORD} -n todo```
- ```git clone https://github.com/harik8/sg-kube-deployer.git```
- ```cd sg-kube-deployer```
- ```kubectl apply -f manifests/todo/```
- ```kubectl port-forward po/<todo_pod_name> -n todo 8080:8080```