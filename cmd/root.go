package cmd

import (
	"fmt"
	"os"

	"github.com/aponx/book-management/app/domain"
	"github.com/aponx/book-management/common"
	"github.com/aponx/book-management/driver"
	phttp "github.com/aponx/book-management/http"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	_bookDelivery "github.com/aponx/book-management/app/book/delivery"
	_bookUsecase "github.com/aponx/book-management/app/book/usecase"
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

	datajson, err := driver.NewLoadJson(conf.JSON.Data)

	if err != nil {
		log.Error().Err(err).Msg("DB Connection error")
		panic(err)
	}

	repos := wiringRepository(*datajson)

	usecase := wiringUsecase(*&repos)

	handler := wiringHttpHandler(conf, *usecase)

	router := phttp.Router(*handler)

	svr := phttp.NewServer(&conf.Server, router)

	svr.Start()

}

func wiringRepository(data []domain.Book) *domain.Repository {
	userRepo := _bookRepository.NewBookRepository(data)
	return &domain.Repository{
		UserRepo: userRepo,
	}
}

func wiringUsecase() *domain.Usecase {
	bookUsecase := _bookUsecase.NewBookUsecase()
	return &domain.Usecase{
		BookUsecase: bookUsecase,
	}
}

func wiringHttpHandler(conf *common.Config, usecase domain.Usecase) *domain.Delivery {
	bookHandler := _bookDelivery.NewBookHandler(conf, usecase.BookUsecase)

	return &domain.Delivery{
		BookDelivery: bookHandler,
	}

}
