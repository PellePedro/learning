

## DeleteHelmHook
```
func (l *K8SLCM) DeleteHelmHook(helmhooks *builder.Hook) error {
		name := hook.Resource.Object.GetName()
		gracePeriod := int64(0)
		switch hook.Resource.GVK.Kind {
		case "Job":
			// Delete the job with a Background propagation policy
			deletePolicy := metav1.DeletePropagationBackground
			_ = l.jobClient.Delete(context.TODO(), name, metav1.DeleteOptions{
				GracePeriodSeconds: &gracePeriod,
				PropagationPolicy:  &deletePolicy,
			})
			// Delete any associate Pods
			selector := fmt.Sprintf("job-name=%s", name)
			_ = retry.RetryOnConflict(retry.DefaultRetry, func() error {
				err := l.podClient.DeleteCollection(context.Background(),
					metav1.DeleteOptions{
						GracePeriodSeconds: &gracePeriod,
						PropagationPolicy:  &deletePolicy,
						TypeMeta:           metav1.TypeMeta{},
					}, metav1.ListOptions{
						LabelSelector: selector,
					})
				return err
			})
		case "Pod":
			//
			_ = l.podClient.Delete(context.TODO(), name, metav1.DeleteOptions{
				GracePeriodSeconds: &gracePeriod,
			})
		}
	return nil
}
```

## updatePodLabelfromParentJob
```
// updates the labels of pods to match the parent Job
func (l *K8SLCM) updatePodLabelfromParentJob(job *apiBatchv1.Job) error {
	// Skyramp managed Jobs creates Pods workloadsbut the labels are not
	// propagated from the Job to the created Pod.
	job, err := l.jobClient.Get(context.Background(), job.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}
	// matchlabel on a job
	// selector:
	//   matchLabels:
	//     controller-uid: 771aebf1-fe43-4068-9d63-c2f60a2142bd
	matchLabels := job.Spec.Selector.MatchLabels
	selector := labels.SelectorFromSet(matchLabels)

	// Get the pods that match the selector
	podList, err := l.podClient.List(context.Background(), metav1.ListOptions{
		LabelSelector: selector.String(),
	})
	if err != nil {
		return err
	}

	// Set the label to the pod
	for _, pod := range podList.Items {
		pod.Labels = builder.MergeStringMap(pod.Labels, job.Labels)
		_, err := l.podClient.Update(context.Background(), &pod, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}

	return nil
}
```
