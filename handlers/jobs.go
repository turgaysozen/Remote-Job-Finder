package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"github.com/remote-job-finder/service/rss"
	"github.com/remote-job-finder/utils/logger"
	"github.com/remote-job-finder/utils/redis"
)

func JobsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	keys, _ := redis.RedisClient.LRange(ctx, "categories", 0, -1).Result()
	logger.Info.Println(
		"Key fetched from redis and jobs are fething from the cache,",
		"keys:", keys,
	)

	var jobs []rss.Channel
	for _, key := range keys {
		jobData, _ := redis.GetJobs(ctx, key)

		var job rss.Channel
		err := json.Unmarshal(jobData, &job)
		if err == nil {
			jobs = append(jobs, job)
		}
	}

	jobsByte, _ := json.Marshal(jobs)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobsByte)
}

func JobDetailsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, slug string) {
	logger.Info.Println("Handling slug:", slug)

	slugArr := strings.Split(slug, "--")
	jobData, _ := redis.GetJobs(ctx, slugArr[0])

	var job rss.Channel
	var jobDetaByte []byte
	err := json.Unmarshal(jobData, &job)
	if err == nil {
		for _, job := range job.Jobs {
			slug := createSlug(job.Title)
			if slug == slugArr[1] {
				logger.Info.Println("Target job found for slug:", slug)
				jobDetaByte, _ = json.Marshal(job)
				break
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jobDetaByte)
}

func createSlug(title string) string {
	slug := strings.ToLower(title)

	reg := regexp.MustCompile(`[^\w\s-]`)
	slug = reg.ReplaceAllString(slug, "")

	reg = regexp.MustCompile(`\s+`)
	slug = reg.ReplaceAllString(slug, "-")

	return slug
}
