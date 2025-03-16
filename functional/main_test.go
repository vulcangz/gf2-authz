//go:build functional
// +build functional

package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"testing"
	"time"

	_ "time/tzdata"

	"github.com/vulcangz/gf2-authz/internal/lib/database"
	_ "github.com/vulcangz/gf2-authz/internal/logic"
	"github.com/vulcangz/gf2-authz/internal/service"

	"github.com/vulcangz/gf2-authz/internal/fixtures"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/spf13/pflag"
	"gorm.io/gorm"
)

var (
	appName     string
	db          *gorm.DB
	logger      *slog.Logger
	initializer fixtures.Initializer
	opts        = godog.Options{
		Format: "pretty",
		Paths:  []string{"features"},
		Output: colors.Colored(os.Stdout),
	}

	initialTime = time.Date(2100, time.January, 1, 1, 0, 0, 0, time.UTC)
	// ErrUnsupportedDriver is returned when specified database driver does not exists.
	ErrUnsupportedDriver = errors.New("unsupported database driver")
)

func init() {
	godog.BindCommandLineFlags("godog.", &opts)
	appName = g.Cfg().MustGet(context.Background(), "app.name").String()

	db = database.GetDatabase(context.Background())
	logger = slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	pflag.Parse()
	opts.Paths = pflag.Args()

	testServer()

	cfg, _ := service.SysConfig().GetUser(ctx)
	initializer = fixtures.NewInitializer(cfg)

	status := godog.TestSuite{
		Name:                appName,
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	if status != 0 {
		os.Exit(1)
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{
		httpClient: &http.Client{},
		logger:     logger,
	}

	//mysql
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		cfg, _ := service.SysConfig().GetDatabase(ctx)

		var truncateSQLSlice []string
		db.Raw("SELECT CONCAT(\"TRUNCATE TABLE `\", t.TABLE_NAME, '`;') FROM information_schema.`TABLES` t WHERE t.TABLE_SCHEMA =?", cfg.Dbname).Scan(&truncateSQLSlice)
		truncateSQLSlice = append(truncateSQLSlice, "SET FOREIGN_KEY_CHECKS = 1;")

		if err := db.Exec(`
			SET FOREIGN_KEY_CHECKS = 0;
			`).Error; err != nil {
			logger.Error("Unable to set foreign key checks: ", slog.Any("err", err))
		}

		for _, v := range truncateSQLSlice {
			if err := db.Exec(v).Error; err != nil {
				logger.Error("Unable to truncate tables: ", slog.Any("err", err))
				return ctx, err
			}
		}

		if err := api.reset(sc); err != nil {
			logger.Error("Cannot reset api: ", slog.Any("err", err))
			return ctx, err
		}

		if err := initializer.Initialize(); err != nil {
			logger.Error("Cannot initialize fixtures:", slog.Any("err", err))
			return ctx, err
		}

		return ctx, nil
	})

	ctx.After(func(ctx context.Context, sc *godog.Scenario, err error) (context.Context, error) {
		return ctx, nil
	})

	ctx.Step(`^I wait "([^"]*)"$`, func(value string) error {
		duration, err := time.ParseDuration(value)
		if err != nil {
			return err
		}

		time.Sleep(duration)
		return nil
	})
	logger.Info("马上进入到step……")
	ctx.Step(`^I authenticate with username "([^"]*)" and password "([^"]*)"$`, api.iAuthenticateWithUsernameAndPassword)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.iSendRequestTo)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)" with payload:$`, api.iSendRequestToWithPayload)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
}
