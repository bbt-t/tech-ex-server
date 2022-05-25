package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"tech-ex-server/internal/app/storage"
	"time"
)

type APIServer struct {
	config  *Config
	logger  *logrus.Logger
	router  *mux.Router
	storage *storage.Storage
}

func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *APIServer) StartServer() error {
	if err := a.makeLogger(); err != nil {
		return err
	}
	a.makeRouter()
	if err := a.makeStorage(); err != nil {
		return err
	}
	a.logger.Infoln("---> START API SERVER! <---")

	return http.ListenAndServe(a.config.BinAddr, a.router)
}

func (a *APIServer) makeLogger() error {
	level, err := logrus.ParseLevel(a.config.LogLevel)
	if err != nil {
		return err
	}

	a.logger.SetLevel(level)
	return nil
}

func (a *APIServer) makeRouter() {
	a.router.HandleFunc("/api", a.helloHandler())
}

func (a *APIServer) makeStorage() error {
	st := storage.New(a.config.Storage)
	if err := st.Open(); err != nil {
		return err
	}
	a.storage = st
	return nil
}

func (a *APIServer) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rg, err := a.storage.Item().SelectItem()
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		if _, err := w.Write(rg); err != nil {
			log.Fatal(err)
		}
	}
}

type ResponseJsonData []struct {
	Symbol         string  `json:"symbol"`
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type dataToAdd struct {
	Price24H       float64 `json:"price_24h"`
	Volume24H      float64 `json:"volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

func RequestToExternalSource(url string) *http.Response {
	client := http.Client{}
	client.Timeout = time.Second * 5

	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func GetDataFromResp(response *http.Response) ResponseJsonData {
	var result ResponseJsonData

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	return result
}

func EditedData() map[string]dataToAdd {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}
	editedData := map[string]dataToAdd{}

	response := RequestToExternalSource(os.Getenv("RQ_API"))
	result := GetDataFromResp(response)

	for _, val := range result {
		editedData[val.Symbol] = dataToAdd{
			Price24H:       val.Price24H,
			Volume24H:      val.Volume24H,
			LastTradePrice: val.LastTradePrice,
		}
	}
	return editedData
}
