/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dcos

type InstallResponse struct {
	AppId string `json:"appId,omitempty"`
	Cli Cli `json:"cli,omitempty"`
	PackageName string `json:"packageName"`
	PackageVersion string `json:"packageVersion"`
	PostInstallNotes string `json:"postInstallNotes,omitempty"`
}