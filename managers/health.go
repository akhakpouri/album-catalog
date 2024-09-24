package managers

import (
	"fmt"
	"mock-api/model"

	helpers "github.com/akhakpouri/go-helpers/math"
)

var list = []string{"service", "fwd", "mx", "foo", "bar"}
var listNum = helpers.GetRandNumInRange(0, len(list))

func Get() model.Health {
	num := helpers.GetRandNumInRange(1, 10)
	status := ""
	if num%2 != 0 {
		status = "unhealthy"
	} else {
		status = "healthy"
	}

	details := []model.Detail{}
	for i := 0; i <= num; i++ {
		details = append(details, model.Detail{
			Key:         fmt.Sprintf("%s api", list[listNum]),
			Description: getDesc(),
			Status:      status,
		})
	}

	health := model.Health{Status: status, Details: details}
	return health
}

func getDesc() string {
	status := "healthy"
	if listNum%2 == 1 {
		status = "unhealthy"
	}

	return fmt.Sprintf(`%s api is %s`, list[listNum], status)
}
