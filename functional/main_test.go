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
	appName = "authz"

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
		Name:                 appName,
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if status != 0 {
		os.Exit(1)
	}
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
		// drop all tables in the database.
		dropAll()
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	api := &apiFeature{
		httpClient: &http.Client{},
		logger:     logger,
	}

	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// truncate all tables in the database.
		truncateAll()

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

	ctx.Step(`^I authenticate with username "([^"]*)" and password "([^"]*)"$`, api.iAuthenticateWithUsernameAndPassword)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)"$`, api.iSendRequestTo)
	ctx.Step(`^I send "(GET|POST|PUT|DELETE)" request to "([^"]*)" with payload:$`, api.iSendRequestToWithPayload)
	ctx.Step(`^the response code should be (\d+)$`, api.theResponseCodeShouldBe)
	ctx.Step(`^the response should match json:$`, api.theResponseShouldMatchJSON)
}

// tableNames return all table names in the database.
func tableNames() ([]string, error) {
	var tx *gorm.DB

	cfg, _ := service.SysConfig().GetDatabase(context.Background())
	dbname := cfg.Dbname
	switch cfg.Driver {
	case "sqlite":
		tx = db.Table("sqlite_master").
			Select("tbl_name").
			Where("type = ?", "table").
			Where("tbl_name NOT LIKE ?", "sqlite_%")
	case "postgres":
		dbname = "public"
		fallthrough
	default:
		tx = db.Table("information_schema.tables").
			Select("table_name").
			Where("table_type = ?", "BASE TABLE").
			Where("table_schema = ?", dbname)
	}

	var names []string
	err := tx.Scan(&names).Error
	if err != nil {
		return nil, err
	}
	return names, nil
}

// truncateAll truncate all tables in the database.
func truncateAll() error {
	return db.Transaction(func(tx *gorm.DB) error {
		names, err := tableNames()
		if err != nil {
			return err
		}
		if tx.Name() == "mysql" {
			if err = tx.Exec("SET FOREIGN_KEY_CHECKS = 0").Error; err != nil {
				return err
			}
			defer tx.Exec("SET FOREIGN_KEY_CHECKS = 1")
		}
		for _, name := range names {
			raw := "TRUNCATE TABLE " + name
			if tx.Name() == "sqlite" {
				raw = "DELETE FROM " + name
			}
			if err = tx.Exec(raw).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

// dropAll drop all tables in the database.
func dropAll() error {
	return db.Transaction(func(tx *gorm.DB) error {
		names, err := tableNames()
		if err != nil {
			return err
		}
		for _, name := range names {
			raw := "DROP TABLE " + name
			if err = tx.Exec(raw).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
