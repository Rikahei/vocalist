package youtubeApi

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
	"google.golang.org/grpc"
)

const missingClientSecretsMessage = `
Please configure OAuth 2.0
`

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func getTokenFromWeb(config *oauth2.Config, code string) *oauth2.Token {

	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}
	return tok
}

func getTokenUrl(config *oauth2.Config) string {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func handleError(err error, message string) {
	if message == "" {
		message = "Error making API call"
	}
	if err != nil {
		log.Fatalf(message+": %v", err.Error())
	}
}

func channelsListMine(service *youtube.Service, part string) *youtube.ChannelListResponse {
	call := service.Channels.List(part)
	call = call.Mine(true)
	response, err := call.Do()
	handleError(err, "")
	fmt.Println(fmt.Sprintf("This channel's ID is %s. Its title is '%s', Its description is '%s' "+
		"and it has %d views.",
		response.Items[0].Id,
		response.Items[0].Snippet.Title,
		response.Items[0].Snippet.Description,
		response.Items[0].Statistics.ViewCount))

	return response
}

type MineChannel struct {
	MineId      string `json:"mine_id,omitempty"`
	MineTitle   string `json:"mine_title,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
	Expiry      string `json:"expiry,omitempty"`
}

var (
	dgraph = flag.String("d", "dgraph-db:9080", "Dgraph Alpha address")
)

func hasUserChannelId(channelId string) bool {

	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	println("Check in DB")

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	resp, err := dg.NewTxn().Query(context.Background(), `{
		acc (func: eq(mine_id, `+channelId+`)) {
			uid
			mine_id
			mine_title
			access_token
			token_type
		}
	}`)

	var decode struct {
		Acc []struct {
			Uid         string `json:"uid,omitempty"`
			MineTitle   string `json:"mine_title,omitempty"`
			AccessToken string `json:"access_token,omitempty"`
			TokenType   string `json:"token_type,omitempty"`
		}
	}

	if err := json.Unmarshal(resp.GetJson(), &decode); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("decode : %s\n", decode.Acc[0].Uid)
	if resp.Json != nil {
		return true
	} else {
		return false
	}
}

func saveMineToken(mineResponse *youtube.ChannelListResponse, tokenSet *oauth2.Token) {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	mineSet := MineChannel{
		MineId:      mineResponse.Items[0].Id,
		MineTitle:   mineResponse.Items[0].Snippet.Title,
		AccessToken: tokenSet.AccessToken,
		TokenType:   tokenSet.TokenType,
		Expiry:      tokenSet.Expiry.String(),
	}

	mu := &api.Mutation{
		CommitNow: true,
	}
	pb, err := json.Marshal(mineSet)
	if err != nil {
		log.Fatal(err)
	}

	mu.SetJson = pb
	resp, err := dg.NewTxn().Mutate(context.Background(), mu)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ADD channel ID to DB: ", resp)

}

func updateMineToken(mineResponse *youtube.ChannelListResponse, tokenSet *oauth2.Token) {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	ctx := context.Background()

	q := `
		query{
			acc as var (func: eq(mine_id, ` + mineResponse.Items[0].Id + `))
		}`

	mu := &api.Mutation{
		SetNquads: []byte(`uid(acc) <access_token> "` + tokenSet.AccessToken + `" .
			uid(acc) <expiry> "` + tokenSet.Expiry.String() + `" .`),
	}

	req := &api.Request{
		Query:     q,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}
	if _, err := dg.NewTxn().Do(ctx, req); err != nil {
		log.Fatal(err)
	}
}

func GetMineInfo(code string) *youtube.ChannelListResponse {
	ctx := context.Background()

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tok := getTokenFromWeb(config, code)
	client := config.Client(ctx, tok)
	service, err := youtube.New(client)

	mineResponse := channelsListMine(service, "snippet,contentDetails,statistics")
	handleError(err, "Error creating YouTube client")

	HasChannelId := hasUserChannelId(mineResponse.Items[0].Id)

	println("HasChannelId is : ", HasChannelId)

	if HasChannelId {
		println("User had data")
		println(tok.AccessToken)
		updateMineToken(mineResponse, tok)
	} else {
		saveMineToken(mineResponse, tok)
	}

	return mineResponse
}

func GetUrl() string {

	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/youtube-go-quickstart.json
	config, err := google.ConfigFromJSON(b, youtube.YoutubeReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	url := getTokenUrl(config)
	return url
}
