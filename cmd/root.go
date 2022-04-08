package cmd

import (
	"fmt"
	"os"
	"umu/golang-api/app/domain"
	"umu/golang-api/common"
	"umu/golang-api/driver"
	phttp "umu/golang-api/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/gorp.v2"

	_userDelivery "umu/golang-api/app/user/delivery/http"
	_userRepository "umu/golang-api/app/user/repository"
	_userUsecase "umu/golang-api/app/user/usecase"
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "golang-api",
	Long:  `golang-api`,
	Run: func(cmd *cobra.Command, args []string) {
		startHttpServer()
	},
}

func init() {
	cobra.OnInitialize()
}

func startHttpServer() {
	conf, err := common.NewConfig()
	if err != nil {
		panic(err)
	}

	dbInstance, err := driver.NewPostgreDatabase(conf.DB)

	if err != nil {
		log.Error().Err(err).Msg("DB Connection error")
		panic(err)
	}

	initDb(dbInstance)

	repos := wiringRepository(dbInstance)

	usecase := wiringUsecase(*repos)

	handler := wiringHttpHandler(conf, *usecase)

	router := phttp.Router(*handler)

	svr := phttp.NewServer(&conf.Server, router)

	svr.Start()

}

func initDb(db *gorp.DbMap) {
	db.AddTableWithName(domain.User{}, "users").SetKeys(true, "Id")
}

func wiringRepository(db *gorp.DbMap) *domain.Repository {
	userRepo := _userRepository.NewUserRepository(db)
	return &domain.Repository{
		UserRepo: userRepo,
	}
}

func wiringUsecase(repos domain.Repository) *domain.Usecase {
	userUsecase := _userUsecase.NewUserUsecase(repos.UserRepo)
	return &domain.Usecase{
		UserUsecase: userUsecase,
	}
}

func wiringHttpHandler(conf *common.Config, usecase domain.Usecase) *domain.Delivery {
	userHandler := _userDelivery.NewUserHandler(conf, usecase.UserUsecase)

	return &domain.Delivery{
		UserDelivery: userHandler,
	}

}
