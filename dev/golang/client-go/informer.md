```
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	labelSelectorKey    = "app"
	labelSelectorValue  = "myapp"
	resyncPeriodSeconds = 30
)

func main() {
	// Configuring the Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Setting up a watch for pod terminations in the specified namespaces
	namespaceList := getNamespaceList()
	labelSelector := labels.Set{labelSelectorKey: labelSelectorValue}.AsSelector()
	watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", metav1.NamespaceAll, nil)
	eventHandler := cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			// Handle the event of a pod being terminated
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				log.Printf("Unexpected object type: %T", obj)
				return
			}
			log.Printf("Pod %s terminated in namespace %s", pod.ObjectMeta.Name, pod.ObjectMeta.Namespace)
		},
	}

	// Setting up the informer
	informer := cache.NewSharedIndexInformer(watchlist, &corev1.Pod{}, resyncPeriodSeconds*time.Second, cache.Indexers{})
	informer.AddEventHandler(eventHandler)

	// Filtering informer to only watch pods in the specified namespaces
	informerWithFilter := cache.NewFilteredListWatchFromClient(informer.GetClient(), "pods", metav1.NamespaceAll, func(options *metav1.ListOptions) {
		options.LabelSelector = labelSelector.String()
		options.FieldSelector = fmt.Sprintf("metadata.namespace in (%s)", strings.Join(namespaceList, ","))
	})

	// Starting the informer
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		log.Println("Starting informer")
		if err := informerWithFilter.Run(ctx.Done()); err != nil {
			log.Fatalf("Error running informer: %v", err)
		}
		log.Println("Stopping informer")
	}()

	// Wait for termination signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	<-signalCh
}

func getNamespaceList() []string {
	namespaceList := os.Getenv("NAMESPACE_LIST")
	if namespaceList == "" {
		return []string{"namespace1", "namespace2", "namespace3"}
	}
	return strings.Split(namespaceList, ",")
}


```



```

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Configuring the Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	// Setting up a watch for pod terminations in the specified namespaces
	namespaceList := []string{"namespace1", "namespace2", "namespace3"}
	labelSelector := labels.Set{"app": "myapp"}.AsSelector()
	watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", metav1.NamespaceAll, nil)
	resyncPeriod := 30 * time.Second
	eventHandler := cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			// Handle the event of a pod being terminated
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				log.Printf("Unexpected object type: %T", obj)
				return
			}
			log.Printf("Pod %s terminated in namespace %s", pod.ObjectMeta.Name, pod.ObjectMeta.Namespace)
		},
	}

	// Setting up the informer with filter
	informerWithFilter := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", metav1.NamespaceAll, nil)
	informerWithFilter.ListFunc = func(options metav1.ListOptions) (runtime.Object, error) {
		options.LabelSelector = labelSelector.String()
		options.FieldSelector = fmt.Sprintf("metadata.namespace in (%s)", strings.Join(namespaceList, ","))
		return clientset.CoreV1().Pods(metav1.NamespaceAll).List(context.Background(), options)
	}
	informerWithFilter.WatchFunc = func(options metav1.ListOptions) (watch.Interface, error) {
		options.LabelSelector = labelSelector.String()
		options.FieldSelector = fmt.Sprintf("metadata.namespace in (%s)", strings.Join(namespaceList, ","))
		return clientset.CoreV1().Pods(metav1.NamespaceAll).Watch(context.Background(), options)
	}

	// Setting up the informer
	informer := cache.NewSharedInformer(watchlist, &corev1.Pod{}, resyncPeriod)
	informer.AddEventHandler(eventHandler)
	informer.AddIndexers(cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})

	// Adding informerWithFilter to the informer
	informer.AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: func(obj interface{}) bool {
			pod, ok := obj.(*corev1.Pod)
			if !ok {
				return false
			}
			if !labelSelector.Matches(labels.Set(pod.GetLabels())) {
				return false
			}
			return true
		},
		Handler: cache.ResourceEventHandlerFuncs{
			AddFunc:    informer.Add,
			UpdateFunc: informer.Update,
			DeleteFunc: informer.Delete,
		},
	})

	// Starting the informer
	stopper := make(chan struct{})
	defer close(stopper)
	go informer.Run(stopper)

	// Waiting indefinitely for the

```