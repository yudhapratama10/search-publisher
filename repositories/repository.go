package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	model "github.com/yudhapratama10/search-publisher/models"
)

// var resp map[string]interface{}

func (repo *footballRepository) Insert(footballClub model.FootballClub) (model.FootballClub, error) {

	var footballClubs []model.FootballClub
	var id int

	cursor, err := repo.db.Query(context.Background(), "INSERT INTO footballclub (name, tournaments, nation, hasstadium, description, rating) VALUES ($1, $2, $3, $4, $5, $6)", footballClub.Name, footballClub.Tournaments, footballClub.Nation, footballClub.HasStadium, footballClub.Rating)

	if err != nil {
		return model.FootballClub{}, err
	}

	cursor.Scan(&id)
	footballClub.Id = id
	topic := "test-messages"

	for _, v := range data.Hits.Hits {
		var tournaments []string

		for _, data := range v.Source["tournaments"].([]interface{}) {
			tournaments = append(tournaments, data.(string))
		}

		id, _ := strconv.Atoi(v.Id)

		footballClubs = append(footballClubs, model.FootballClub{
			Id:          id,
			Name:        v.Source["name"].(string),
			Nation:      v.Source["nation"].(string),
			Tournaments: tournaments,
			HasStadium:  v.Source["has_stadium"].(bool),
			Description: v.Source["description"].(string),
			Rating:      v.Source["rating"].(float64),
		})

	}
	//fmt.Println(hits)

	return footballClubs, nil
}

func (repo *footballRepository) Synonym(keyword string, hasStadium bool, page, take int) ([]model.FootballClub, error) {
	return []model.FootballClub{}, nil
}

func (repo *footballRepository) Autocomplete(keyword string) ([]model.FootballClub, error) {
	var footballClubs []model.FootballClub
	var buf bytes.Buffer

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"function_score": map[string]interface{}{
				"query": map[string]interface{}{
					"bool": map[string]interface{}{
						"should": []map[string]interface{}{
							{
								"multi_match": map[string]interface{}{
									"query":     keyword,
									"operator":  "and",
									"type":      "bool_prefix",
									"fuzziness": "AUTO",
									"fields": []string{
										"name",
										"name._2gram",
										"name._3gram",
									},
								},
							},
							{
								"match": map[string]interface{}{
									"name": map[string]interface{}{
										"query":     keyword,
										"fuzziness": "AUTO",
										"operator":  "and",
									},
								},
							},
						},
					},
				},
				"min_score": 0.01,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
	res, err := repo.client.Search(
		//   es.Search.WithContext(context.Background()),
		repo.client.Search.WithIndex("footballclubs"),
		repo.client.Search.WithBody(&buf),
		repo.client.Search.WithTrackTotalHits(true),
		repo.client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
			return []model.FootballClub{}, err
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
			return []model.FootballClub{}, err
		}
	}

	var data model.SearchResult

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
		return []model.FootballClub{}, err
	}

	for _, v := range data.Hits.Hits {
		var tournaments []string

		for _, data := range v.Source["tournaments"].([]interface{}) {
			tournaments = append(tournaments, data.(string))
		}

		id, _ := strconv.Atoi(v.Id)

		footballClubs = append(footballClubs, model.FootballClub{
			Id:          id,
			Name:        v.Source["name"].(string),
			Nation:      v.Source["nation"].(string),
			Tournaments: tournaments,
			HasStadium:  v.Source["has_stadium"].(bool),
			Description: v.Source["description"].(string),
			Rating:      v.Source["rating"].(float64),
		})

	}
	//fmt.Println(hits)

	return footballClubs, nil
}
