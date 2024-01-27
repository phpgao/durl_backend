package biz

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/phpgao/durl_backend/internal/conf"
	"github.com/phpgao/durl_backend/internal/data/ent"
)

type Leaf struct {
	ID        int64
	BizTag    string
	MaxID     int64
	Step      int64
	Desc      string
	Version   int32
	CreatedAt int64
	UpdatedAt int64
}

var once sync.Once

type LeafRepo interface {
	Save(context.Context, *Leaf) (int64, error)
	Update(context.Context, int64, *Leaf) error
	Find(context.Context, int64) (*Leaf, error)
	FindByBizTag(context.Context, string) (*Leaf, error)
	CountByBizTag(context.Context, string) (int, error)
	ListAll(context.Context) ([]*Leaf, error)
	UpdateBuffer(context.Context, int64) ([]int64, error)
}

type LeafUseCase struct {
	repo   LeafRepo
	log    *log.Helper
	buffer chan int64
	lock   *sync.Map
	config *conf.App
}

func NewLeafUseCase(repo LeafRepo, logger log.Logger, config *conf.App) *LeafUseCase {
	l := &LeafUseCase{
		repo:   repo,
		log:    log.NewHelper(logger),
		buffer: make(chan int64, config.DefaultStep),
		lock:   &sync.Map{},
		config: config,
	}
	once.Do(func() {
		l.log.Infof("init leaf,config: %v", config)
		err := l.init(context.Background())
		if err != nil {
			l.log.Errorf("init leaf error: %v", err)
			panic(err)
		}
	})
	return l
}

func (lu *LeafUseCase) FindByBizID(ctx context.Context, bizID int64) (*Leaf, error) {
	return lu.repo.Find(ctx, bizID)
}

func (lu *LeafUseCase) CountByBizTag(ctx context.Context, bizTag string) (int, error) {
	return lu.repo.CountByBizTag(ctx, bizTag)
}

func (lu *LeafUseCase) init(ctx context.Context) error {
	_, err := lu.repo.Find(ctx, 1)
	if err != nil {
		if ent.IsNotFound(err) {
			leaf := &Leaf{
				ID:        1,
				BizTag:    "initial_leaf",
				Step:      lu.config.DefaultStep,
				MaxID:     int64(math.Pow(62, float64(lu.config.ReservedDigits-1))),
				Version:   0,
				Desc:      "default leaf, don't change it",
				CreatedAt: time.Now().Unix(),
				UpdatedAt: time.Now().Unix(),
			}
			_, err := lu.repo.Save(ctx, leaf)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	err = lu.generateID(ctx, 1)
	if err != nil {
		lu.log.Errorf("error while init id for biz id: %s,err: %v", 1, err)
	}
	return nil
}

func (lu *LeafUseCase) generateID(ctx context.Context, bizID int64) error {
	_, locked := lu.lock.LoadOrStore(bizID, true)
	if locked {
		return nil
	}
	defer func() {
		lu.lock.Delete(bizID)
	}()
	ids, err := lu.repo.UpdateBuffer(ctx, bizID)
	if err != nil {
		lu.log.Error("error while generating id for biz id: %d,err: %v", bizID, err)
		return err
	}
	for _, v := range ids {
		lu.buffer <- v
	}
	return nil
}

func (lu *LeafUseCase) GetID(ctx context.Context) int64 {
	length := len(lu.buffer)
	capacity := cap(lu.buffer)
	percentage := float64(length) / float64(capacity) * 100
	if length == 0 || percentage <= 50 {
		lu.log.Infof("generate new id, length: %d,capacity: %d,percentage: %f", length, capacity, percentage)
		newCtx := context.Background()
		go func(ctx context.Context) {
			_ = lu.generateID(ctx, 1)
		}(newCtx)
	}
	newID := <-lu.buffer
	return newID
}

func (lu *LeafUseCase) Save(ctx context.Context, leaf *Leaf) (int64, error) {
	_, err := lu.repo.FindByBizTag(ctx, leaf.BizTag)
	if err != nil {
		if !ent.IsNotFound(err) {
			return 0, err
		}
	}
	return lu.repo.Save(ctx, leaf)
}
