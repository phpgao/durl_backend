package service

import (
	"context"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/phpgao/durl_backend/api/url_shorter/v1"
	"github.com/phpgao/durl_backend/internal/biz"
	"github.com/phpgao/durl_backend/internal/conf"
	"github.com/phpgao/durl_backend/internal/util"
)

type UrlShortenerService struct {
	v1.UnimplementedUrlShortenerServiceServer
	Logger          *log.Helper
	LeafUseCase     *biz.LeafUseCase
	ShortUrlUseCase *biz.ShortUrlUseCase
	config          *conf.App
}

func NewUrlShortenerService(config *conf.App, leafUseCase *biz.LeafUseCase, ShortUrlUseCase *biz.ShortUrlUseCase, logger log.Logger) *UrlShortenerService {
	return &UrlShortenerService{
		LeafUseCase:     leafUseCase,
		ShortUrlUseCase: ShortUrlUseCase,
		Logger:          log.NewHelper(logger),
		config:          config,
	}
}

func (s *UrlShortenerService) CreateShortUrl(ctx context.Context, in *v1.CreateUrlRequest) (*v1.CreateUrlReply, error) {
	if in.BizId <= 0 {
		in.BizId = 1
	} else if _, err := s.LeafUseCase.FindByBizID(ctx, in.BizId); err != nil {
		return nil, errors.New(500, "BIZ_NOT_FOUND", "business ID not found")
	}

	in.Url = strings.TrimSpace(in.Url)
	if !util.IsValidURL(in.Url) {
		return nil, errors.New(500, "INVALID_URL_ADDRESS", "invalid url address")
	}

	url := &biz.ShortUrl{
		BizID:     in.BizId,
		Origin:    in.Url,
		ExpiredAt: in.ExpiredAt,
	}

	in.ShortKey = strings.TrimSpace(in.ShortKey)
	if in.ShortKey == "" {
		url.ShortInt = s.LeafUseCase.GetID(ctx)
	} else if util.IsValidBase62String(in.ShortKey) {
		if int64(len(in.ShortKey)) > s.config.ReservedDigits {
			url.NeedCheck = true
		}
		url.ShortInt = util.ConvertFromBase62(in.ShortKey)
	} else {
		return nil, errors.New(500, "INVALID_SHORT_KEY", "invalid short key")
	}

	if err := s.ShortUrlUseCase.Save(ctx, url); err != nil {
		return nil, err
	}

	return &v1.CreateUrlReply{Url: util.ConvertToBase62(url.ShortInt)}, nil
}

func (s *UrlShortenerService) GetRedirectURL(ctx context.Context, in *v1.JumpRequest) (*v1.JumpReply, error) {
	url, err := s.ShortUrlUseCase.FindByShort(ctx, in.Url)
	if err != nil {
		return nil, err
	}
	if url.ExpiredAt > 0 && url.ExpiredAt <= time.Now().Unix() {
		return nil, errors.New(500, "URL_EXPIRED", "url expired")
	}
	return &v1.JumpReply{Id: url.ID, Url: url.Origin}, nil
}
