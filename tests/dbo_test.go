package tests

import (
	"os"
	"testing"

	"github.com/Cyberpull/dbo"
	"github.com/Cyberpull/dbo/clauses"
	"github.com/Cyberpull/dbo/pagination"
	"github.com/Cyberpull/dbo/tests/ctes"
	"github.com/Cyberpull/dbo/tests/models"
	"github.com/Cyberpull/dbo/tests/seeders"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DBOTestSuite struct {
	suite.Suite

	inst dbo.Instance
}

func (x *DBOTestSuite) SetupSuite() {
	var err error

	connector := dbo.NewConnector(dbo.Options{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		DBName:   env("DB_DATABASE", "dbo_test"),
		Username: env("DB_USERNAME", "dbo_test"),
		Password: env("DB_PASSWORD", "nonstop"),
		Config: &gorm.Config{
			CreateBatchSize: 3000,
			Logger:          logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	})

	x.inst, err = connector.Connect()
	require.NoError(x.T(), err)

	require.NoError(x.T(), x.inst.Migrate(
		&models.Car{},
		&models.House{},

		&models.Person{},
		&models.PersonStatus{},
		&models.PersonLocation{},
	))

	require.NoError(x.T(), x.inst.Seed(
		&seeders.Car{},
		&seeders.House{},

		&seeders.Person{},
		&seeders.PersonStatus{},
		&seeders.PersonLocation{},
	))
}

func (x *DBOTestSuite) TearDownSuite() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	// Drop Association Tables
	require.NoError(x.T(), db.Migrator().DropTable(
		"person_car",
		"person_house",
	))

	// Drop Tables
	require.NoError(x.T(), db.Migrator().DropTable(
		&models.Car{},
		&models.House{},

		&models.Person{},
		&models.PersonStatus{},
		&models.PersonLocation{},
	))
}

func (x *DBOTestSuite) TestPersons() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	persons := []*models.Person{}

	err = db.Find(&persons).Error
	require.NoError(x.T(), err)

	count := len(persons)
	require.Equal(x.T(), 2, count)
}

func (x *DBOTestSuite) TestPersonsPagination() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	var engine pagination.Engine[*pagination.Pagination[*models.Person], *models.Person]

	records, err := engine.Paginate(db, pagination.Options{Limit: 10})
	require.NoError(x.T(), err)

	require.Equal(x.T(), 10, records.PerPage)
	require.Equal(x.T(), uint64(1), records.CurrentPage)
	require.Equal(x.T(), uint64(2), records.Total)

	count := len(records.Data)
	require.Equal(x.T(), 2, count)
}

func (x *DBOTestSuite) TestApprovedPersonsPagination() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	db = db.Clauses(clauses.With{CTE: []clauses.CTE{
		&ctes.PersonStatusCTE{},
	}})

	db = db.Joins("INNER JOIN `ps` ON ps.personId = `person`.id AND `ps`.row = 1")
	db = db.Where("`ps`.name = ?", "APPROVED")

	var engine pagination.Engine[*pagination.Pagination[*models.Person], *models.Person]

	records, err := engine.Paginate(db, pagination.Options{Limit: 10})
	require.NoError(x.T(), err)

	require.Equal(x.T(), 10, records.PerPage)
	require.Equal(x.T(), uint64(1), records.CurrentPage)
	require.Equal(x.T(), uint64(1), records.Total)

	count := len(records.Data)
	require.Equal(x.T(), 1, count)
}

func (x *DBOTestSuite) TestSuspendedPersonsPagination() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	db = db.Clauses(clauses.With{CTE: []clauses.CTE{
		&ctes.PersonStatusCTE{},
	}})

	db = db.Joins("INNER JOIN `ps` ON ps.personId = `person`.id AND `ps`.row = 1")
	db = db.Where("`ps`.name = ?", "SUSPENDED")

	var engine pagination.Engine[*pagination.Pagination[*models.Person], *models.Person]

	records, err := engine.Paginate(db, pagination.Options{Limit: 10})
	require.NoError(x.T(), err)

	require.Equal(x.T(), 10, records.PerPage)
	require.Equal(x.T(), uint64(1), records.CurrentPage)
	require.Equal(x.T(), uint64(1), records.Total)

	count := len(records.Data)
	require.Equal(x.T(), 1, count)
}

func (x *DBOTestSuite) TestApprovedPersons() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	persons := []*models.Person{}

	db = db.Clauses(clauses.With{CTE: []clauses.CTE{
		&ctes.PersonStatusCTE{},
	}})

	db = db.Joins("INNER JOIN `ps` ON ps.personId = `person`.id AND `ps`.row = 1")
	db = db.Where("`ps`.name = ?", "APPROVED")

	err = db.Find(&persons).Error
	require.NoError(x.T(), err)

	count := len(persons)
	require.Equal(x.T(), 1, count)
}

func (x *DBOTestSuite) TestSuspendedPersons() {
	db, err := x.inst.Get()
	require.NoError(x.T(), err)

	persons := []*models.Person{}

	db = db.Clauses(clauses.With{CTE: []clauses.CTE{
		&ctes.PersonStatusCTE{},
	}})

	db = db.Joins("INNER JOIN `ps` ON ps.personId = `person`.id AND `ps`.row = 1")
	db = db.Where("`ps`.name = ?", "SUSPENDED")

	err = db.Find(&persons).Error
	require.NoError(x.T(), err)

	count := len(persons)
	require.Equal(x.T(), 1, count)
}

// ===============================

func env(key string, defaultValue ...string) (value string) {
	value = os.Getenv(key)

	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}

	return
}

// ===============================

func TestDBO(t *testing.T) {
	suite.Run(t, new(DBOTestSuite))
}
