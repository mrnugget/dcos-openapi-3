/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dcos

type PackageDescribeRequest struct {
	PackageName string `json:"packageName"`
	PackageVersion string `json:"packageVersion,omitempty"`
}
