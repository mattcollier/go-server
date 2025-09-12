package app

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func GetReplicasets() ([]string, error) {
	fmt.Println("getting replicasets")
	ctx := context.Background()
	cfg, _ := rest.InClusterConfig()
	clientset, _ := kubernetes.NewForConfig(cfg)
	// use clientset.CoreV1().Pods("default").List(ctx, opts)
	rsList, err := clientset.AppsV1().ReplicaSets("default").List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	names := make([]string, 0, len(rsList.Items))
	for _, rs := range rsList.Items {
		names = append(names, rs.Name)
	}
	return names, nil
}
