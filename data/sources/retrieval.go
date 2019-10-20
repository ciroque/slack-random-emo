package sources

import "slack-random-emo/data"

type Retrieval interface {
	Retrieve() []data.Emo
}
