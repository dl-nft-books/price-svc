package runners

import (
	"context"
	"fmt"
	"github.com/dl-nft-books/price-svc/internal/config"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko"
	"github.com/dl-nft-books/price-svc/internal/service/coingecko/models"
	"github.com/dl-nft-books/price-svc/resources"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"
)

type UpdatePricer struct {
	name       string
	logger     *logan.Entry
	runnerData config.RunnerData
	coingecko  *coingecko.Service
	platforms  map[string]config.MockedPlatform
}

func New(cfg config.Config) *UpdatePricer {
	return &UpdatePricer{
		name:       cfg.UpdatePricerCfg().Name,
		logger:     cfg.Log(),
		runnerData: cfg.UpdatePricerCfg().Runner,
		coingecko:  cfg.Coingecko(),
		platforms:  cfg.Mocked().Platforms,
	}
}

func (pc *UpdatePricer) updatePrices(ctx context.Context) error {
	platforms := make([]config.MockedPlatform, 0)
	for _, platform := range pc.platforms {
		// q is hardcoded
		if platform.Id == "q" {
			platforms = append(platforms, config.MockedPlatform{
				Id:               platform.Id,
				ChainId:          platform.ChainId,
				Name:             platform.Name,
				ShortName:        platform.ShortName,
				PricePerOneToken: platform.PricePerOneToken,
				PricePerOneNft:   platform.PricePerOneNft,
			})
			continue
		}

		price, err := pc.coingecko.GetPrice(platform.Id, "", "usd")
		if err != nil {
			pc.logger.WithError(err).Error("failed to get price")
			return errors.Wrap(err, "failed to get price")
		}
		fmt.Println("hope it works")
		fmt.Println(price)
		platforms = append(platforms, config.MockedPlatform{
			Id:               platform.Id,
			ChainId:          platform.ChainId,
			Name:             platform.Name,
			ShortName:        platform.ShortName,
			PricePerOneToken: price,
			PricePerOneNft:   platform.PricePerOneNft,
		})
	}
	return nil
}

func (pc *UpdatePricer) setCoigeckoPlatforms(platforms []config.MockedPlatform) *models.Platforms {
	platformsResp := resources.PlatformListResponse{}
	mapped := make(map[string]string)
	for _, platform := range platforms {
		mapped[platform.Name] = platform.Id
		chainIdentifier := platform.ChainId

		platformsResp.Data = append(platformsResp.Data, resources.Platform{
			Key: resources.Key{
				ID:   platform.Id,
				Type: resources.PLATFORMS,
			},
			Attributes: resources.PlatformAttributes{
				ChainIdentifier: chainIdentifier,
				Name:            platform.Name,
				Shortname:       platform.ShortName,
			},
		})
	}

	return &models.Platforms{
		Mapped:   mapped,
		Response: platformsResp,
	}
}

func (pc *UpdatePricer) Run(ctx context.Context) {
	running.WithBackOff(
		ctx, pc.logger, pc.name,
		pc.updatePrices,
		pc.runnerData.NormalPeriod, pc.runnerData.MinAbnormalPeriod, pc.runnerData.MaxAbnormalPeriod,
	)
}

func RunUpdatePricer(ctx context.Context, cfg config.Config) {
	New(cfg).Run(ctx)
}
