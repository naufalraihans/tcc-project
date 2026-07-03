package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"tcc-itpln/backend/config"
	"tcc-itpln/backend/internal/handler"
	"tcc-itpln/backend/internal/middleware"
	"tcc-itpln/backend/internal/repository"
	"tcc-itpln/backend/internal/usecase"
	"tcc-itpln/backend/pkg/midtrans"
	"tcc-itpln/backend/pkg/utils"
)

func New(cfg config.Config, db *pgxpool.Pool) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS(cfg.CORSAllowedOrigins))

	r.GET("/health", func(c *gin.Context) {
		if err := db.Ping(c.Request.Context()); err != nil {
			utils.Err(c, 500, "INTERNAL_ERROR", "database unreachable")
			return
		}
		utils.OK(c, gin.H{"status": "ok"})
	})

	profileRepo := repository.NewProfileRepository(db)
	txRepo := repository.NewTransaksiRepository(db)
	gamRepo := repository.NewGamifikasiRepository(db)
	pengRepo := repository.NewPengumumanRepository(db)
	midClient := midtrans.New(cfg.MidtransServerKey, cfg.MidtransIsProduction)

	authMw := middleware.Auth(cfg.SupabaseJWTSecret, profileRepo.Role)
	adminMw := middleware.RequireAdmin()

	topikH := handler.NewTopikHandler(usecase.NewTopikUsecase(repository.NewTopikRepository(db)))
	instrukturH := handler.NewInstrukturHandler(usecase.NewInstrukturUsecase(repository.NewInstrukturRepository(db)))
	kelasH := handler.NewKelasHandler(usecase.NewKelasUsecase(repository.NewKelasRepository(db)))
	authH := handler.NewAuthHandler(usecase.NewAuthUsecase(profileRepo, repository.NewAuthRepository(db), gamRepo, cfg.SupabaseJWTSecret))
	pendaftaranH := handler.NewPendaftaranHandler(usecase.NewPendaftaranUsecase(repository.NewPendaftaranRepository(db), txRepo, profileRepo, gamRepo, midClient))
	konsultasiH := handler.NewKonsultasiHandler(usecase.NewKonsultasiUsecase(repository.NewKonsultasiRepository(db), gamRepo))
	transaksiH := handler.NewTransaksiHandler(usecase.NewTransaksiUsecase(txRepo, gamRepo, midClient))
	sertifikatH := handler.NewSertifikatHandler(usecase.NewSertifikatUsecase(repository.NewSertifikatRepository(db)))
	materiH := handler.NewMateriHandler(usecase.NewMateriUsecase(repository.NewMateriRepository(db), gamRepo))
	gamifikasiH := handler.NewGamifikasiHandler(usecase.NewGamifikasiUsecase(gamRepo, pengRepo, profileRepo))
	pengumumanH := handler.NewPengumumanHandler(usecase.NewPengumumanUsecase(pengRepo))

	api := r.Group("/api/v1")

	api.GET("/topik", topikH.List)
	api.GET("/topik/:slug", topikH.Detail)
	api.GET("/instruktur", instrukturH.List)
	api.GET("/instruktur/:id", instrukturH.Detail)
	api.GET("/kelas", kelasH.List)
	api.GET("/kelas/:slug", kelasH.Detail)
	api.GET("/sertifikat/verify/:nomor", sertifikatH.Verify)
	api.GET("/pengumuman", pengumumanH.ListActive)
	api.POST("/webhook/midtrans", transaksiH.Webhook)
	api.POST("/auth/register", authH.Register)
	api.POST("/auth/login", authH.Login)

	auth := api.Group("")
	auth.Use(authMw)
	{
		auth.GET("/auth/me", authH.Me)
		auth.PUT("/auth/profile", authH.UpdateProfile)
		auth.GET("/kelas/:slug/materi", materiH.ListForUser)
		auth.POST("/pendaftaran", pendaftaranH.Daftar)
		auth.GET("/pendaftaran/saya", pendaftaranH.ListSaya)
		auth.GET("/pendaftaran/saya/:id", pendaftaranH.DetailSaya)
		auth.POST("/konsultasi", konsultasiH.Create)
		auth.GET("/konsultasi/saya", konsultasiH.ListSaya)
		auth.GET("/transaksi/saya", transaksiH.ListSaya)
		auth.GET("/sertifikat/saya", sertifikatH.ListSaya)
		auth.GET("/me/dashboard", gamifikasiH.Beranda)
		auth.GET("/me/progress", gamifikasiH.Progress)
		auth.GET("/me/misi", gamifikasiH.Misi)
	}

	admin := api.Group("/admin")
	admin.Use(authMw, adminMw)
	{
		admin.POST("/topik", topikH.Create)
		admin.PUT("/topik/:id", topikH.Update)
		admin.DELETE("/topik/:id", topikH.Delete)

		admin.POST("/instruktur", instrukturH.Create)
		admin.PUT("/instruktur/:id", instrukturH.Update)
		admin.DELETE("/instruktur/:id", instrukturH.Delete)

		admin.POST("/kelas", kelasH.Create)
		admin.PUT("/kelas/:id", kelasH.Update)
		admin.PATCH("/kelas/:id/status", kelasH.UpdateStatus)
		admin.DELETE("/kelas/:id", kelasH.Delete)
		admin.POST("/kelas/:id/materi", materiH.Create)

		admin.PUT("/materi/:id", materiH.Update)
		admin.DELETE("/materi/:id", materiH.Delete)

		admin.GET("/pendaftaran", pendaftaranH.ListAll)
		admin.PATCH("/pendaftaran/:id/status", pendaftaranH.UpdateStatus)

		admin.GET("/konsultasi", konsultasiH.ListAll)
		admin.GET("/konsultasi/:id", konsultasiH.Detail)
		admin.PATCH("/konsultasi/:id", konsultasiH.Respond)

		admin.POST("/sertifikat", sertifikatH.Issue)

		admin.GET("/transaksi", transaksiH.ListAll)
		admin.PATCH("/transaksi/:id/status", transaksiH.UpdateStatus)

		admin.GET("/misi", gamifikasiH.ListMisi)
		admin.POST("/misi", gamifikasiH.CreateMisi)
		admin.PUT("/misi/:id", gamifikasiH.UpdateMisi)
		admin.DELETE("/misi/:id", gamifikasiH.DeleteMisi)

		admin.GET("/pengumuman", pengumumanH.ListAll)
		admin.POST("/pengumuman", pengumumanH.Create)
		admin.PUT("/pengumuman/:id", pengumumanH.Update)
		admin.DELETE("/pengumuman/:id", pengumumanH.Delete)
	}

	return r
}
