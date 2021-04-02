/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PciDeviceInfo Information about a PCI device
type PciDeviceInfo struct {
	Id  string `json:"id"`
	Bdf string `json:"bdf"`
}