package shorty

import (
	"context"
)

type Service interface {
	Get(ctx context.Context, shortname string) (ShortyURL, error)
	Shorten(ctx context.Context, original string) (ShortyURL, error)
}
