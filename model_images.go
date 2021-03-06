/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package dcos

type Images struct {
	// PNG icon URL, preferably 256 by 256 pixels.
	IconLarge string `json:"icon-large,omitempty"`
	// PNG icon URL, preferably 128 by 128 pixels.
	IconMedium string `json:"icon-medium,omitempty"`
	// PNG icon URL, preferably 48 by 48 pixels.
	IconSmall string `json:"icon-small,omitempty"`
	Screenshots []string `json:"screenshots,omitempty"`
}
