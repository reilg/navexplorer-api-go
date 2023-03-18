package resource

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/navcoin/navexplorer-api-go/v2/internal/service/block"
)

type StatResource struct {
	blockService block.Service
}

func NewStatResource(blockService block.Service) *StatResource {
	return &StatResource{blockService}
}

func (r *StatResource) GetTotalSpendableSupply(c *gin.Context) {
	size, _ := strconv.Atoi(c.Query("size"))
	from, _ := strconv.Atoi(c.Query("from"))

	supply, err := r.blockService.GetTotalSpendableSupply(network(c), size, from)
	if err != nil {
		handleError(c, err, http.StatusInternalServerError)
		return
	}

	// GetAggregatedSpendableSupply (if more than 1k size?)
	// GetTotalSpendableSupply (if less than or equal to 1k size)

	c.JSON(200, supply)
}
