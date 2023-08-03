// Copyright 2023 Jimmi Dyson.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var dummylog = logf.Log.WithName("dummy-resource")

func (r *Dummy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//nolint:lll // Long kubebuilder annotation for controller-gen.
//+kubebuilder:webhook:path=/mutate-dummy-example-com-v1alpha1-dummy,mutating=true,failurePolicy=fail,sideEffects=None,groups=dummy.example.com,resources=dummies,verbs=create;update,versions=v1alpha1,name=mdummy.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Dummy{}

// Default implements webhook.Defaulter so a webhook will be registered for the type.
func (r *Dummy) Default() {
	dummylog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//nolint:lll // Long kubebuilder annotation for controller-gen.
//+kubebuilder:webhook:path=/validate-dummy-example-com-v1alpha1-dummy,mutating=false,failurePolicy=fail,sideEffects=None,groups=dummy.example.com,resources=dummies,verbs=create;update,versions=v1alpha1,name=vdummy.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Dummy{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type.
func (r *Dummy) ValidateCreate() (admission.Warnings, error) {
	dummylog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type.
func (r *Dummy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	dummylog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type.
func (r *Dummy) ValidateDelete() (admission.Warnings, error) {
	dummylog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
