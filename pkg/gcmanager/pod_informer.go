// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package gcmanager

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

// startPodInformer will set up k8s pod informer in circle
func (s *SpiderGC) startPodInformer(ctx context.Context) {
	logger.Sugar().Infof("try to register pod informer")

	innerCtx, innerCancel := context.WithCancel(ctx)
	defer innerCancel()

	for {
		select {
		case <-ctx.Done():
			return
		case isLeader := <-s.leader.IsElected():
			// Proceed only if this pod is the leader
			if !isLeader {
				logger.Warn("Leader lost, stopping IP GC pod informer.")
				innerCancel()
				return
			}

			logger.Info("Create Pod informer")
			informerFactory := informers.NewSharedInformerFactory(s.k8ClientSet, 0)
			podInformer := informerFactory.Core().V1().Pods().Informer()
			_, err := podInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
				AddFunc:    s.onPodAdd,
				UpdateFunc: s.onPodUpdate,
				DeleteFunc: s.onPodDel,
			})
			if err != nil {
				logger.Error(err.Error())
				innerCancel()
				continue
			}
			s.informerFactory = informerFactory
			informerFactory.Start(innerCtx.Done())

			logger.Debug("Triggering scan all with leader elected")
			cacheSync := cache.WaitForCacheSync(innerCtx.Done())
			if !cacheSync {
				innerCancel()
				continue
			}
			// Notify the system to trigger a GC scan
			s.gcSignal <- struct{}{}

			// Wait for informer context to be done
			<-innerCtx.Done()

			logger.Error("K8s pod informer broken, restarting process")
		}
	}
}

// onPodAdd represents Pod informer Add Event
func (s *SpiderGC) onPodAdd(obj interface{}) {
	// backup controller could be elected as master
	isLeader := <-s.leader.IsElected()
	if !isLeader {
		return
	}

	pod := obj.(*corev1.Pod)
	podEntry, err := s.buildPodEntry(nil, pod, false)
	if nil != err {
		logger.Sugar().Errorf("onPodAdd: failed to build Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		return
	}

	// flush the pod database
	if podEntry != nil {
		err = s.GetPodDatabase().ApplyPodEntry(podEntry)
		if nil != err {
			logger.Sugar().Errorf("onPodAdd: failed to apply Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		}
	}
}

// onPodUpdate represents Pod informer Update Event
func (s *SpiderGC) onPodUpdate(oldObj interface{}, newObj interface{}) {
	// backup controller could be elected as master
	isLeader := <-s.leader.IsElected()
	if !isLeader {
		return
	}

	oldPod := oldObj.(*corev1.Pod)
	pod := newObj.(*corev1.Pod)
	podEntry, err := s.buildPodEntry(oldPod, pod, false)
	if nil != err {
		logger.Sugar().Errorf("onPodUpdate: failed to build Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		return
	}

	// flush the pod database
	if podEntry != nil {
		err = s.GetPodDatabase().ApplyPodEntry(podEntry)
		if nil != err {
			logger.Sugar().Errorf("onPodUpdate: failed to apply Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		}
	}
}

// onPodDel represents Pod informer Delete Event
func (s *SpiderGC) onPodDel(obj interface{}) {
	// backup controller could be elected as master
	isLeader := <-s.leader.IsElected()
	if !isLeader {
		return
	}

	pod := obj.(*corev1.Pod)
	logger.Sugar().Debugf("onPodDel: receive pod '%s/%s' deleted event", pod.Namespace, pod.Name)
	podEntry, err := s.buildPodEntry(nil, pod, true)
	if nil != err {
		logger.Sugar().Errorf("onPodDel: failed to build Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		return
	}

	if podEntry != nil {
		err = s.GetPodDatabase().ApplyPodEntry(podEntry)
		if nil != err {
			logger.Sugar().Errorf("onPodDel: failed to apply Pod Entry '%s/%s', error: %v", pod.Namespace, pod.Name, err)
		}
	} else {
		logger.Sugar().Debugf("onPodDel: discard to apply status '%v' PodEntry '%s/%s'", pod.Status.Phase, pod.Namespace, pod.Name)
	}
}
