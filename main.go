package main

import (
	setnamespace "github.com/GoogleContainerTools/kpt-functions-catalog/functions/go/set-namespace/transformer"
	"github.com/GoogleContainerTools/kpt-functions-sdk/go/fn"
	"sigs.k8s.io/kustomize/kyaml/filesys"
	"sigs.k8s.io/kustomize/kyaml/kio"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func main() {
	rfr := kio.LocalPackageReadWriter{
		PackagePath: "manifests",
		FileSystem:  filesys.FileSystemOrOnDisk{FileSystem: filesys.MakeFsOnDisk()},
	}

	var rl fn.ResourceList

	nodes, _ := rfr.Read()
	for _, node := range nodes {
		kubeObject, _ := fn.ParseKubeObject([]byte(node.MustString()))
		rl.Items = append(rl.Items, kubeObject)
	}

	rl.FunctionConfig, _ = fn.ParseKubeObject([]byte(`
apiVersion: v1
kind: ConfigMap
metadata:
  name: kptfile.kpt.dev
data:
  name: newNamespace
`))

	setnamespace.Run(&rl)

	var newNodes []*yaml.RNode
	for _, obj := range rl.Items {
		newNodes = append(newNodes, yaml.MustParse(obj.String()))
	}

	rfr.Write(newNodes)
}
