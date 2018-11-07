package stub

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"

	dnsv1alpha1 "github.com/openshift/cluster-dns-operator/pkg/apis/dns/v1alpha1"
	"github.com/openshift/cluster-dns-operator/pkg/util/clusteroperator"
	operatorversion "github.com/openshift/cluster-dns-operator/version"
	osv1 "github.com/openshift/cluster-version-operator/pkg/apis/operatorstatus.openshift.io/v1"

	"github.com/operator-framework/operator-sdk/pkg/k8sclient"
	"github.com/operator-framework/operator-sdk/pkg/sdk"
	"github.com/operator-framework/operator-sdk/pkg/util/k8sutil"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// syncOperatorStatus computes the operator's current status and therefrom
// creates or updates the ClusterOperator resource for the operator.
func (h *Handler) syncOperatorStatus() {
	co := &osv1.ClusterOperator{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterOperator",
			APIVersion: "operatorstatus.openshift.io/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Namespace: h.Namespace,
			// TODO Use a named constant or get name from config.
			Name: "openshift-dns",
		},
	}
	err := sdk.Get(co)
	isNotFound := errors.IsNotFound(err)
	if err != nil && !isNotFound {
		logrus.Errorf("syncOperatorStatus: error getting ClusterOperator %s/%s: %v",
			co.Namespace, co.Name, err)

		return
	}

	ns, dnses, daemonsets, err := h.getOperatorState()
	if err != nil {
		logrus.Errorf("syncOperatorStatus: getOperatorState: %v", err)

		return
	}

	oldConditions := co.Status.Conditions
	co.Status.Conditions = computeStatusConditions(oldConditions, ns,
		dnses, daemonsets)

	if isNotFound {
		co.Status.Version = operatorversion.Version

		if err := sdk.Create(co); err != nil {
			logrus.Errorf("syncOperatorStatus: failed to create ClusterOperator %s/%s: %v",
				co.Namespace, co.Name, err)
		} else {
			logrus.Infof("syncOperatorStatus: created ClusterOperator %s/%s (UID %v)",
				co.Namespace, co.Name, co.UID)
		}
	}

	if clusteroperator.ConditionsEqual(oldConditions, co.Status.Conditions) {
		return
	}

	unstructObj, err := k8sutil.UnstructuredFromRuntimeObject(co)
	if err != nil {
		logrus.Errorf("syncOperatorStatus: k8sutil.UnstructuredFromRuntimeObject: %v", err)

		return
	}

	resourceClient, _, err := k8sclient.GetResourceClient(co.APIVersion,
		co.Kind, co.Namespace)
	if err != nil {
		logrus.Errorf("syncOperatorStatus: GetResourceClient: %v", err)

		return
	}

	if _, err := resourceClient.UpdateStatus(unstructObj); err != nil {
		logrus.Errorf("syncOperatorStatus: UpdateStatus on %s/%s: %v",
			co.Namespace, co.Name, err)
	}
}

// getOperatorState gets and returns the resources necessary to compute the
// operator's current state.
func (h *Handler) getOperatorState() (*corev1.Namespace, []dnsv1alpha1.ClusterDNS, []appsv1.DaemonSet, error) {
	ns, err := h.ManifestFactory.DNSNamespace()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error building Namespace: %v",
			err)
	}

	if err := sdk.Get(ns); err != nil {
		if errors.IsNotFound(err) {
			return nil, nil, nil, nil
		}

		return nil, nil, nil, fmt.Errorf(
			"error getting Namespace %s: %v", ns.Name, err)
	}

	dnsList := &dnsv1alpha1.ClusterDNSList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ClusterDNS",
			APIVersion: "dns.openshift.io/v1alpha1",
		},
	}
	err = sdk.List(h.Namespace, dnsList,
		sdk.WithListOptions(&metav1.ListOptions{}))
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to list ClusterDNSes: %v", err)
	}

	daemonsetList := &appsv1.DaemonSetList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "DaemonSet",
			APIVersion: "apps/v1",
		},
	}
	err = sdk.List(ns.Name, daemonsetList,
		sdk.WithListOptions(&metav1.ListOptions{}))
	if err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to list DaemonSets: %v", err)
	}

	return ns, dnsList.Items, daemonsetList.Items, nil
}

// computeStatusConditions computes the operator's current state.
func computeStatusConditions(conditions []osv1.ClusterOperatorStatusCondition, ns *corev1.Namespace, dnses []dnsv1alpha1.ClusterDNS, daemonsets []appsv1.DaemonSet) []osv1.ClusterOperatorStatusCondition {
	failingCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorFailing,
		Status: osv1.ConditionUnknown,
	}
	if ns == nil {
		failingCondition.Status = osv1.ConditionTrue
		failingCondition.Reason = "NoNamespace"
		failingCondition.Message = "DNS namespace does not exist"
	} else {
		failingCondition.Status = osv1.ConditionFalse
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		failingCondition)

	progressingCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorProgressing,
		Status: osv1.ConditionUnknown,
	}
	numClusterDNSes := len(dnses)
	numDaemonsets := len(daemonsets)
	if numClusterDNSes == numDaemonsets {
		progressingCondition.Status = osv1.ConditionFalse
	} else {
		progressingCondition.Status = osv1.ConditionTrue
		progressingCondition.Reason = "Reconciling"
		progressingCondition.Message = fmt.Sprintf(
			"have %d DaemonSets, want %d",
			numDaemonsets, numClusterDNSes)
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		progressingCondition)

	availableCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorAvailable,
		Status: osv1.ConditionUnknown,
	}
	dsAvailable := map[string]bool{}
	for _, ds := range daemonsets {
		dsAvailable[ds.Name] = ds.Status.NumberAvailable > 0
	}
	unavailable := []string{}
	for _, dns := range dnses {
		// TODO Use the manifest to derive the name, or use labels or
		// owner references.
		name := "dns-" + dns.Name
		if available, exists := dsAvailable[name]; !exists {
			msg := fmt.Sprintf("no DaemonSet for ClusterDNS %q",
				dns.Name)
			unavailable = append(unavailable, msg)
		} else if !available {
			msg := fmt.Sprintf("DaemonSet %q not available", name)
			unavailable = append(unavailable, msg)
		}
	}
	if len(unavailable) == 0 {
		availableCondition.Status = osv1.ConditionTrue
	} else {
		availableCondition.Status = osv1.ConditionFalse
		availableCondition.Reason = "DaemonSetNotAvailable"
		availableCondition.Message = strings.Join(unavailable, "\n")
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		availableCondition)

	return conditions
}
