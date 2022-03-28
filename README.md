# Elastic GPU

## How to generate apis

```
$ ./hack/update-codegen.sh // codes
$ make manifests // crd
```

## How to use

1. Create EGPU CRD in Kubernetes cluster:

```
$ kubectl apply -f config/crd/bases/elasticgpu.io_elasticgpus.yaml
```

2. Referenced in your codes as follows:

```
package test

import (
	"context"
	"fmt"
	"strings"

	"elasticgpu.io/elastic-gpu/api/elasticgpu.io/v1alpha1"
	"elasticgpu.io/elastic-gpu/clientset/versioned"
)

func main() {
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", m.kubeconf)
	if err != nil {
		return nil, err
	}
	eGPUClient, err := versioned.NewForConfig(kubeconfig)
	if err != nil {
		return panic(err)
	}
	egpu := &v1alpha1.ElasticGPU{}
	egpu.Name = fmt.Sprintf("qgpu-%s", strings.ToLower(info.Devices.Hash()))
	egpu.Spec.Capacity = v1.ResourceList{
		v1alpha1.ResourceQGPUCore:   resource.MustParse(fmt.Sprintf("%d", core)),
		v1alpha1.ResourceQGPUMemory: resource.MustParse(fmt.Sprintf("%d", memory)),
	}
	egpu.Spec.ElasticGPUSource = v1alpha1.ElasticGPUSource{
		QGPU: &v1alpha1.QGPUElasticGPUSource{
			BaseGPUSource: v1alpha1.BaseGPUSource{
				Index: fmt.Sprintf("%d", qgpuIdx),
			},
		},
	}
	egpu.Status = v1alpha1.ElasticGPUStatus{
		Phase: v1alpha1.GPUBound,
	}
	egpu.Spec.NodeName = pod.Spec.NodeName
	_, err := eGPUClient.ElasticgpuV1alpha1().ElasticGPUs(pod.Namespace).Create(context.TODO(), egpu, metav1.CreateOptions{})
	if err != nil {
		return panic(err)
	}
}

```

