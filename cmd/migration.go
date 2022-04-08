package cmd

import (
	"fmt"
	"os"
	"time"

	"umu/golang-api/driver"

	"umu/golang-api/common"

	"github.com/rs/zerolog/log"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/cobra"
)

var migrateUpCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate Up DB",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := common.NewConfig()
		if err != nil {
			log.Error().Msgf("Config error | %v", err)
			panic(err)
		}
		mSource := getMigrateSource()

		doMigrate(conf.DB, mSource, migrate.Up)
	},
}

var migrateDownCmd = &cobra.Command{
	Use:   "migratedown",
	Short: "Migrate Down DB",
	Long:  `Please you know what are you doing by using this command`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := common.NewConfig()
		if err != nil {
			log.Error().Msgf("Config error | %v", err)
			panic(err)
		}
		mSource := getMigrateSource()

		doMigrate(conf.DB, mSource, migrate.Down)
	},
}

var migrateNewCmd = &cobra.Command{
	Use:   "migratenew [migration name]",
	Short: "Create new migration file",
	Long:  `Create new migration file on folder migrations/sql with timestamp as prefix`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mDir := "migrations/sql/"

		createMigrationFile(mDir, args[0])
	},
}

func init() {
	RootCmd.AddCommand(migrateUpCmd)
	RootCmd.AddCommand(migrateDownCmd)
	RootCmd.AddCommand(migrateNewCmd)
}

func getMigrateSource() migrate.FileMigrationSource {
	source := migrate.FileMigrationSource{
		Dir: "migrations/sql",
	}

	return source
}

func doMigrate(config common.DB, mSource migrate.FileMigrationSource, direction migrate.MigrationDirection) error {
	db, err := driver.NewPostgreDatabase(config)

	defer db.Db.Close()

	total, err := migrate.Exec(db.Db, "postgres", mSource, direction)
	if err != nil {
		log.Printf("Fail migration | %v", err)
		return err
	}

	log.Printf("Migrate Success, total migrated: %d", total)
	return nil
}

func createMigrationFile(mDir string, mName string) error {
	var migrationContent = `-- +migrate Up
 		-- SQL in section 'Up' is executed when this migration is applied
 		-- [your SQL script here]

		 -- +migrate Down
 		-- SQL section 'Down' is executed when this migration is rolled back
 		-- [your SQL script here]
 	`
	filename := fmt.Sprintf("%d_%s.sql", time.Now().Unix(), mName)
	filepath := fmt.Sprintf("%s%s", mDir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		log.Printf("Error create migration file | %v", err)
		return err
	}
	defer f.Close()

	f.WriteString(migrationContent)
	f.Sync()

	log.Printf("New migration file has been created: %s)", filepath)
	return nil
}
