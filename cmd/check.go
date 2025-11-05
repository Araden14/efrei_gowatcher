package cmd

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Araden14/efrei_gowatcher/internal/checker"
	"github.com/Araden14/efrei_gowatcher/internal/config"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Vérifie l'accessibilité d'une liste d'URLs",
	Long:  "Vérifie l'accessibilité d'une liste d'URLs, en utilisant des goroutinnes",
	Run: func(cmd *cobra.Command, args []string) {
		targets := []string{
			"https://www.google.com",
			"https://www.notarealwebsite.abc",
			"https://github.com",
			"https://www.movie.database/film/details",
			"https://www.gaming.news/release/new-game",
			"https://www.health.clinic/appointment/online",
			"https://www.car.manufacturer/model/electric",
			"https://www.home.decor/ideas/living-room",
			"https://www.environmental.org/project/clean-water",
			"https://www.space.agency/mission/mars",
			"https://www.fashion.magazine/trend/summer",
			"https://www.tech.conference/schedule/day1",
			"https://www.food.blog/recipe/dessert",
			"https://www.online.course/programming/python",
			"https://www.travel.guide/city/paris",
			"https://www.music.label/artist/new-album",
			"https://www.sports.club/events/match",
			"https://www.photography.tips/technique/lighting",
			"https://www.diy.tools/review/drill",
			"https://www.pet.vet/service/vaccination",
			"https://www.gardening.store/seeds/flower",
			"https://www.finance.advice/retirement/planning",
			"https://www.history.podcast/episode/ww2",
			"https://www.language.exchange/partner/find",
			"https://www.book.review/author/classic",
			"https://www.movie.review/genre/comedy",
			"https://www.gaming.forum/topic/strategy",
		}

		var wg sync.WaitGroup

		wg.Add(len(targets))

		for _, url := range targets {
			go func(u string) {
				defer wg.Done()
				target := config.InputTarget{
					Name:  u, // Use URL as name for simplicity
					URL:   u,
					Owner: "unknown", // Default owner
				}
				result := checker.CheckURL(target)

				if result.Err != nil {
					var unreachable *checker.UnreachableError

					if errors.As(result.Err, &unreachable) {
						fmt.Printf("%s est innaccessible : %v\n", unreachable.URL, unreachable.Unwrap())
					} else {
						fmt.Printf("%s : erreur - %v\n", result.InputTarget.URL, result.Err)
					}
				} else {
					fmt.Printf("OK %s - %s\n", result.InputTarget.URL, result.Status)
				}
			}(url)
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
