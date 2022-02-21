package app

import (
	"context"
	"log"
	_ "net/http/pprof"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"

	"maranatha_web/internal/config"
	"maranatha_web/internal/controllers"
	"maranatha_web/internal/controllers/token"
	"maranatha_web/internal/datasources/fcm_client"
	"maranatha_web/internal/datasources/filestorage"
	mailClient "maranatha_web/internal/datasources/mail"
	redisDb "maranatha_web/internal/datasources/redis"
	"maranatha_web/internal/logger"
	"maranatha_web/internal/repository"
	"maranatha_web/internal/services"
)

var app config.AppConfig

//const jwtKey = "*#*Johnte2536290248"

var (
	serviceName  = os.Getenv("SERVICE_NAME")
	signozToken  = os.Getenv("SIGNOZ_ACCESS_TOKEN")
	collectorURL = os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	insecure     = os.Getenv("INSECURE_MODE")
)

func initTracer() func(context.Context) error {

	headers := map[string]string{
		"signoz-access-token": signozToken,
	}

	secureOption := otlptracegrpc.WithTLSCredentials(credentials.NewClientTLSFromCert(nil, ""))
	if len(insecure) > 0 {
		secureOption = otlptracegrpc.WithInsecure()
	}

	exporter, err := otlptrace.New(
		context.Background(),
		otlptracegrpc.NewClient(
			secureOption,
			otlptracegrpc.WithEndpoint(collectorURL),
			otlptracegrpc.WithHeaders(headers),
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	resources, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
			attribute.String("library.language", "go"),
		),
	)
	if err != nil {
		log.Print("Could not set resources: ", err)
	}

	otel.SetTracerProvider(
		sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithSpanProcessor(sdktrace.NewBatchSpanProcessor(exporter)),
			sdktrace.WithSyncer(exporter),
			sdktrace.WithResource(resources),
		),
	)
	return exporter.Shutdown
}
func StartApplication() {

	///Setup signoz
	//cleanup := initTracer()
	//defer cleanup(context.Background())

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error  .env file")
	}
	jwtKey := os.Getenv("JwtSecret")
	tokenLifeTime := os.Getenv("TLT")

	//configure logger
	zl := logger.GetLogger()
	app.ErrorLog = zl
	app.InfoLog = zl
	_ = config.NewAppConfig(zl, zl)

	//Create the new token maker
	//
	_, err = token.NewJWTMaker(jwtKey)
	if err != nil {
		return
	}
	//
	//Token lifetime
	intVar, err := strconv.Atoi(tokenLifeTime)
	if err != nil {
		log.Panicln(err)
	}
	app.TokenLifeTime = intVar
	//Get database connection
	db := repository.GetDatabaseConnection()
	dao := repository.NewPostgresRepo(db, &app)

	//
	//Get fcm connection
	messagingService := fcm_client.GetFcmConnection()

	//
	//
	//Get file storage connection
	connection, bucketName, minioErr := filestorage.GetMinioConnection()
	if minioErr != nil {
		log.Panicln(minioErr)
	}
	filestorage.NewMinoRepo(connection, &app)
	log.Printf("mino endpoint is %s", connection.EndpointURL())

	fs := filestorage.MinioRepo{
		App:          &app,
		MinioStorage: connection,
	}
	app.StorageURL = connection.EndpointURL()
	app.StorageBucket = bucketName

	//connect to mail server
	mailServer := mailClient.GetMailServer()

	//initiate services
	booksService := services.NewBookService(dao)
	eventsService := services.NewEventsService(dao)
	genresService := services.NewGenreService(dao)
	jobsService := services.NewJobsService(dao)
	newsService := services.NewNewsService(dao, &app)
	podcastService := services.NewPodcastService(dao, &app)

	partnersService := services.NewChurchPartnersService(dao, &app)
	prayerRequestService := services.NewPrayerRequestService(dao)
	sermonServices := services.NewSermonService(dao, &app)
	testimonyService := services.NewTestimoniesService(dao)
	usersService := services.NewUsersService(dao)
	volunteerJobService := services.NewVolunteerChurchJobService(dao)
	dailyVerseService := services.NewDailyVerseService()
	mailService := services.NewMailService(mailServer)
	fcmService := services.NewFcmService(messagingService)

	redisDb.GetRedisClient()

	allServices := controllers.NewRepo(
		mailService,
		fcmService,
		&app,
		booksService,
		&fs,
		dailyVerseService,
		eventsService,
		genresService,
		jobsService,
		newsService,
		podcastService,
		partnersService,
		prayerRequestService,
		sermonServices,
		testimonyService,
		usersService,
		volunteerJobService)

	controllers.NewHandlers(allServices)

	r := SetupRouter()

	err = r.Run("localhost:8090")
	//err = r.Run("192.168.2.70:8090")
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
