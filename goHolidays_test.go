package goHolidays

import (
	"errors"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
	"os"
	"strconv"
	"testing"
	"time"
)

var opt = godog.Options{
	Concurrency: 5,
	Format:      "progress",
	Output:      colors.Colored(os.Stdout),
	Randomize:   1,
}

func init() {
	godog.BindCommandLineFlags("cadence.", &opt)
}

func TestCucumberFeatures(t *testing.T) {
	pflag.Parse()
	opt.Paths = pflag.Args()

	suite := godog.TestSuite{
		Name:                "cadence",
		ScenarioInitializer: InitializeScenario,
		Options:             &opt,
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run cucumber tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {

	//////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Holiday Object Tests                                                                                 //
	//////////////////////////////////////////////////////////////////////////////////////////////////////////

	htd := HolidayTestData{}

	// Given Mappings
	sc.Step(`^I have a NewYears Holiday Object$`, htd.iHaveANewYearsHolidayObject)
	sc.Step(`^I have a MLKHoliday Holiday Object$`, htd.iHaveAMLKHolidayHolidayObject)
	sc.Step(`^I have a PresidentsDay Holiday Object$`, htd.iHaveAPresidentsDayHolidayObject)
	sc.Step(`^I have a MemorialDay Holiday Object$`, htd.iHaveAMemorialDayHolidayObject)
	sc.Step(`^I have a Juneteenth Holiday Object$`, htd.iHaveAJuneteenthHolidayObject)
	sc.Step(`^I have a IndependenceDay Holiday Object$`, htd.iHaveAIndependenceDayHolidayObject)
	sc.Step(`^I have a LaborDay Holiday Object$`, htd.iHaveALaborDayHolidayObject)
	sc.Step(`^I have a ColumbusDay Holiday Object$`, htd.iHaveAColumbusDayHolidayObject)
	sc.Step(`^I have a VeteransDay Holiday Object$`, htd.iHaveAVeteransDayHolidayObject)
	sc.Step(`^I have a Thanksgiving Holiday Object$`, htd.iHaveAThanksgivingHolidayObject)
	sc.Step(`^I have a Christmas Holiday Object$`, htd.iHaveAChristmasHolidayObject)

	// When Mappings
	sc.Step(`^I check "([^"]*)"$`, htd.iCheck)

	// Then Mappings
	sc.Step(`^the description will be "([^"]*)"$`, htd.theDescriptionWillBe)
	sc.Step(`^the name will be "([^"]*)"$`, htd.theNameWillBe)

	sc.Step(`^it returns true$`, htd.itReturnsTrue)
	sc.Step(`^it will return false$`, htd.itWillReturnFalse)

	//////////////////////////////////////////////////////////////////////////////////////////////////////////
	// Holiday List Object Tests                                                                            //
	//////////////////////////////////////////////////////////////////////////////////////////////////////////

	hltd := HolidayListTestData{}

	// Given Mappings
	sc.Step(`^I have a USFederalHoliday HolidayList Object$`, hltd.iHaveAUSFederalHolidayHolidayListObject)

	// When Mappings
	sc.Step(`^I execute check with "([^"]*)"$`, hltd.iExecuteCheckWith)

	// Then Mappings
	sc.Step(`^it returns "([^"]*)"$`, hltd.itReturns)
	sc.Step(`^the title will be "([^"]*)"$`, hltd.theTitleWillBe)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
// Holiday Object Tests                                                                                 //
//////////////////////////////////////////////////////////////////////////////////////////////////////////

type HolidayTestData struct {
	holiday      Holiday
	CheckResult  bool
	StringResult string
}

// Given Callbacks
func (td *HolidayTestData) iHaveANewYearsHolidayObject() (err error) {
	td.holiday = NewYearsDay

	return
}

func (td *HolidayTestData) iHaveAMLKHolidayHolidayObject() (err error) {
	td.holiday = MLKHoliday

	return
}

func (td *HolidayTestData) iHaveAPresidentsDayHolidayObject() (err error) {
	td.holiday = PresidentsDay

	return
}

func (td *HolidayTestData) iHaveAMemorialDayHolidayObject() (err error) {
	td.holiday = MemorialDay

	return
}

func (td *HolidayTestData) iHaveAJuneteenthHolidayObject() (err error) {
	td.holiday = Juneteenth

	return
}

func (td *HolidayTestData) iHaveAIndependenceDayHolidayObject() (err error) {
	td.holiday = IndependenceDay

	return
}

func (td *HolidayTestData) iHaveALaborDayHolidayObject() (err error) {
	td.holiday = LaborDay

	return
}

func (td *HolidayTestData) iHaveAColumbusDayHolidayObject() (err error) {
	td.holiday = ColumbusDay

	return
}

func (td *HolidayTestData) iHaveAVeteransDayHolidayObject() (err error) {
	td.holiday = VeteransDay

	return
}

func (td *HolidayTestData) iHaveAThanksgivingHolidayObject() (err error) {
	td.holiday = Thanksgiving

	return
}

func (td *HolidayTestData) iHaveAChristmasHolidayObject() (err error) {
	td.holiday = Christmas

	return
}

// When Callbacks
func (td *HolidayTestData) iCheck(dateString string) (err error) {
	dateValue, err := time.Parse("2006-01-02", dateString)
	td.CheckResult = td.holiday.Check(dateValue)

	return
}

// Then Callbacks
func (td *HolidayTestData) itReturnsTrue() (err error) {
	if !td.CheckResult {
		err = errors.New("ERROR: HOLIDAY: Did not detect holiday correctly")
	}

	return
}

func (td *HolidayTestData) itWillReturnFalse() (err error) {
	if td.CheckResult {
		err = errors.New("ERROR: HOLIDAY: False positive, failed to error on incorrect date")
	}

	return
}

func (td *HolidayTestData) theDescriptionWillBe(expected string) (err error) {
	if td.holiday.Describe() != expected {
		err = errors.New(fmt.Sprintf("ERROR: HOLIDAY: Name was not as expected, received %s, expected: %s", td.holiday.Describe(), expected))
	}

	return
}

func (td *HolidayTestData) theNameWillBe(expected string) (err error) {
	if td.holiday.Name() != expected {
		err = errors.New(fmt.Sprintf("ERROR: HOLIDAY: Description was not as expected, received %s, expected: %s", td.holiday.Name(), expected))
	}

	return
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////
// Holiday List Object Tests                                                                            //
//////////////////////////////////////////////////////////////////////////////////////////////////////////

type HolidayListTestData struct {
	holidays Holidays
	checkRes bool
}

// Given Callbacks
func (hltd HolidayListTestData) iHaveAUSFederalHolidayHolidayListObject() (err error) {
	hltd.holidays = USFederalHolidays

	return
}

// When Callbacks
func (hltd HolidayListTestData) iExecuteCheckWith(dateString string) (err error) {
	dateValue, err := time.Parse("2006-01-02", dateString)
	hltd.checkRes = hltd.holidays.Check(dateValue)

	return
}

// Then Callbacks
func (hltd HolidayListTestData) itReturns(boolStr string) (err error) {
	boolVal, err := strconv.ParseBool(boolStr)
	if hltd.checkRes != boolVal {
		err = errors.New(fmt.Sprintf("ERROR: HOLIDAYLIST: Incorrect return from check"))
	}

	return
}

func (hltd HolidayListTestData) theTitleWillBe(expected string) (err error) {
	if hltd.holidays.Name() != expected {
		err = errors.New(fmt.Sprintf("ERROR: HOLIDAYLIST: Incorrect return from Name()"))
	}

	return
}
