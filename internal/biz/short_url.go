package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/phpgao/durl_backend/internal/conf"
)

type ShortUrl struct {
	ID        int64
	BizID     int64
	Origin    string
	ShortInt  int64
	CreatedAt int64
	UpdatedAt int64
	ExpiredAt int64
	NeedCheck bool
}

type ShortUrlRepo interface {
	Save(context.Context, *ShortUrl) error
	Update(context.Context, int64, *ShortUrl) error
	Find(context.Context, int64) (*ShortUrl, error)
	ListAll(context.Context) ([]*ShortUrl, error)
	FindByShort(context.Context, string) (*ShortUrl, error)
	UpdateUrlCount(context.Context, int64, int64) error
}

type ShortUrlUseCase struct {
	ShortUrlRepo ShortUrlRepo
	LeafRepo     LeafRepo
	log          *log.Helper
	visitCounter map[int64]int64
	countChan    chan int64
	config       *conf.App
}

func NewShortUrlUseCase(repo ShortUrlRepo, leafRepo LeafRepo, logger log.Logger) *ShortUrlUseCase {
	s := &ShortUrlUseCase{
		ShortUrlRepo: repo,
		LeafRepo:     leafRepo,
		log:          log.NewHelper(logger),
		visitCounter: make(map[int64]int64),
		countChan:    make(chan int64, 1000),
	}
	go s.CountForever(context.Background())
	return s
}

func (su *ShortUrlUseCase) Save(ctx context.Context, short *ShortUrl) error {
	return su.ShortUrlRepo.Save(ctx, short)
}

func (su *ShortUrlUseCase) FindByShort(ctx context.Context, short string) (*ShortUrl, error) {
	s, err := su.ShortUrlRepo.FindByShort(ctx, short)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (su *ShortUrlUseCase) CountForever(ctx context.Context) {
	for {
		select {
		case urlID := <-su.countChan:
			su.visitCounter[urlID]++
		case <-time.After(5 * time.Second):
			for urlID, visits := range su.visitCounter {
				if visits > 0 {
					err := su.ShortUrlRepo.UpdateUrlCount(ctx, urlID, visits)
					if err != nil {
						log.Errorf("failed to update url %d: %v", urlID, err)
					} else {
						su.visitCounter[urlID] = 0
					}
				}
			}
		}
	}
}

// UpdateVisitCount url visit count + 1
func (su *ShortUrlUseCase) UpdateVisitCount(ctx context.Context, id int64) {
	su.log.Infof("update visit count to channel: %d", id)
	su.countChan <- id
}

//todo 自定义链接
//todo 延迟消费
//todo 多DB支持
