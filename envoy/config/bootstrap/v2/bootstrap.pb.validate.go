// Code generated by protoc-gen-validate
// source: envoy/config/bootstrap/v2/bootstrap.proto
// DO NOT EDIT!!!

package v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Bootstrap) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "Node",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStaticResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "StaticResources",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDynamicResources()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "DynamicResources",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetClusterManager()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "ClusterManager",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "HdsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	// no validation rules for FlagsPath

	for idx, item := range m.GetStatsSinks() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					Field:  fmt.Sprintf("StatsSinks[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetStatsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "StatsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetStatsFlushInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "StatsFlushInterval",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetWatchdog()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "Watchdog",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTracing()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "Tracing",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetRuntime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "Runtime",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAdmin()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "Admin",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetOverloadManager()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BootstrapValidationError{
				Field:  "OverloadManager",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = BootstrapValidationError{}

// Validate checks the field values on Admin with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Admin) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessLogPath

	// no validation rules for ProfilePath

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdminValidationError{
				Field:  "Address",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// AdminValidationError is the validation error returned by Admin.Validate if
// the designated constraints aren't met.
type AdminValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e AdminValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmin.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = AdminValidationError{}

// Validate checks the field values on ClusterManager with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClusterManager) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for LocalClusterName

	if v, ok := interface{}(m.GetOutlierDetection()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterManagerValidationError{
				Field:  "OutlierDetection",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpstreamBindConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterManagerValidationError{
				Field:  "UpstreamBindConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLoadStatsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ClusterManagerValidationError{
				Field:  "LoadStatsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// ClusterManagerValidationError is the validation error returned by
// ClusterManager.Validate if the designated constraints aren't met.
type ClusterManagerValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ClusterManagerValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterManager.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ClusterManagerValidationError{}

// Validate checks the field values on Watchdog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Watchdog) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetMissTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchdogValidationError{
				Field:  "MissTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMegamissTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchdogValidationError{
				Field:  "MegamissTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetKillTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchdogValidationError{
				Field:  "KillTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMultikillTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return WatchdogValidationError{
				Field:  "MultikillTimeout",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// WatchdogValidationError is the validation error returned by
// Watchdog.Validate if the designated constraints aren't met.
type WatchdogValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e WatchdogValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchdog.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = WatchdogValidationError{}

// Validate checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Runtime) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSymlinkRoot()) < 1 {
		return RuntimeValidationError{
			Field:  "SymlinkRoot",
			Reason: "value length must be at least 1 bytes",
		}
	}

	// no validation rules for Subdirectory

	// no validation rules for OverrideSubdirectory

	return nil
}

// RuntimeValidationError is the validation error returned by Runtime.Validate
// if the designated constraints aren't met.
type RuntimeValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e RuntimeValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = RuntimeValidationError{}

// Validate checks the field values on Bootstrap_StaticResources with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Bootstrap_StaticResources) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetListeners() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Bootstrap_StaticResourcesValidationError{
					Field:  fmt.Sprintf("Listeners[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetClusters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Bootstrap_StaticResourcesValidationError{
					Field:  fmt.Sprintf("Clusters[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetSecrets() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Bootstrap_StaticResourcesValidationError{
					Field:  fmt.Sprintf("Secrets[%v]", idx),
					Reason: "embedded message failed validation",
					Cause:  err,
				}
			}
		}

	}

	return nil
}

// Bootstrap_StaticResourcesValidationError is the validation error returned by
// Bootstrap_StaticResources.Validate if the designated constraints aren't met.
type Bootstrap_StaticResourcesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Bootstrap_StaticResourcesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap_StaticResources.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Bootstrap_StaticResourcesValidationError{}

// Validate checks the field values on Bootstrap_DynamicResources with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Bootstrap_DynamicResources) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Bootstrap_DynamicResourcesValidationError{
				Field:  "LdsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Bootstrap_DynamicResourcesValidationError{
				Field:  "CdsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAdsConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Bootstrap_DynamicResourcesValidationError{
				Field:  "AdsConfig",
				Reason: "embedded message failed validation",
				Cause:  err,
			}
		}
	}

	return nil
}

// Bootstrap_DynamicResourcesValidationError is the validation error returned
// by Bootstrap_DynamicResources.Validate if the designated constraints aren't met.
type Bootstrap_DynamicResourcesValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e Bootstrap_DynamicResourcesValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap_DynamicResources.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = Bootstrap_DynamicResourcesValidationError{}

// Validate checks the field values on ClusterManager_OutlierDetection with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterManager_OutlierDetection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EventLogPath

	return nil
}

// ClusterManager_OutlierDetectionValidationError is the validation error
// returned by ClusterManager_OutlierDetection.Validate if the designated
// constraints aren't met.
type ClusterManager_OutlierDetectionValidationError struct {
	Field  string
	Reason string
	Cause  error
	Key    bool
}

// Error satisfies the builtin error interface
func (e ClusterManager_OutlierDetectionValidationError) Error() string {
	cause := ""
	if e.Cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.Cause)
	}

	key := ""
	if e.Key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterManager_OutlierDetection.%s: %s%s",
		key,
		e.Field,
		e.Reason,
		cause)
}

var _ error = ClusterManager_OutlierDetectionValidationError{}
