// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes Controllers
package controller

import (
	simple_skv2_solo_io_v1alpha1 "github.com/solo-io/solo-kit-example/simple/pkg/api/simple.skv2.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the Circle Resource.
// implemented by the user
type CircleReconciler interface {
	ReconcileCircle(obj *simple_skv2_solo_io_v1alpha1.Circle) (reconcile.Result, error)
}

// Reconcile deletion events for the Circle Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type CircleDeletionReconciler interface {
	ReconcileCircleDeletion(req reconcile.Request)
}

type CircleReconcilerFuncs struct {
	OnReconcileCircle         func(obj *simple_skv2_solo_io_v1alpha1.Circle) (reconcile.Result, error)
	OnReconcileCircleDeletion func(req reconcile.Request)
}

func (f *CircleReconcilerFuncs) ReconcileCircle(obj *simple_skv2_solo_io_v1alpha1.Circle) (reconcile.Result, error) {
	if f.OnReconcileCircle == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileCircle(obj)
}

func (f *CircleReconcilerFuncs) ReconcileCircleDeletion(req reconcile.Request) {
	if f.OnReconcileCircleDeletion == nil {
		return
	}
	f.OnReconcileCircleDeletion(req)
}

// Reconcile and finalize the Circle Resource
// implemented by the user
type CircleFinalizer interface {
	CircleReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	CircleFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeCircle(obj *simple_skv2_solo_io_v1alpha1.Circle) error
}

type CircleReconcileLoop interface {
	RunCircleReconciler(rec CircleReconciler, predicates ...predicate.Predicate) error
}

type circleReconcileLoop struct {
	loop reconcile.Loop
}

func NewCircleReconcileLoop(name string, mgr manager.Manager) CircleReconcileLoop {
	return &circleReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &simple_skv2_solo_io_v1alpha1.Circle{}),
	}
}

func (c *circleReconcileLoop) RunCircleReconciler(reconciler CircleReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericCircleReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(CircleFinalizer); ok {
		reconcilerWrapper = genericCircleFinalizer{
			genericCircleReconciler: genericReconciler,
			finalizingReconciler:    finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(reconcilerWrapper, predicates...)
}

// genericCircleHandler implements a generic reconcile.Reconciler
type genericCircleReconciler struct {
	reconciler CircleReconciler
}

func (r genericCircleReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*simple_skv2_solo_io_v1alpha1.Circle)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Circle handler received event for %T", object)
	}
	return r.reconciler.ReconcileCircle(obj)
}

func (r genericCircleReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(CircleDeletionReconciler); ok {
		deletionReconciler.ReconcileCircleDeletion(request)
	}
}

// genericCircleFinalizer implements a generic reconcile.FinalizingReconciler
type genericCircleFinalizer struct {
	genericCircleReconciler
	finalizingReconciler CircleFinalizer
}

func (r genericCircleFinalizer) FinalizerName() string {
	return r.finalizingReconciler.CircleFinalizerName()
}

func (r genericCircleFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*simple_skv2_solo_io_v1alpha1.Circle)
	if !ok {
		return errors.Errorf("internal error: Circle handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeCircle(obj)
}

// Reconcile Upsert events for the Square Resource.
// implemented by the user
type SquareReconciler interface {
	ReconcileSquare(obj *simple_skv2_solo_io_v1alpha1.Square) (reconcile.Result, error)
}

// Reconcile deletion events for the Square Resource.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type SquareDeletionReconciler interface {
	ReconcileSquareDeletion(req reconcile.Request)
}

type SquareReconcilerFuncs struct {
	OnReconcileSquare         func(obj *simple_skv2_solo_io_v1alpha1.Square) (reconcile.Result, error)
	OnReconcileSquareDeletion func(req reconcile.Request)
}

func (f *SquareReconcilerFuncs) ReconcileSquare(obj *simple_skv2_solo_io_v1alpha1.Square) (reconcile.Result, error) {
	if f.OnReconcileSquare == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileSquare(obj)
}

func (f *SquareReconcilerFuncs) ReconcileSquareDeletion(req reconcile.Request) {
	if f.OnReconcileSquareDeletion == nil {
		return
	}
	f.OnReconcileSquareDeletion(req)
}

// Reconcile and finalize the Square Resource
// implemented by the user
type SquareFinalizer interface {
	SquareReconciler

	// name of the finalizer used by this handler.
	// finalizer names should be unique for a single task
	SquareFinalizerName() string

	// finalize the object before it is deleted.
	// Watchers created with a finalizing handler will a
	FinalizeSquare(obj *simple_skv2_solo_io_v1alpha1.Square) error
}

type SquareReconcileLoop interface {
	RunSquareReconciler(rec SquareReconciler, predicates ...predicate.Predicate) error
}

type squareReconcileLoop struct {
	loop reconcile.Loop
}

func NewSquareReconcileLoop(name string, mgr manager.Manager) SquareReconcileLoop {
	return &squareReconcileLoop{
		loop: reconcile.NewLoop(name, mgr, &simple_skv2_solo_io_v1alpha1.Square{}),
	}
}

func (c *squareReconcileLoop) RunSquareReconciler(reconciler SquareReconciler, predicates ...predicate.Predicate) error {
	genericReconciler := genericSquareReconciler{
		reconciler: reconciler,
	}

	var reconcilerWrapper reconcile.Reconciler
	if finalizingReconciler, ok := reconciler.(SquareFinalizer); ok {
		reconcilerWrapper = genericSquareFinalizer{
			genericSquareReconciler: genericReconciler,
			finalizingReconciler:    finalizingReconciler,
		}
	} else {
		reconcilerWrapper = genericReconciler
	}
	return c.loop.RunReconciler(reconcilerWrapper, predicates...)
}

// genericSquareHandler implements a generic reconcile.Reconciler
type genericSquareReconciler struct {
	reconciler SquareReconciler
}

func (r genericSquareReconciler) Reconcile(object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*simple_skv2_solo_io_v1alpha1.Square)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: Square handler received event for %T", object)
	}
	return r.reconciler.ReconcileSquare(obj)
}

func (r genericSquareReconciler) ReconcileDeletion(request reconcile.Request) {
	if deletionReconciler, ok := r.reconciler.(SquareDeletionReconciler); ok {
		deletionReconciler.ReconcileSquareDeletion(request)
	}
}

// genericSquareFinalizer implements a generic reconcile.FinalizingReconciler
type genericSquareFinalizer struct {
	genericSquareReconciler
	finalizingReconciler SquareFinalizer
}

func (r genericSquareFinalizer) FinalizerName() string {
	return r.finalizingReconciler.SquareFinalizerName()
}

func (r genericSquareFinalizer) Finalize(object ezkube.Object) error {
	obj, ok := object.(*simple_skv2_solo_io_v1alpha1.Square)
	if !ok {
		return errors.Errorf("internal error: Square handler received event for %T", object)
	}
	return r.finalizingReconciler.FinalizeSquare(obj)
}
