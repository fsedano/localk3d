package main

import (
	"fmt"
	"log"

	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(app)
}

func createPod(ctx *pulumi.Context, n string, parent pulumi.Resource, image string, envNamePod *corev1.Pod) (*corev1.Pod, error) {

	depends := pulumi.DependsOn([]pulumi.Resource{})
	if parent != nil {
		depends = pulumi.DependsOn([]pulumi.Resource{parent})
	}

	env := corev1.EnvVarArray{}
	if envNamePod != nil {
		nm := envNamePod.Metadata.Name()
		env = corev1.EnvVarArray{
			corev1.EnvVarArgs{
				Name:  pulumi.String("PARENT_POD"),
				Value: nm.Elem().ToStringOutput(),
			},
		}

	} else {
		env = nil
	}
	pod, err := corev1.NewPod(ctx, n, &corev1.PodArgs{
		Spec: corev1.PodSpecArgs{
			Containers: corev1.ContainerArray{
				corev1.ContainerArgs{
					Name:  pulumi.String("nginx"),
					Image: pulumi.String(fmt.Sprintf("k3d-myreg:5000/%s", image)),
					Env:   env,
				},
			},
		},
	}, depends)
	if err != nil {
		return nil, err
	}
	return pod, err

}
func app(ctx *pulumi.Context) error {
	var pods []*corev1.Pod
	for i := 0; i < 3; i++ {

		p, err := createPod(ctx, fmt.Sprintf("pod-%d", i), nil, "nginx", nil)
		if err != nil {
			return err
		}
		pods = append(pods, p)
	}

	var xpods []*corev1.Pod

	for i := 0; i < 3; i++ {

		p, err := createPod(ctx, fmt.Sprintf("xpod-%d", i), pods[i], "nginx", pods[i])
		if err != nil {
			return err
		}
		xpods = append(xpods, p)
	}
	log.Printf("Pods created")
	ctx.Log.Info("Pods created", nil)
	for _, p := range pods {
		p.Metadata.Name().ApplyT(func(name *string) error {
			log.Printf("pod name=%s", *name)
			return nil
		})
	}
	for _, p := range xpods {
		p.Metadata.Name().ApplyT(func(name *string) error {
			log.Printf("xpod name=%s", *name)
			return nil
		})
	}

	return nil
}
