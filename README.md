# TungstenFabric Operator
## Quick Start
### Create Operator
```
kubectl apply -f https://raw.githubusercontent.com/michaelhenkel/tungstenfabric-operator/v0.0.3/deploy/1-create-operator.yaml
```
### Start Cluster
#### 1 Node
```
kubectl apply -f https://github.com/michaelhenkel/tungstenfabric-operator/blob/v0.0.3/deploy/2-start-operator-1node.yaml
```
#### 3 Node
```
kubectl apply -f https://github.com/michaelhenkel/tungstenfabric-operator/blob/v0.0.3/deploy/2-start-operator-3node.yaml
```
## Customize configuration
```
curl -o mycustomtf.yaml https://github.com/michaelhenkel/tungstenfabric-operator/
blob/v0.0.3/deploy/2-start-operator-1node.yaml
```
Customize settings in mycustomtf.yaml and apply it
```
kubectl apply -f mycustomtf.yaml
```

## Check resources

