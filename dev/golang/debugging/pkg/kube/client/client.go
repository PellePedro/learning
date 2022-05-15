package client

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var (
	clientset *kubernetes.Clientset
)

func init() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

// ListPods returns a map of pod Names and podIP for a given namespace
func ListPods(namespace string) map[string]string {
	m := make(map[string]string)
	pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("Failed to list Pods")
		return m
	}
	for _, pod := range pods.Items {
		podIP := pod.Status.PodIP
		podName := pod.ObjectMeta.Name
		fmt.Printf("Found pod with name [%s] host IP[%s]", podName, podIP)
		m[podName] = podIP
	}
	return m
}
