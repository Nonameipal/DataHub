package main
import (
	"DataHub/internal/app"
	"DataHub/internal/logger"
)
func main() {
	log := logger.GetLogger()

	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("application stopped")
	}
}