package controllers

import (
	"maranatha_web/internal/config"
	"maranatha_web/internal/datasources/filestorage"
	"maranatha_web/internal/services"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App                  *config.AppConfig
	MinoStorage          *filestorage.MinioRepo
	dailyVerse           services.DailVerseService
	bookService          services.BooksService
	eventsService        services.EventsService
	genreService         services.GenreService
	jobService           services.JobsService
	newsService          services.NewsService
	partnersService      services.ChurchPartnersService
	prayerRequestService services.PrayerRequestService
	sermonService        services.SermonService
	testimonyService     services.TestimoniesService
	userServices         services.UsersService
	volunteerService     services.VolunteerChurchJobService
}

// NewRepo creates a new repository
func NewRepo(
	a *config.AppConfig,
	booksService services.BooksService,
	m *filestorage.MinioRepo,
	dailyVerse services.DailVerseService,
	eventsService services.EventsService,
	genreService services.GenreService,
	jobService services.JobsService,
	newsService services.NewsService,
	partnersService services.ChurchPartnersService,
	prayerRequestService services.PrayerRequestService,
	sermonService services.SermonService,
	testimonyService services.TestimoniesService,
	userServices services.UsersService,
	volunteerService services.VolunteerChurchJobService,
) *Repository {
	return &Repository{

		App:                  a,
		MinoStorage:          m,
		dailyVerse:           dailyVerse,
		bookService:          booksService,
		eventsService:        eventsService,
		genreService:         genreService,
		jobService:           jobService,
		newsService:          newsService,
		partnersService:      partnersService,
		prayerRequestService: prayerRequestService,
		sermonService:        sermonService,
		testimonyService:     testimonyService,
		userServices:         userServices,
		volunteerService:     volunteerService,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
