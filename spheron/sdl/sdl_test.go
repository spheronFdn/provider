package sdl

import (
	"testing"

	"github.com/akash-network/provider/spheron/entities"
)

func TestSdlBuilder(t *testing.T) {

	sdlManifest, _ := ReadFile("../testdata/deployment.yaml")

	groups, _ := sdlManifest.DeploymentGroups()

	order := entities.TransformGroupToOrder(groups[0])

	t.Logf("order: %v+", order)
}
