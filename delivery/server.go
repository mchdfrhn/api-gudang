package delivery

import (
	"barang/config"
	"barang/delivery/controller"
	"barang/repository"
	"barang/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gudangUc usecase.GudangUsecase
	barangUc usecase.BarangUsecase

	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	controller.NewGudangController(s.gudangUc, s.engine).Route()
	controller.NewBarangController(s.barangUc, s.engine).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("Server Error : ", err.Error())
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatal("Config : ", err.Error())
	}

	db, err := config.NewDbConnection(cfg)

	if err != nil {
		log.Fatal("DB Connect : ", err.Error())
	}

	// Gudang
	gudangRepo := repository.NewGudangRepository(db.Conn())
	gudangUc := usecase.NewGudangUsecase(gudangRepo)

	// Barang
	barangRepo := repository.NewBarangRepository(db.Conn())
	barangUc := usecase.NewBarangUsecase(barangRepo)


	// Gin Engine
	engine := gin.Default()

	return &Server{
		gudangUc: gudangUc,
		barangUc: barangUc,

		engine: engine,
		host:   ":8080",
	}
}
