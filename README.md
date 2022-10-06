# Road to Kubernetes: Go Lang App Demo

This is a demo Go application for the Road to Kubernetes series. This repo is primarily for educational purposes but can be modified for production use.


----

## Deploy to Kubernetes

### Create Deployment and Service
Create a `deployment`:
```
kubectl apply -f https://raw.githubusercontent.com/roadtokubernetes/go.demo.roadtok8s.com/main/k8s/deployment.yaml --namespace=apps
``` 

Create a `service`:
```
kubectl apply -f https://raw.githubusercontent.com/roadtokubernetes/go.demo.roadtok8s.com/main/k8s/service.yaml --namespace=apps
```


Verify `deployment`:

```
kubectl get deployment go-demos-roadtok8s-com-deployment
```

Verify `service`:

```
kubectl get service go-demos-roadtok8s-com-service
```

Change `apps` to the namespace of your choosing. 



### Remove Deployment and Service
Delete the `deployment`:
```
kubectl delete -f https://raw.githubusercontent.com/roadtokubernetes/go.demo.roadtok8s.com/main/k8s/deployment.yaml --namespace=apps
```

Delete the `service`:
```
kubectl delete -f https://raw.githubusercontent.com/roadtokubernetes/go.demo.roadtok8s.com/main/k8s/service.yaml --namespace=apps
```

