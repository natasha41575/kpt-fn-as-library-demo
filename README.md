# kpt-fn-as-library-demo

This is a demo of how to use the kpt set-namespace function as a library. To run the demo:

```sh 
# clone the repo
git clone https://github.com/natasha41575/kpt-fn-as-library-demo.git

# run the demo
cd kpt-fn-as-library-demo; go run main.go

# view the diff
git diff
```

You should observe that the namespaces under the `manifests` directory have been set correctly,
including special handling for the ClusterRoleBinding and Namespace objects.