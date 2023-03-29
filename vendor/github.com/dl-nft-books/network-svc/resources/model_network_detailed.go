/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NetworkDetailed struct {
	Key
	Attributes NetworkDetailedAttributes `json:"attributes"`
}
type NetworkDetailedResponse struct {
	Data     NetworkDetailed `json:"data"`
	Included Included        `json:"included"`
}

type NetworkDetailedListResponse struct {
	Data     []NetworkDetailed `json:"data"`
	Included Included          `json:"included"`
	Links    *Links            `json:"links"`
}

// MustNetworkDetailed - returns NetworkDetailed from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNetworkDetailed(key Key) *NetworkDetailed {
	var networkDetailed NetworkDetailed
	if c.tryFindEntry(key, &networkDetailed) {
		return &networkDetailed
	}
	return nil
}
