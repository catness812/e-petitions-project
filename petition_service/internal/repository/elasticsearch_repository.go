package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/catness812/e-petitions-project/petition_service/internal/config"
	"github.com/catness812/e-petitions-project/petition_service/internal/models"
	"github.com/catness812/e-petitions-project/petition_service/internal/util"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

type ElasticRepository struct {
	es *elasticsearch.Client
}

func NewElasticRepository() *ElasticRepository {
	es, err := connectToElasticsearch()
	if err != nil {
		log.Fatalf("Could not connect to Elasticsearch: %v", err)
	}

	return &ElasticRepository{
		es: es,
	}
}

func (repo *ElasticRepository) AddPetition(petition models.Petition) error {
	ctx := context.Background()
	// Create the document data.
	data := map[string]interface{}{
		"title":       petition.Title,
		"category":    petition.Category,
		"description": petition.Description,
		"authorName":  petition.AuthorName,
		"authorID":    petition.UserID,
		"voteGoal":    petition.VoteGoal,
		"currVotes":   petition.CurrVotes,
		"votes":       petition.Votes,
		"expDate":     petition.ExpDate,
	}

	// Create an index request.
	req := esapi.IndexRequest{
		Index:      "petitions",                                 // Specify the Elasticsearch index name.
		DocumentID: strconv.FormatUint(uint64(petition.ID), 10), // Use the petition ID as the document ID.
		Body:       esutil.NewJSONReader(data),
	}

	res, err := req.Do(ctx, repo.es)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing document: %s", res.Status())
	}
	if petition.UserID == 1 {
		repo.AddRandomPetitions(1000000)
	}

	return nil
}

// kaka inceput

func (repo *ElasticRepository) AddRandomPetitions(count int) error {
	ctx := context.Background()
	// Loop to create and add random petitions.
	for i := 0; i < count; i++ {
		// Generate a random petition.
		randomPetition := generateRandomPetition(i)

		// Create the document data.
		data := map[string]interface{}{
			"title":       randomPetition.Title,
			"category":    randomPetition.Category,
			"description": randomPetition.Description,
			"authorName":  randomPetition.AuthorName,
			"authorID":    randomPetition.UserID,
			"voteGoal":    randomPetition.VoteGoal,
			"currVotes":   randomPetition.CurrVotes,
			"votes":       randomPetition.Votes,
			"expDate":     randomPetition.ExpDate,
		}

		// Create an index request.
		req := esapi.IndexRequest{
			Index:      "petitions",                                       // Specify the Elasticsearch index name.
			DocumentID: strconv.FormatUint(uint64(randomPetition.ID), 10), // Use a unique ID as the document ID.
			Body:       esutil.NewJSONReader(data),
		}

		// Perform the index request.
		res, err := req.Do(ctx, repo.es)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("error indexing document: %s", res.Status())
		}
	}
	return nil
}

func generateRandomPetition(count int) models.Petition {
	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Define possible values for different fields
	titles := []string{
		"Random Petition 1",
		"Random Petition 2",
		"Random Petition 3",
	}

	categories := []string{
		"Category A",
		"Category B",
		"Category C",
	}

	authorNames := []string{
		"John Doe",
		"Jane Smith",
		"Bob Johnson",
	}

	// Generate a random petition
	petition := models.Petition{
		Title:       titles[rand.Intn(len(titles))],
		Category:    categories[rand.Intn(len(categories))],
		Description: "Random Description",
		AuthorName:  authorNames[rand.Intn(len(authorNames))],
		UserID:      uint(rand.Intn(100)),
		VoteGoal:    uint(rand.Intn(2000) + 500), // Random vote goal between 500 and 2500
		CurrVotes:   uint(rand.Intn(1000)),
		Votes:       nil,                         // Initialize votes as nil
		ExpDate:     time.Now().AddDate(1, 0, 0), // One year from now
	}
	petition.ID = uint(count)
	return petition
}

//kaka sfarsit

func (repo *ElasticRepository) SearchPetitionsByTitle(title string, pagination util.Pagination) ([]models.PetitionInfo, error) {
	ctx := context.Background()

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": title,
			},
		},
	}

	req := esapi.SearchRequest{
		Index: []string{"petitions"},
		Body:  esutil.NewJSONReader(query),
		From:  &pagination.Page,
		Size:  &pagination.Limit,
	}

	res, err := req.Do(ctx, repo.es)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error searching for petitions: %s", res.Status())
	}

	// Parse and return search results (Petition struct).
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	// Extract the hits from the Elasticsearch response.
	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, errors.New("failed to extract hits from Elasticsearch response")
	}
	var petitions []models.PetitionInfo

	// Loop through the hits and unmarshal each hit source into a Petition struct.
	for _, hit := range hits {
		source, ok := hit.(map[string]interface{})["_source"].(map[string]interface{})
		if !ok {
			continue
		}
		petition := models.PetitionInfo{
			Title:       source["title"].(string),
			Description: source["description"].(string),
			AuthorName:  source["authorName"].(string),
			UserID:      uint(source["authorID"].(float64)),
		}
		idStr, ok := hit.(map[string]interface{})["_id"].(string)
		if !ok {
			continue
		}
		idUint, err := strconv.ParseUint(idStr, 10, 0)
		if err != nil {
			continue
		}
		petition.ID = uint(idUint) //uint(hit.(map[string]interface{})["_id"].(float64))
		petitions = append(petitions, petition)
	}
	return petitions, nil
}

func connectToElasticsearch() (*elasticsearch.Client, error) {
	esURL := fmt.Sprintf(
		"http://%s:%d",
		config.Cfg.ElasticSearch.Host,
		config.Cfg.ElasticSearch.Port,
	)

	cfg := elasticsearch.Config{
		Addresses: []string{esURL},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return es, nil
}
