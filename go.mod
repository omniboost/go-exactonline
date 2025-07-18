module github.com/omniboost/go-exactonline

go 1.24

require (
	github.com/aodin/date v0.0.0-20160219192542-c5f6146fc644
	github.com/gorilla/schema v1.1.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/oauth2 v0.27.0
)

require (
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20191030093734-a170fe1a7240
