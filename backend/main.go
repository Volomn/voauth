package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/Volomn/voauth/backend/api"
	"github.com/Volomn/voauth/backend/app"
	"github.com/Volomn/voauth/backend/infra"
	"github.com/Volomn/voauth/backend/infra/repository"
	"github.com/caarlos0/env/v9"
	"github.com/go-chi/chi/v5"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

type config struct {
	Port             int    `env:"PORT" envDefault:"5000"`
	DatabaseHost     string `env:"DATABASE_HOST" envDefault:"voauth_db"`
	DatabasePort     int    `env:"DATABASE_PORT" envDefault:"5432" `
	DatabaseUser     string `env:"DATABASE_USER" envDefault:"voauth"`
	DatabasePassword string `env:"DATABASE_PASSWORD" envDefault:"voauth"`
	DatabaseName     string `env:"DATABASE_NAME" envDefault:"voauth"`
	AuthSecretKey    string `env:"AUTH_SECRET_KEY" envDefault:"devsecret"`
}

func DatabaseMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "db", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func initServeCommand(cfg config) *cobra.Command {
	return &cobra.Command{
		Use:     "serve",
		Aliases: []string{"serve"},
		Short:   "Serve voauth web server",
		Run: func(cmd *cobra.Command, args []string) {

			// initialize database
			db := infra.InitDb(cfg.DatabaseHost, cfg.DatabaseUser, cfg.DatabasePassword, cfg.DatabaseName, cfg.DatabasePort)

			// create and update database tables
			infra.AutoMigrateDB(db)

			// Instantiate new application
			application := app.NewApplication(app.ApplicationConfig{
				AuthSecretKey: cfg.AuthSecretKey,
			}, db, &infra.PasswordHasher{}, &infra.UUIDGenerator{}, &repository.UserRepository{}, &repository.NoteRepository{})

			// get api router
			apiRouter := api.GetApiRouter(application)

			mainRouter := chi.NewRouter()

			// add datbase middleware
			// TODO: might make sense to create a new read only db connection instance
			mainRouter.Use(DatabaseMiddleware(db))

			// mount api router on path /api
			mainRouter.Mount("/api", apiRouter)

			serverPort, _ := cmd.Flags().GetInt("port")

			slog.Info(fmt.Sprintf("Starting server on port %d", serverPort))
			http.ListenAndServe(fmt.Sprintf(":%d", serverPort), mainRouter)
		},
	}
}

func initRootCommand(cfg config) *cobra.Command {
	return &cobra.Command{
		Use:   "voauth",
		Short: "Voauth entry point command",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
			slog.Info("Trying about voauth")
		},
	}
}

func main() {
	cfg := config{}

	// populate cfg from env file
	if err := env.Parse(&cfg); err != nil {
		slog.Error("Error parsing environment variables", "error", err.Error())
		os.Exit(1)
	}

	// Initialize root command
	rootCmd := initRootCommand(cfg)

	// Initialzie serve subcommand with config
	serveCmd := initServeCommand(cfg)

	// Add serve sub command
	rootCmd.AddCommand(serveCmd)

	// Allow setting server port with port flag
	serveCmd.Flags().Int("port", cfg.Port, "Web server port")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
