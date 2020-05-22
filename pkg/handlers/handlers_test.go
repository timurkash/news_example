package handlers

import (
	"testing"
	//"gitlab.mcsolutions.ru/lib/common/config"
)

const (
	EXPECTED = "123"
)

var (
//LOG_LEVEL  = config.GetEnv("LOG_LEVEL", "info")
//SENTRY_DSN = config.GetEnv("SENTRY_DSN", "")
)

func TestExampleGetRoute(t *testing.T) {
	//req, err := http.NewRequest(http.MethodGet, configmap.NEWS_URL + "/example?value=" + EXPECTED, nil)
	////req = mux.SetURLVars(req, map[string]string{"parameter": "parameter"})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetBodyHandlerFunc(routes.exampleGetFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//body, err := ioutil.ReadAll(rr.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//exampleOut := &models.ExampleOut{}
	//if err = json.Unmarshal(body, exampleOut); err != nil {
	//	t.Fatal(err)
	//}
	//if exampleOut.Example != EXPECTED {
	//	t.Errorf("Expected %s but got %s", EXPECTED, exampleOut.Example)
	//}
}

func TestExamplePostRoute(t *testing.T) {
	//example := &models.Example{Example: EXPECTED}
	//bytes_, err := json.Marshal(example)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//req, err := http.NewRequest(http.MethodPost, configmap.NEWS_URL + "/example", bytes.NewBuffer(bytes_))
	//if err != nil {
	//	t.Fatal(err)
	//}
	//rr := httptest.NewRecorder()
	//routes := Routes{Logger: logger.NewLogger(LOG_LEVEL, SENTRY_DSN)}
	//handler := http.HandlerFunc(hh.GetBodyHandlerFunc(routes.examplePostFunc, routes.Logger))
	//req.Header.Set(consts.CONTENT_TYPE, consts.CONTENT_TYPE_JSON)
	//handler.ServeHTTP(rr, req)
	//if rr.Code != http.StatusOK {
	//	t.Errorf("Expected %d but got %d", http.StatusOK, rr.Code)
	//}
	//body, err := ioutil.ReadAll(rr.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//exampleOut := &models.ExampleOut{}
	//err = json.Unmarshal(body, exampleOut)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if exampleOut.Example != EXPECTED {
	//	t.Errorf("Expected %s but got %s", EXPECTED, exampleOut.Example)
	//}
}
