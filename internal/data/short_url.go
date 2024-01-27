package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/phpgao/durl_backend/internal/biz"
	"github.com/phpgao/durl_backend/internal/data/ent"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
	"github.com/phpgao/durl_backend/internal/util"
)

type ShortUrlRepo struct {
	data *Data
	log  *log.Helper
}

func (sr *ShortUrlRepo) Save(ctx context.Context, url *biz.ShortUrl) (err error) {
	if url.NeedCheck {
		// Start a new transaction
		tx, err := sr.data.dbClient.Tx(ctx)
		if err != nil {
			sr.log.Errorf("Failed to create transaction: %v", err)
			return err
		}

		// Ensure the transaction is rolled back if an error occurs
		defer func() {
			if err != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					// Log error if transaction rollback fails
					sr.log.Errorf("failed rolling back transaction: %v", rollbackErr)
				}
			}
		}()

		// Fetch the leaf with the given business ID
		var leaf *ent.TLeaf
		leaf, err = tx.TLeaf.Get(ctx, url.BizID)
		if err != nil {
			sr.log.Errorf("Failed to query leaf with biz id: %d, error: %v", url.BizID, err)
			// Return a specific error if the business ID is not found
			if ent.IsNotFound(err) {
				err = errors.New(500, "BIZ_NOT_FOUND", "biz id not found")
			}
			return err
		}

		// Prepare the maxId ID for the update
		maxId := leaf.MaxID + leaf.Step

		// Check if the short URL ID has been used
		if url.ShortInt <= maxId {
			sr.log.Infof("The short id has been used,maxId id(maxId+step): %d, short id: %d", maxId, url.ShortInt)
			err = errors.New(500, "SHORT_URL_USED", "short url has been used")
			return err
		}

		// Create a new short URL
		_, err = tx.TShortUrl.
			Create().
			SetBizID(url.BizID).
			SetOrigin(url.Origin).
			SetShort(url.ShortInt).
			SetCreatedAt(time.Now().Unix()).
			SetUpdatedAt(time.Now().Unix()).
			SetExpiredAt(url.ExpiredAt).
			Save(ctx)

		if err != nil {
			sr.log.Errorf("Failed to create short url: %v, error: %v", url, err)
			// Return a specific error if a constraint error occurs
			if ent.IsConstraintError(err) {
				err = errors.New(500, "SHORT_URL_USED", "short url has been used")
			}
			return err
		}

		// Update the leaf version
		affected := 0
		affected, err = tx.TLeaf.Update().
			Where(tleaf.ID(url.BizID)).
			Where(tleaf.Version(leaf.Version)).
			SetVersion(leaf.Version + 1).
			Save(ctx)

		if err != nil || affected == 0 {
			sr.log.Errorf("Failed to update leaf with biz id: %d,affected: %d, error: %v",
				url.BizID, affected, err)
			err = errors.New(500, "SHORT_URL_ERROR", "short url has been used")
			return err
		}

		// Commit the transaction
		if err = tx.Commit(); err != nil {
			sr.log.Errorf("Failed to commit transaction: %v", err)
			return err
		}
	} else {
		// Create a new short URL without checking
		_, err = sr.data.dbClient.TShortUrl.
			Create().
			SetBizID(url.BizID).
			SetOrigin(url.Origin).
			SetShort(url.ShortInt).
			SetCreatedAt(time.Now().Unix()).
			SetUpdatedAt(time.Now().Unix()).
			SetExpiredAt(url.ExpiredAt).
			Save(ctx)
	}
	if ent.IsConstraintError(err) {
		err = errors.New(500, "SHORT_URL_USED", "short url has been used")
	}

	return err
}

func (sr *ShortUrlRepo) Update(ctx context.Context, id int64, ShortUrl *biz.ShortUrl) error {
	l, err := sr.data.dbClient.TShortUrl.Get(ctx, id)
	if err != nil {
		return err
	}
	_, err = l.Update().
		SetBizID(ShortUrl.BizID).
		SetOrigin(ShortUrl.Origin).
		SetUpdatedAt(time.Now().Unix()).
		SetExpiredAt(ShortUrl.ExpiredAt).
		Save(ctx)
	return err
}

func (sr *ShortUrlRepo) Find(ctx context.Context, id int64) (*biz.ShortUrl, error) {
	l, err := sr.data.dbClient.TShortUrl.Get(ctx, id)
	if err != nil {
		sr.log.Errorf("error while querying ShortUrl biz id: %d,err: %v", id, err)
		return nil, err
	}
	return &biz.ShortUrl{
		ID:        l.ID,
		BizID:     l.BizID,
		Origin:    l.Origin,
		ShortInt:  l.Short,
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		ExpiredAt: l.UpdatedAt,
	}, nil
}

func (sr *ShortUrlRepo) FindByShort(ctx context.Context, short string) (*biz.ShortUrl, error) {
	id := util.ConvertFromBase62(short)
	leaf, err := sr.data.dbClient.TShortUrl.Query().Where(tshorturl.Short(id)).Only(ctx)
	if err != nil {
		sr.log.Errorf("error while querying ShortUrl : %s,err: %v", short, err)
		if ent.IsNotFound(err) {
			err = errors.New(404, "SHORT_URL_NOT_FOUND", "short url not found")
		}
		return nil, err
	}
	return &biz.ShortUrl{
		ID:        leaf.ID,
		BizID:     leaf.BizID,
		Origin:    leaf.Origin,
		ShortInt:  leaf.Short,
		CreatedAt: leaf.CreatedAt,
		UpdatedAt: leaf.UpdatedAt,
		ExpiredAt: leaf.ExpiredAt,
	}, nil
}

func (sr *ShortUrlRepo) ListAll(ctx context.Context) ([]*biz.ShortUrl, error) {
	l, err := sr.data.dbClient.TShortUrl.Query().All(ctx)
	if err != nil {
		sr.log.Errorf("error while querying ShortUrl: %v", err)
		if ent.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	ShortUrls := make([]*biz.ShortUrl, len(l))
	for i, ShortUrl := range l {
		ShortUrls[i] = &biz.ShortUrl{
			ID:        ShortUrl.ID,
			BizID:     ShortUrl.BizID,
			Origin:    ShortUrl.Origin,
			ShortInt:  ShortUrl.Short,
			CreatedAt: ShortUrl.CreatedAt,
			UpdatedAt: ShortUrl.UpdatedAt,
			ExpiredAt: ShortUrl.ExpiredAt,
		}
	}
	return ShortUrls, nil
}

func (sr *ShortUrlRepo) UpdateUrlCount(ctx context.Context, id int64, added int64) error {
	sr.log.Infof("adding %d to url id: %d", added, id)
	_, err := sr.data.dbClient.TShortUrl.UpdateOneID(id).AddVisit(added).Save(ctx)
	if err != nil {
		log.Errorf("failed to update url %d: %v", id, err)
		return err
	}

	return nil
}

// NewShortUrlRepo creates a new instance of ShortUrlRepo.
func NewShortUrlRepo(data *Data, logger log.Logger) biz.ShortUrlRepo {
	return &ShortUrlRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
