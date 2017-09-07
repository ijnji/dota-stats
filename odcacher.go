package odcacher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
)

var (
	qpsLimiter = time.Tick(time.Second)
)

func init() {
	http.Handle("/api/", appHandler(odCacheHandler))
}

type Entity struct {
	Data []byte
}

type appHandler func(http.ResponseWriter, *http.Request) error

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Get OpenDota API results from the datastore, or if not present,
// from OpenDota and put it in the datastore.
func odCacheHandler(w http.ResponseWriter, r *http.Request) error {
	// Get from datastore
	ctx := appengine.NewContext(r)
	p := r.URL.Path
	k := datastore.NewKey(ctx, "od", p, 0, nil)
	e := &Entity{}
	if err := datastore.Get(ctx, k, e); err != nil {
		// Get from OpenDota, retrying if necessary, and wait on a timer
		<-qpsLimiter
		client := urlfetch.Client(ctx)
		resp, err := client.Get("https://api.opendota.com" + p)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New(resp.Status)
		}
		defer resp.Body.Close()
		e.Data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		go datastore.Put(ctx, k, e)
	}
	fmt.Fprintf(w, string(e.Data))
	return nil
}
