// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	api "github.com/aveshagarwal/rescheduler/pkg/api"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds,
		Convert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds,
		Convert_v1alpha1_ReschedulerPolicy_To_api_ReschedulerPolicy,
		Convert_api_ReschedulerPolicy_To_v1alpha1_ReschedulerPolicy,
		Convert_v1alpha1_ReschedulerStrategy_To_api_ReschedulerStrategy,
		Convert_api_ReschedulerStrategy_To_v1alpha1_ReschedulerStrategy,
		Convert_v1alpha1_StrategyParameters_To_api_StrategyParameters,
		Convert_api_StrategyParameters_To_v1alpha1_StrategyParameters,
	)
}

func autoConvert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds(in *NodeResourceUtilizationThresholds, out *api.NodeResourceUtilizationThresholds, s conversion.Scope) error {
	out.Thresholds = *(*api.ResourceThresholds)(unsafe.Pointer(&in.Thresholds))
	out.TargetThresholds = *(*api.ResourceThresholds)(unsafe.Pointer(&in.TargetThresholds))
	out.NumberOfNodes = in.NumberOfNodes
	return nil
}

// Convert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds is an autogenerated conversion function.
func Convert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds(in *NodeResourceUtilizationThresholds, out *api.NodeResourceUtilizationThresholds, s conversion.Scope) error {
	return autoConvert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds(in, out, s)
}

func autoConvert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds(in *api.NodeResourceUtilizationThresholds, out *NodeResourceUtilizationThresholds, s conversion.Scope) error {
	out.Thresholds = *(*ResourceThresholds)(unsafe.Pointer(&in.Thresholds))
	out.TargetThresholds = *(*ResourceThresholds)(unsafe.Pointer(&in.TargetThresholds))
	out.NumberOfNodes = in.NumberOfNodes
	return nil
}

// Convert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds is an autogenerated conversion function.
func Convert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds(in *api.NodeResourceUtilizationThresholds, out *NodeResourceUtilizationThresholds, s conversion.Scope) error {
	return autoConvert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds(in, out, s)
}

func autoConvert_v1alpha1_ReschedulerPolicy_To_api_ReschedulerPolicy(in *ReschedulerPolicy, out *api.ReschedulerPolicy, s conversion.Scope) error {
	out.Strategies = *(*api.StrategyList)(unsafe.Pointer(&in.Strategies))
	return nil
}

// Convert_v1alpha1_ReschedulerPolicy_To_api_ReschedulerPolicy is an autogenerated conversion function.
func Convert_v1alpha1_ReschedulerPolicy_To_api_ReschedulerPolicy(in *ReschedulerPolicy, out *api.ReschedulerPolicy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ReschedulerPolicy_To_api_ReschedulerPolicy(in, out, s)
}

func autoConvert_api_ReschedulerPolicy_To_v1alpha1_ReschedulerPolicy(in *api.ReschedulerPolicy, out *ReschedulerPolicy, s conversion.Scope) error {
	out.Strategies = *(*StrategyList)(unsafe.Pointer(&in.Strategies))
	return nil
}

// Convert_api_ReschedulerPolicy_To_v1alpha1_ReschedulerPolicy is an autogenerated conversion function.
func Convert_api_ReschedulerPolicy_To_v1alpha1_ReschedulerPolicy(in *api.ReschedulerPolicy, out *ReschedulerPolicy, s conversion.Scope) error {
	return autoConvert_api_ReschedulerPolicy_To_v1alpha1_ReschedulerPolicy(in, out, s)
}

func autoConvert_v1alpha1_ReschedulerStrategy_To_api_ReschedulerStrategy(in *ReschedulerStrategy, out *api.ReschedulerStrategy, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Weight = in.Weight
	if err := Convert_v1alpha1_StrategyParameters_To_api_StrategyParameters(&in.Params, &out.Params, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ReschedulerStrategy_To_api_ReschedulerStrategy is an autogenerated conversion function.
func Convert_v1alpha1_ReschedulerStrategy_To_api_ReschedulerStrategy(in *ReschedulerStrategy, out *api.ReschedulerStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_ReschedulerStrategy_To_api_ReschedulerStrategy(in, out, s)
}

func autoConvert_api_ReschedulerStrategy_To_v1alpha1_ReschedulerStrategy(in *api.ReschedulerStrategy, out *ReschedulerStrategy, s conversion.Scope) error {
	out.Enabled = in.Enabled
	out.Weight = in.Weight
	if err := Convert_api_StrategyParameters_To_v1alpha1_StrategyParameters(&in.Params, &out.Params, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_ReschedulerStrategy_To_v1alpha1_ReschedulerStrategy is an autogenerated conversion function.
func Convert_api_ReschedulerStrategy_To_v1alpha1_ReschedulerStrategy(in *api.ReschedulerStrategy, out *ReschedulerStrategy, s conversion.Scope) error {
	return autoConvert_api_ReschedulerStrategy_To_v1alpha1_ReschedulerStrategy(in, out, s)
}

func autoConvert_v1alpha1_StrategyParameters_To_api_StrategyParameters(in *StrategyParameters, out *api.StrategyParameters, s conversion.Scope) error {
	if err := Convert_v1alpha1_NodeResourceUtilizationThresholds_To_api_NodeResourceUtilizationThresholds(&in.NodeResourceUtilizationThresholds, &out.NodeResourceUtilizationThresholds, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_StrategyParameters_To_api_StrategyParameters is an autogenerated conversion function.
func Convert_v1alpha1_StrategyParameters_To_api_StrategyParameters(in *StrategyParameters, out *api.StrategyParameters, s conversion.Scope) error {
	return autoConvert_v1alpha1_StrategyParameters_To_api_StrategyParameters(in, out, s)
}

func autoConvert_api_StrategyParameters_To_v1alpha1_StrategyParameters(in *api.StrategyParameters, out *StrategyParameters, s conversion.Scope) error {
	if err := Convert_api_NodeResourceUtilizationThresholds_To_v1alpha1_NodeResourceUtilizationThresholds(&in.NodeResourceUtilizationThresholds, &out.NodeResourceUtilizationThresholds, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_StrategyParameters_To_v1alpha1_StrategyParameters is an autogenerated conversion function.
func Convert_api_StrategyParameters_To_v1alpha1_StrategyParameters(in *api.StrategyParameters, out *StrategyParameters, s conversion.Scope) error {
	return autoConvert_api_StrategyParameters_To_v1alpha1_StrategyParameters(in, out, s)
}
