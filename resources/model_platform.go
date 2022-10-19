/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Platform struct {
	Key
	Attributes PlatformAttributes `json:"attributes"`
}
type PlatformResponse struct {
	Data     Platform `json:"data"`
	Included Included `json:"included"`
}

type PlatformListResponse struct {
	Data     []Platform `json:"data"`
	Included Included   `json:"included"`
	Links    *Links     `json:"links"`
}

// MustPlatform - returns Platform from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustPlatform(key Key) *Platform {
	var platform Platform
	if c.tryFindEntry(key, &platform) {
		return &platform
	}
	return nil
}
