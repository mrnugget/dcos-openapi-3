/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dcos

type Binaries struct {
	Darwin Os `json:"darwin,omitempty"`
	Linux Os `json:"linux,omitempty"`
	Windows Os `json:"windows,omitempty"`
}
