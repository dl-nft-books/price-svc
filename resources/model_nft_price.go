/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type NftPrice struct {
	Key
	Attributes NftPriceAttributes `json:"attributes"`
}
type NftPriceResponse struct {
	Data     NftPrice `json:"data"`
	Included Included `json:"included"`
}

type NftPriceListResponse struct {
	Data     []NftPrice `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustNftPrice - returns NftPrice from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustNftPrice(key Key) *NftPrice {
	var nftPrice NftPrice
	if c.tryFindEntry(key, &nftPrice) {
		return &nftPrice
	}
	return nil
}
