package main

import (
	"context"
	"encoding/json"
	_ "expvar"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/stnmrshx/baratum/pb/auth"
	"github.com/stnmrshx/baratum/pb/geo"
	"github.com/stnmrshx/baratum/pb/profile"
	"github.com/stnmrshx/baratum/pb/rate"

	uuid "github.com/nu7hatch/gouuid"

	"github.com/stnmrshx/authtoken"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type inventory struct {
	Hotels    []*profile.Hotel `json:"hotels"`
	RatePlans []*rate.RatePlan `json:"ratePlans"`
}

type client struct {
	auth.AuthClient
	geo.GeoClient
	profile.ProfileClient
	rate.RateClient
}

func requestHandler(c client, w http.ResponseWriter, r *http.Request) {
	tr := trace.New("api.v1", r.URL.Path)
	defer tr.Finish()
	ctx := context.Background()
	ctx = trace.NewContext(ctx, tr)
	if traceID, err := uuid.NewV4(); err == nil {
		ctx = metadata.NewContext(ctx, metadata.Pairs(
			"traceID", traceID.String(),
			"fromName", "api.v1",
		))
	}

	token, err := authtoken.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	if _, err = c.VerifyToken(ctx, &auth.Request{AuthToken: token}); err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// checkin checkout date query params
	inDate, outDate := r.URL.Query().Get("inDate"), r.URL.Query().Get("outDate")
	if inDate == "" || outDate == "" {
		http.Error(w, "Please specify inDate/outDate params", http.StatusBadRequest)
		return
	}

	// nearby
	nearby, err := c.Nearby(ctx, &geo.Request{
		Lat: -7.8011069,
		Lon: 110.3955194,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// reqeusts profiles & rates
	profileCh := getHotelProfiles(c, ctx, nearby.HotelIds)
	rateCh := getRatePlans(c, ctx, nearby.HotelIds, inDate, outDate)

	// wait profiles reply
	profileReply := <-profileCh
	if err := profileReply.err; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// wait rates reply
	rateReply := <-rateCh
	if err := rateReply.err; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// render response
	json.NewEncoder(w).Encode(inventory{
		Hotels:    profileReply.hotels,
		RatePlans: rateReply.ratePlans,
	})
}

type rateResults struct {
	ratePlans []*rate.RatePlan
	err       error
}

func getRatePlans(c client, ctx context.Context, hotelIDs []string, inDate string, outDate string) chan rateResults {
	ch := make(chan rateResults, 1)

	go func() {
		res, err := c.GetRates(ctx, &rate.Request{
			HotelIds: hotelIDs,
			InDate:   inDate,
			OutDate:  outDate,
		})
		ch <- rateResults{res.RatePlans, err}
	}()

	return ch
}

type profileResults struct {
	hotels []*profile.Hotel
	err    error
}

func getHotelProfiles(c client, ctx context.Context, hotelIDs []string) chan profileResults {
	ch := make(chan profileResults, 1)

	go func() {
		res, err := c.GetProfiles(ctx, &profile.Request{
			HotelIds: hotelIDs,
			Locale:   "en",
		})
		ch <- profileResults{res.Hotels, err}
	}()

	return ch
}

// guaranted tcp connection.
func mustDial(addr *string) *grpc.ClientConn {
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
		panic(err)
	}

	return conn
}

func main() {
	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) {
		return true, true
	}

	var (
		port        = flag.String("port", "8080", "Server port")
		authAddr    = flag.String("auth", "auth:8080", "Auth server address format host:port")
		geoAddr     = flag.String("geo", "geo:8080", "Geo server address format of host:port")
		profileAddr = flag.String("profile", "profile:8080", "Pofile server address format host:port")
		rateAddr    = flag.String("rate", "rate:8080", "Rate server address format host:port")
	)
	flag.Parse()

	// client with all grpc connections
	c := client{
		AuthClient:    auth.NewAuthClient(mustDial(authAddr)),
		GeoClient:     geo.NewGeoClient(mustDial(geoAddr)),
		ProfileClient: profile.NewProfileClient(mustDial(profileAddr)),
		RateClient:    rate.NewRateClient(mustDial(rateAddr)),
	}

	// handle http requests
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHandler(c, w, r)
	}))
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
