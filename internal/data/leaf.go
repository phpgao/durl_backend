package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/phpgao/durl_backend/internal/biz"
	"github.com/phpgao/durl_backend/internal/conf"
	"github.com/phpgao/durl_backend/internal/data/ent"
	"github.com/phpgao/durl_backend/internal/data/ent/tleaf"
	"github.com/phpgao/durl_backend/internal/data/ent/tshorturl"
	"github.com/phpgao/durl_backend/internal/util"
)

type LeafRepo struct {
	data   *Data
	log    *log.Helper
	config *conf.App
}

func (lf *LeafRepo) UpdateBuffer(ctx context.Context, bizId int64) ([]int64, error) {
	lf.log.Infof("Start updating buffer for biz id: %d", bizId)
	ids := make([]int64, 0)

	// Start a new transaction
	tx, err := lf.data.dbClient.Tx(ctx)
	if err != nil {
		lf.log.Errorf("Failed to create transaction: %v", err)
		return nil, err
	}

	// Ensure the transaction is rolled back if an error occurs
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				// Log error if transaction rollback fails
				lf.log.Errorf("failed rolling back transaction: %v", rollbackErr)
			}
		}
	}()

	// Fetch the leaf with the given business ID
	leaf, err := tx.TLeaf.Get(ctx, bizId)
	if err != nil {
		lf.log.Errorf("Failed to query leaf with biz id: %d, error: %v", bizId, err)
		// Return a specific error if the business ID is not found
		if ent.IsNotFound(err) {
			err = errors.New(500, "BIZ_NOT_FOUND", "biz id not found")
		}
		return nil, err
	}

	// Prepare the new version and IDs for the update
	newVersion := leaf.Version + 1
	oldMaxID := leaf.MaxID + 1
	newMaxID := leaf.MaxID + leaf.Step
	affected := 0
	// Update the leaf with the new version and ID
	affected, err = tx.TLeaf.Update().
		Where(tleaf.ID(bizId)).
		Where(tleaf.Version(leaf.Version)).
		SetMaxID(newMaxID).
		SetVersion(newVersion).
		Save(ctx)

	// Return an error if the update fails or affects zero rows
	if err != nil || affected == 0 {
		lf.log.Errorf("Failed to update leaf with biz id: %d, error: %v", bizId, err)
		err = errors.New(500, "FAILED_TO_UPDATE_LEAF", "failed to update leaf")
		return nil, err
	}

	// Fetch the IDs to be skipped
	var skippedIDs []int64
	err = lf.data.dbClient.TShortUrl.Query().
		Where(tshorturl.ShortGTE(oldMaxID), tshorturl.ShortLTE(newMaxID)).
		Select(tshorturl.FieldShort).
		Scan(ctx, &skippedIDs)
	if err != nil {
		return nil, err
	}

	lf.log.Infof("Skipped ids: %v", skippedIDs)

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		lf.log.Errorf("Failed to commit transaction: %v", err)
		return nil, err
	}

	// Generate the IDs, skipping the ones fetched earlier
	ids = util.GenerateIDs(oldMaxID, newMaxID, skippedIDs)
	lf.log.Infof("Generated ids: %v", ids)

	// Return the generated IDs
	return ids, nil
}
func (lf *LeafRepo) Save(ctx context.Context, leaf *biz.Leaf) (int64, error) {
	currentTime := time.Now().Unix()
	l, err := lf.data.dbClient.TLeaf.
		Create().
		SetID(leaf.ID).
		SetBizTag(leaf.BizTag).
		SetStep(leaf.Step).
		SetMaxID(leaf.MaxID).
		SetVersion(leaf.Version).
		SetDesc(leaf.Desc).
		SetCreatedAt(currentTime).
		SetUpdatedAt(currentTime).
		Save(ctx)
	if err != nil {
		lf.log.Errorf("error while creating ShortUrl biz id: %s,err: %v", leaf.BizTag, err)
		return 0, err
	}
	return l.ID, nil
}

func (lf *LeafRepo) Update(ctx context.Context, id int64, leaf *biz.Leaf) error {
	l, err := lf.data.dbClient.TLeaf.Get(ctx, id)
	if err != nil {
		lf.log.Errorf("error while querying ShortUrl biz id: %s,err: %v", id, err)
		return err
	}
	currentTime := time.Now().Unix()
	_, err = l.Update().
		SetStep(leaf.Step).
		SetMaxID(leaf.MaxID).
		SetBizTag(leaf.BizTag).
		SetVersion(leaf.Version).
		SetDesc(leaf.Desc).
		SetUpdatedAt(currentTime).
		Save(ctx)
	return err
}

func (lf *LeafRepo) Find(ctx context.Context, id int64) (*biz.Leaf, error) {
	leaf, err := lf.data.dbClient.TLeaf.Get(ctx, id)
	if err != nil {
		lf.log.Errorf("error while querying leaf biz id: %s,err: %v", id, err)
		return nil, err
	}
	return &biz.Leaf{
		ID:        leaf.ID,
		BizTag:    leaf.BizTag,
		Step:      leaf.Step,
		MaxID:     leaf.MaxID,
		Desc:      leaf.Desc,
		Version:   leaf.Version,
		CreatedAt: leaf.CreatedAt,
		UpdatedAt: leaf.UpdatedAt,
	}, nil
}

func (lf *LeafRepo) FindByBizTag(ctx context.Context, tag string) (*biz.Leaf, error) {
	leaf, err := lf.data.dbClient.TLeaf.Query().Where(tleaf.BizTag(tag)).Only(ctx)
	if err != nil {
		lf.log.Errorf("error while querying leaf biz id: %s,err: %v", tag, err)
		return nil, err
	}
	return &biz.Leaf{
		ID:        leaf.ID,
		BizTag:    leaf.BizTag,
		Step:      leaf.Step,
		MaxID:     leaf.MaxID,
		Desc:      leaf.Desc,
		Version:   leaf.Version,
		CreatedAt: leaf.CreatedAt,
		UpdatedAt: leaf.UpdatedAt,
	}, nil
}

func (lf *LeafRepo) CountByBizTag(ctx context.Context, tag string) (int, error) {
	c, err := lf.data.dbClient.TLeaf.Query().Where(tleaf.BizTag(tag)).Count(ctx)
	if err != nil {
		lf.log.Errorf("error while querying leaf biz tag: %s,err: %v", tag, err)
		return 0, err
	}
	return c, nil
}

func (lf *LeafRepo) ListAll(ctx context.Context) ([]*biz.Leaf, error) {
	l, err := lf.data.dbClient.TLeaf.Query().All(ctx)
	if err != nil {
		lf.log.Errorf("error while querying leaf: %v", err)
		return nil, err
	}
	leafs := make([]*biz.Leaf, len(l))
	for i, leaf := range l {
		leafs[i] = &biz.Leaf{
			ID:        leaf.ID,
			BizTag:    leaf.BizTag,
			Step:      leaf.Step,
			MaxID:     leaf.MaxID,
			Desc:      leaf.Desc,
			Version:   leaf.Version,
			CreatedAt: leaf.CreatedAt,
			UpdatedAt: leaf.UpdatedAt,
		}
	}
	return leafs, nil
}

// NewLeafRepo .
func NewLeafRepo(data *Data, logger log.Logger, AppConfig *conf.App) biz.LeafRepo {
	return &LeafRepo{
		data:   data,
		config: AppConfig,
		log:    log.NewHelper(logger),
	}
}
