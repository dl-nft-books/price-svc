/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Price struct {
	Key
	Attributes PriceAttributes `json:"attributes"`
}
type PriceResponse struct {
	Data     Price    `json:"data"`
	Included Included `json:"included"`
}

type PriceListResponse struct {
	Data     []Price  `json:"data"`
	Included Included `json:"included"`
	Links    *Links   `json:"links"`
}

// MustPrice - returns Price from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPrice(key Key) *Price {
	var price Price
	if c.tryFindEntry(key, &price) {
		return &price
	}
	return nil
}
