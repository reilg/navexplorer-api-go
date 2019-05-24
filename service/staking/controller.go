package staking

import (
	"errors"
	"github.com/NavExplorer/navexplorer-api-go/error"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type Controller struct{}

func (controller *Controller) GetStakingReport(c *gin.Context) {
	stakingEstimate, err := GetStakingReport()
	if err != nil {
		error.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(200, stakingEstimate)
}

func (controller *Controller) GetStakingByBlockCount(c *gin.Context) {
	blockCount, err := strconv.Atoi(c.DefaultQuery("blocks", "1000"))
	if err != nil {
		blockCount = 1000
	}
	if blockCount > 50000000 {
		blockCount = 50000000
	}

	staking, err := GetStakingByBlockCount(blockCount)
	if err != nil {
		error.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(200, staking)
}

func (controller *Controller) GetStakingRewardsForAddresses(c *gin.Context) {
	addresses := strings.Split(c.Query("addresses"), ",")
	if len(addresses) == 0 {
		error.HandleError(c, errors.New("No addresses provided"), http.StatusBadRequest)
		return
	}

	rewards, err := GetStakingRewardsForAddresses(addresses)
	if err != nil {
		error.HandleError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(200, rewards)
}