package godogs

import (
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-resty/resty/v2"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
	"testing"
)

type TestCounter struct {
	Name    string
	Counter int64
}

var baseUri = getBaseURI()
var request *resty.Request
var response resty.Response
var UUID string

func aRestService() {
	httpClient := resty.New()
	request = httpClient.R()
}

func iCallTheHelloWorldEndpoint() error {
	log.Printf("Hitting GET endpoint %s\n", baseUri)
	httpResponse, err := request.Get(baseUri + "/hello")

	if err != nil {
		return fmt.Errorf("call to %s was unsuccessful, error: %v", baseUri, err)
	}

	response = *httpResponse
	return nil
}

func anOkResponseIsReturned() error {
	if response.IsSuccess() == true {
		return nil
	}
	return fmt.Errorf("response not successful, response code: %d, error: %v", response.StatusCode(), response.Error())
}

func iCallTheGetCounterWithTheNameARandomName(name string) error {
	path := baseUri + "/counter/" + name
	log.Printf("Hitting GET endpoint %s\n", path)
	httpResponse, err := request.Get(path)

	if err != nil {
		return fmt.Errorf("call to %s was unsuccessful, error: %v", baseUri, err)
	}
	response = *httpResponse
	return nil
}

func theResponseBodyIs(responseBody *godog.DocString) error {
	log.Printf("Response body as string is: %s", response.String())
	log.Printf("actual response body: %s", responseBody.Content)
	if !strings.EqualFold(response.String(), responseBody.Content) {
		return fmt.Errorf("expected responseBody : %s did not match actual: %s", responseBody.Content, response.String())
	}
	return nil
}

func aRandomUUID() {
	generatedUUID, _ := uuid.NewV4()
	UUID = generatedUUID.String()
}

func iCallTheGetCounterWithTheRandomUUID() error {
	path := baseUri + "/counter/" + UUID
	log.Printf("Hitting GET endpoint %s\n", path)
	httpResponse, err := request.Get(path)

	if err != nil {
		return fmt.Errorf("call to %s was unsuccessful, error: %v", baseUri, err)
	}
	response = *httpResponse
	return nil
}

func theResponseBodyFieldCounterIsEqualTo(expectedCounterValue int64) error {
	var counter TestCounter

	err := json.Unmarshal(response.Body(), &counter)
	if err != nil {
		return fmt.Errorf("failed while unmarshalling %v", response.String())
	}
	log.Printf("Counter: %v", counter)
	if counter.Counter != expectedCounterValue {
		return fmt.Errorf("expected %v to equal %v", counter.Counter, expectedCounterValue)
	}
	return nil
}

func iCallThePutCounterWithTheRandomUUID() error {
	path := baseUri + "/counter/" + UUID
	log.Printf("Hitting PUT endpoint %s\n", path)
	httpResponse, err := request.Put(path)

	if err != nil {
		return fmt.Errorf("call to %s was unsuccessful, error: %v", baseUri, err)
	}
	response = *httpResponse
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a rest service$`, aRestService)
	ctx.Step(`^an ok response is returned$`, anOkResponseIsReturned)
	ctx.Step(`^I call the hello world endpoint$`, iCallTheHelloWorldEndpoint)
	ctx.Step(`^I call the get counter with the name '(.*)'$`, iCallTheGetCounterWithTheNameARandomName)
	ctx.Step(`^the response body is$`, theResponseBodyIs)
	ctx.Step(`^a random UUID$`, aRandomUUID)
	ctx.Step(`^I call the get counter with the random UUID$`, iCallTheGetCounterWithTheRandomUUID)
	ctx.Step(`^the response body field \'counter\' is equal to \'(\d+)\'$`, theResponseBodyFieldCounterIsEqualTo)
	ctx.Step(`^I call the put counter with the random UUID$`, iCallThePutCounterWithTheRandomUUID)
}

func getBaseURI() string {
	serviceEndpoint := os.Getenv("SERVICE_ENDPOINT")

	if serviceEndpoint == "" {
		return "http://service:8080"
	}
	return serviceEndpoint
}

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}
