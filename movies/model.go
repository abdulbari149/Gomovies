package movies

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Movie struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Year     int       `json:"year"`
	Director *Director `json:"director"`
	Rating   int       `json:"rating"`
	Genre    string    `json:"genre"`
	Isbn     string    `json:"isbn"`
	Poster   string    `json:"poster"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	BirthDate string `json:"birthDate"`
}

var movies []Movie

func Init() {
	f, err := os.Open("data/movies.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	defer f.Close()
	buf := make([]byte, 1024)
	jsonBytes := []byte{}
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}
		if n > 0 {
			jsonBytes = append(jsonBytes, buf[:n]...)
		}
	}

	var moviesData []interface{}

	err = json.Unmarshal(jsonBytes, &moviesData)

	if err != nil {
		log.Fatalf("unable to parse json: %v", err)
	}

	for _, movieData := range moviesData {
		movie := movieData.(map[string]interface{})
		directorData := movie["director"].(map[string]interface{})

		director := &Director{
			FirstName: directorData["firstName"].(string),
			LastName:  directorData["lastName"].(string),
			BirthDate: directorData["birthDate"].(string),
		}

		movieItem := &Movie{
			ID:       movie["id"].(string),
			Title:    movie["title"].(string),
			Year:     int(movie["year"].(float64)),
			Director: director,
			Rating:   int(movie["rating"].(float64)),
			Genre:    movie["genre"].(string),
			Isbn:     movie["isbn"].(string),
			Poster:   movie["poster"].(string),
		}

		movies = append(movies, *movieItem)
	}

}

type MovieRepoImpl struct{}

type MovieRepo interface {
	ListMovies() []Movie
	GetMovie(id string) (*Movie, error)
	CreateMovie(data map[string]interface{}) (*Movie, error)
	UpdateMovie(id string, data map[string]interface{}) (*Movie, error)
	DeleteMovie(id string) error
}

func (m *MovieRepoImpl) ListMovies() []Movie {
	return movies
}

func (m *MovieRepoImpl) GetMovie(id string) (*Movie, error) {
	for _, movie := range movies {
		if movie.ID == id {
			return &movie, nil
		}
	}

	return nil, errors.New("Movie not found")
}

func WriteMoviesToFile() error {
	// write to file, open file

	f, err := os.OpenFile("data/movies.json", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.FileMode(0644))

	if err != nil {
		log.Fatalf("unable to read file: %v", err)

		return err
	}

	defer f.Close()

	// convert movies to json

	moviesJson, err := json.Marshal(movies)

	if err != nil {
		log.Fatalf("unable to parse json: %v", err)

		return err
	}

	// write to file

	_, err = f.Write(moviesJson)

	if err != nil {
		log.Fatalf("unable to write to file: %v", err)

		return err
	}

	return nil
}

func validateMovieData(data map[string]interface{}) error {

	keys := []string{"title", "year", "rating", "genre", "isbn", "director", "poster", "director.firstName", "director.lastName", "director.birthDate"}
	for _, key := range keys {

		if strings.Contains(key, ".") {
			keyParts := strings.Split(key, ".")
			tempData := data
			for i, part := range keyParts {

				if _, ok := tempData[part]; !ok {
					return errors.New(part + " is required")
				}

				if i == len(keyParts)-1 {
					continue
				}

				tempData = tempData[part].(map[string]interface{})
			}

			continue
		}

		if _, ok := data[key]; !ok {
			return errors.New(key + " is required")
		}
	}

	return nil
}

func (m *MovieRepoImpl) CreateMovie(data map[string]interface{}) (*Movie, error) {
	err := validateMovieData(data)
	if err != nil {
		return nil, err
	}

	directorData := data["director"].(map[string]interface{})

	director := &Director{
		FirstName: directorData["firstName"].(string),
		LastName:  directorData["lastName"].(string),
		BirthDate: directorData["birthDate"].(string),
	}

	movie := &Movie{
		ID:       strconv.Itoa(rand.Intn(1000000)),
		Title:    data["title"].(string),
		Year:     int(data["year"].(float64)),
		Rating:   int(data["rating"].(float64)),
		Genre:    data["genre"].(string),
		Poster:   data["poster"].(string),
		Isbn:     data["isbn"].(string),
		Director: director,
	}

	movies = append(movies, *movie)

	// write to file

	err = WriteMoviesToFile()

	if err != nil {
		return nil, errors.New("Failed to save movie to file")
	}

	return movie, nil
}

func (m *MovieRepoImpl) UpdateMovie(id string, data map[string]interface{}) (*Movie, error) {
	movie, err := m.GetMovie(id)

	if err != nil {
		return nil, err
	}

	if title, ok := data["title"]; ok {
		movie.Title = title.(string)
	}

	if isbn, ok := data["isbn"]; ok {
		movie.Isbn = isbn.(string)
	}

	if genre, ok := data["genre"]; ok {
		movie.Genre = genre.(string)
	}

	if rating, ok := data["rating"]; ok {
		movie.Rating = int(rating.(float64))
	}

	if year, ok := data["year"]; ok {
		movie.Year = int(year.(float64))
	}

	if poster, ok := data["poster"]; ok {
		movie.Poster = poster.(string)
	}

	if director, ok := data["director"]; ok {
		directorData := director.(map[string]interface{})

		if firstName, ok := directorData["firstName"]; ok {
			movie.Director.FirstName = firstName.(string)
		}

		if lastName, ok := directorData["lastName"]; ok {
			movie.Director.LastName = lastName.(string)
		}

		if birthDate, ok := directorData["birthDate"]; ok {
			movie.Director.BirthDate = birthDate.(string)
		}
	}

	for i, m := range movies {
		if m.ID == id {
			movies[i] = *movie
		}
	}

	err = WriteMoviesToFile()

	if err != nil {
		return nil, errors.New("Failed to update movie to file")
	}

	return movie, nil
}

func (m *MovieRepoImpl) DeleteMovie(id string) error {
	for i, movie := range movies {
		if movie.ID == id {
			movies = append(movies[:i], movies[i+1:]...)
			return nil
		}
	}

	err := WriteMoviesToFile()

	if err != nil {
		return errors.New("Failed to delete movie to file")
	}

	return errors.New("Movie not found")
}
