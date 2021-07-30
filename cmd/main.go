package main

import (
	"os"
	"time"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	"k8s.io/vk8scontroller/pkg/controller"
	cdclientset "k8s.io/vk8scontroller/pkg/generated/clientset/versioned"
	cdinformers "k8s.io/vk8scontroller/pkg/generated/informers/externalversions"
)

func main() {
	klog.Info("Setting Kube client")
	kubeconfig := os.Getenv("KUBECONFIG")

	// setup connection
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}
	// setup clientset
	clientset, err := cdclientset.NewForConfig(config)

	//kubeInformerFactory := informers.NewSharedInformerFactory(kclientset, time.Second*3)
	zbminformerfactory := cdinformers.NewSharedInformerFactory(clientset, time.Second*3)
	//informer := factory.Ibm().V1alpha1().Zbms().Informer()
	// create controller
	klog.Info("Creating controller")
	stopCh := make(chan struct{})
	vcontroller := controller.NewController(clientset, zbminformerfactory.Ibm().V1alpha1().Zbms())
	zbminformerfactory.Start(stopCh)
	//vcontroller := controller.NewController(clientset)
	//controller.NewController(clientset).Run(stopCh)
	vcontroller.Run(stopCh)
	klog.Info("Completed")
}
