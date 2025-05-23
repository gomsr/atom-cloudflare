package kvs

import (
	"github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/gomsr/atom-cloudflare/util"
)

var (
	TokenKey     = "CLOUDFLARE_API_KEY_KV"
	Token        = "ENC(0TaSRlLDFeZJumH6XMXdi9hUnHJnMw0Y2jULnACnFiz7AmQbqo3HDoiv2HYgSSrmpNQ22uRLPTgIG9utt/RpYb96hftCjG/vlliesXo6bA4=)"
	AccountIDKey = "CLOUDFLARE_ACCOUNT_ID"
	AccountID    = "ENC(KvPG+Lg3btQhGyVT61jzH0Qt/1IyCjX4ki/XH+CSLSD5qsg/CJdvnJtgOgmCfaPn5+8qI+2/jTbt//5sBE7RV1PunI9nyq8iOTiJdGU2XVU=)"
	Client       *cloudflare.Client
)

func init() {
	// 1. parse kv storage key
	AccountID = util.GetEnvDefault(AccountIDKey, AccountID)

	// 2. parse kv storage key
	Token = util.GetEnvDefault(TokenKey, Token)

	// 3. init cloudflare api
	Client = cloudflare.NewClient(option.WithAPIToken(Token))
}
