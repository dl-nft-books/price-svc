/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Network struct {
	Key
	Attributes NetworkAttributes `json:"attributes"`
}
type NetworkResponse struct {
	Data     Network  `json:"data"`
	Included Included `json:"included"`
}

type NetworkListResponse struct {
	Data     []Network `json:"data"`
	Included Included  `json:"included"`
	Links    *Links    `json:"links"`
}

// MustNetwork - returns Network from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNetwork(key Key) *Network {
	var network Network
	if c.tryFindEntry(key, &network) {
		return &network
	}
	return nil
}
