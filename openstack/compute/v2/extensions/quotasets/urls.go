package quotasets

import "github.com/eugenetaranov/gophercloud"

const resourcePath = "os-quota-sets"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func getURL(c *gophercloud.ServiceClient, tenantID string) string {
	return c.ServiceURL(resourcePath, tenantID)
}

func updateURL(c *gophercloud.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}

func deleteURL(c *gophercloud.ServiceClient, tenantID string) string {
	return getURL(c, tenantID)
}
