package csv

import (
	"bufio"
	"fmt"
	"os"
	"parser/internal/domain"
	"strings"
)

func SaveInfluencers(influencers []domain.Influencer, filename string) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintln(writer, "Rank;Name;Nick;Category;Followers;Country;Eng. (Auth.);Eng. (Avg.)")
	for _, influencer := range influencers {
		row := fmt.Sprintf(
			"%s;%s;%s;%s;%s;%s;%s;%s",
			influencer.Rank,
			influencer.Name,
			influencer.Nick,
			stringSliceToString(influencer.Categories),
			influencer.Subscribes,
			influencer.Audience,
			influencer.Authentic,
			influencer.Engagement,
		)
		fmt.Fprintln(writer, row)
	}
	return nil
}

func stringSliceToString(stringSlice []string) string {
	sb := strings.Builder{}
	for i, s := range stringSlice {
		sb.WriteString(s)
		if i != len(stringSlice)-1 {
			sb.WriteString(", ")
		}
	}
	return "\"" + strings.Replace(sb.String(), "\n", "", -1) + "\""
}
