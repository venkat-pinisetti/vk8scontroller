package controller

import (
	"fmt"
	"time"

	runtime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"

	clientset "k8s.io/vk8scontroller/pkg/generated/clientset/versioned"
	informers "k8s.io/vk8scontroller/pkg/generated/informers/externalversions/zbm/v1alpha1"
	listers "k8s.io/vk8scontroller/pkg/generated/listers/zbm/v1alpha1"
)

type Controller struct {
	zbmclientset clientset.Interface
	zbmsLister        listers.ZbmLister
	zbmsSynced        cache.InformerSynced
	workqueue workqueue.RateLimitingInterface
}

func NewController(
	zbmclientset clientset.Interface,
	zbmInformer informers.ZbmInformer) *Controller {

	controller := &Controller{
		zbmclientset:	zbmclientset,
		zbmsLister:		zbmInformer.Lister(),
		zbmsSynced:		zbmInformer.Informer().HasSynced,
		workqueue:		workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Zbms"),
	}

	klog.Info("Setting up event handlers")
	zbmInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueZbm,
		},
	)
	return controller
}

func (c *Controller) Run(stopCh <-chan struct{}) error {
	klog.Info("Starting Zbm controller")
	klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.zbmsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}
	klog.Info("Starting workers")
	go wait.Until(c.runWorker, time.Second, stopCh)
	klog.Info("Started workers")
	<-stopCh
	klog.Info("Shutting down workers")

	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

func (c *Controller) processNextWorkItem() bool {
	obj,_ := c.workqueue.Get()
	defer c.workqueue.Done(obj)
	c.workqueue.Forget(obj)
	klog.Infof("Successfully synced '%s'", obj)
	return true
}

func (c *Controller) enqueueZbm(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}
