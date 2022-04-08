package shorty

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
)

type ShortyURL struct {
	OriginalURL string
	Shortname   string
}

type shortyService struct{}

func NewService() Service { return &shortyService{} }

func (w *shortyService) Get(_ context.Context, shortname string) (ShortyURL, error) {
	// TODO
	res := ShortyURL{}
	return res, nil
}

func (w *shortyService) Shorten(_ context.Context, original string) (ShortyURL, error) {
	// TODO
	res := ShortyURL{}
	return res, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
