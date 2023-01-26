package cmd

import (
	"RegLog/app"
	"RegLog/app/response"
	"RegLog/config"
	infraPostgres "RegLog/infra/postgres"
	"RegLog/repo"
	"RegLog/service"
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var srvrInfo *response.ServerInfo
var cfgPostgres *response.Postgres

// srvCmd is the serve sub command to start the api server
var srvCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve serves the api server",
	RunE:  serve,
}

func init() {
	srvCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "", "config file path")
}

func serve(cmd *cobra.Command, args []string) error {

	cfgPostgres, srvrInfo = config.GetPostgres(cfgPath)
	db, err := infraPostgres.PostNew(*cfgPostgres)

	signBytes, _ := ioutil.ReadFile("./app.rsa")
	PrivateKey, _ := jwt.ParseRSAPrivateKeyFromPEM(signBytes)

	verifyBytes, _ := ioutil.ReadFile("./app.rsa.pub")
	PublicKey, _ := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)

	if err != nil {
		fmt.Println("problem database connection")
		return err
	}
	brandRepo := repo.NewBrand(db, PrivateKey, PublicKey)
	svc := service.NewBrand(brandRepo)

	errChan := make(chan error)
	go func() {
		if err := startApiServer(svc); err != nil {
			errChan <- err
		}
	}()
	return <-errChan
	//startApiServer(svc, cfgPostgres)
	//return nil
}

func startApiServer(svc service.BrandService) error {

	brndsCtrl := app.NewBrandsController(svc)

	r := chi.NewRouter()

	r.Mount("/api/v1", app.NewRouter(brndsCtrl))

	srvr := http.Server{
		Addr:    getAddressFromHostAndPort(cfgPostgres.Host, cfgPostgres.AppPort),
		Handler: r,
		//ErrorLog: logger.DefaultErrLogger,
		//WriteTimeout: cfg.WriteTimeout,
		//ReadTimeout:  cfg.ReadTimeout,
		ReadTimeout:       time.Duration(srvrInfo.ReadTimeout) * time.Second,
		WriteTimeout:      time.Duration(srvrInfo.WriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(srvrInfo.IdleTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(srvrInfo.ReadHeaderTimeout) * time.Second,
	}

	return ManageServer(&srvr, time.Duration(srvrInfo.GracePeriod)*time.Second)
}

func ManageServer(srvr *http.Server, gracePeriod time.Duration) error {
	errCh := make(chan error)

	sigs := []os.Signal{syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, os.Interrupt}

	graceful := func() error {
		log.Println("Shutting down server gracefully with in", gracePeriod)
		log.Println("To shutdown immedietly press again")

		ctx, cancel := context.WithTimeout(context.Background(), gracePeriod)
		defer cancel()

		return srvr.Shutdown(ctx)
	}

	forced := func() error {
		log.Println("Shutting down server forcefully")
		return srvr.Close()
	}

	go func() {
		log.Println("Starting server on", srvr.Addr)
		if err := srvr.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	go func() {
		errCh <- HandleSignals(sigs, graceful, forced)
	}()

	return <-errCh
}

func HandleSignals(sigs []os.Signal, gracefulHandler, forceHandler func() error) error {
	sigCh := make(chan os.Signal)
	errCh := make(chan error, 1)

	signal.Notify(sigCh, sigs...)
	defer signal.Stop(sigCh)

	grace := true
	for {
		select {
		case err := <-errCh:
			return err
		case <-sigCh:
			if grace {
				grace = false
				go func() {
					errCh <- gracefulHandler()
				}()
			} else if forceHandler != nil {
				err := forceHandler()
				errCh <- err
			}
		}
	}
}

func getAddressFromHostAndPort(host string, port int) string {
	addr := host
	if port > 0 {
		addr = addr + ":" + strconv.Itoa(port)
	}
	return addr
}
