package redis

// redis keys use namespaces to facilitate query and splitting
const (
	Prefix                 = "acana:"      // project key prefix
	KeyPostTimeZSet        = "post:time"   // zset;post and the time of post
	KeyPostScoreZSet       = "post:score"  // zset;post abd the vates of post
	KeyPostVotedZSetPrefix = "post:voted:" // zset;record users and vote-types;param is the post_id

	KeyCommunitySetPrefix = "community:" // set;save the IDs of the posts in each community section
)

// concat prefix for key
func getRedisKey(key string) string {
	return Prefix + key
}
