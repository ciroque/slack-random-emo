module slack-random-emo.org/slack-random-emo

go 1.19

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	unit.nginx.org/go v0.0.0-20221215124201-37cac7fec92e // indirect
)

require (
	github.com/Sirupsen/logrus v1.0.6
	slack-random-emo.org/config v0.0.0
	slack-random-emo.org/data v0.0.0
	slack-random-emo.org/server v0.0.0
	slack-random-emo.org/metrics v0.0.0
)

replace (
	slack-random-emo.org/config => ./config
	slack-random-emo.org/data => ./data
	slack-random-emo.org/http => ./server
	slack-random-emo.org/metrics => ./metrics
)
