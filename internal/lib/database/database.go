package database

import (
	"context"
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmode"
	"github.com/vulcangz/gf2-authz/internal/lib/ctime"
	"github.com/vulcangz/gf2-authz/internal/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB

	// ErrUnsupportedDriver is returned when specified database driver does not exists.
	ErrUnsupportedDriver = errors.New("unsupported database driver")
)

func Initialize(ctx context.Context, clock ctime.Clock) (db *gorm.DB, err error) {
	var dialector gorm.Dialector

	logger := gormLogger.Default.LogMode(gormLogger.Silent)
	dbConfig, _ := GetDatabaseConfig(ctx)
	switch dbConfig.Driver {
	case entity.DriverMysql:
		dialector = mysql.Open(dbConfig.MysqlDSN())
	case entity.DriverSqlite:
		dialector = sqlite.Open(dbConfig.SqliteDSN())
	case entity.DriverPostgres:
		dialector = postgres.New(postgres.Config{DSN: dbConfig.PostgresDSN()})
	default:
		return nil, ErrUnsupportedDriver
	}

	db, err = gorm.Open(dialector, &gorm.Config{
		Logger:  logger,
		NowFunc: clock.Now,
	})
	if err != nil {
		return nil, err
	}

	if dbConfig.Driver == entity.DriverSqlite {
		checkErr(ctx, logger, db.AutoMigrate(entity.Action{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Attribute{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Audit{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Client{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.CompiledPolicy{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Policy{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Principal{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Stats{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Resource{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Role{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.Token{}))
		checkErr(ctx, logger, db.AutoMigrate(entity.User{}))
	} else {
		db.AutoMigrate(entity.Action{}, entity.Attribute{}, entity.Audit{},
			entity.Client{}, entity.CompiledPolicy{}, entity.Policy{}, entity.Principal{},
			entity.Stats{}, entity.Resource{}, entity.Role{}, entity.Token{}, entity.User{},
		)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, _ := db.DB()
	// Ping
	err = sqlDB.Ping()
	if err == nil {
		g.Log().Line().Infof(ctx, "%s database is alive!\n", dbConfig.Driver)
	}

	return dbInstance, nil
}

func GetDatabase(ctx context.Context) *gorm.DB {
	// 获取dbInstance，这个连接使用时自动生成副本，当前使用的永远是副本
	if dbInstance == nil {
		var clock ctime.Clock
		if gmode.IsTesting() {
			clock = ctime.NewStaticClock()
		} else {
			clock = ctime.NewClock()
		}
		dbInstance, _ = Initialize(ctx, clock)
	}
	return dbInstance
}

func GetTestDatabase(ctx context.Context) *gorm.DB {
	// 获取dbInstance，这个连接使用时自动生成副本，当前使用的永远是副本
	if dbInstance == nil {
		clock := ctime.NewStaticClock()
		dbInstance, _ = Initialize(ctx, clock)
	}
	return dbInstance
}

// GetDatabase get database configuration options
func GetDatabaseConfig(ctx context.Context) (conf *entity.DatabaseConfig, err error) {
	val, err := g.Cfg().GetWithEnv(ctx, "database")
	if err != nil {
		model := &entity.DatabaseConfig{}
		conf = model.DefaultConfig()
		return
	}
	if val != nil {
		err = val.Scan(&conf)
		return
	}

	// container env setup for sqlite
	val = g.Cfg().MustGetWithEnv(ctx, "database.driver", "sqlite")
	driver := val.String()
	if driver == "sqlite" {
		name := g.Cfg().MustGetWithEnv(ctx, "database.dbname", ":memory:")
		conf = &entity.DatabaseConfig{
			Driver: driver,
			Dbname: name.String(),
		}
		return
	}

	return
}

func checkErr(ctx context.Context, logger gormLogger.Interface, err error) {
	if err != nil {
		logger.Error(ctx, "Cannot migrate database", err)
	}
}
