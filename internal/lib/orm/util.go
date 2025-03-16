package orm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	// filterHTTPDelimiter character allows to delimit the field
	// to filter (example: name), the operator (example: contains) and the value (example: hello).
	filterHTTPDelimiter = ":"

	// sortHTTPDelimiter character allows to delimit the field
	// to sort (example: name) and the order (example: asc).
	sortHTTPDelimiter = ":"
)

func Paginate(c *ghttp.Request) (int64, int64, error) {
	pageStr := c.GetQuery("page").String()
	sizeStr := c.GetQuery("size").String()

	page, detailErr := convertStringToInt64(pageStr)
	if detailErr != nil {
		return 0, 0, detailErr
	}

	// Decrement page to avoid having page 0. Starts at 1.
	if page > 0 {
		page--
	}

	size, detailErr := convertStringToInt64(sizeStr)
	if detailErr != nil {
		return 0, 0, detailErr
	}

	// Set a default size value in case no size is specified in HTTP request.
	if size == 0 || size > 1000 {
		size = 100
	}

	return page, size, nil
}

func HttpFilterToORM(c *ghttp.Request) map[string]FieldValue {
	var result = map[string]FieldValue{}

	if c == nil {
		return result
	}

	filter := c.Get("filter").String()
	values := strings.Split(filter, filterHTTPDelimiter)

	// We should have 3 values:
	// - field
	// - operator
	// - value
	if len(values) != 3 {
		return result
	}

	field, operator, value := values[0], values[1], values[2]

	switch operator {
	case "contains":
		result[field] = FieldValue{Operator: "LIKE", Value: "%" + value + "%"}
	case "is":
		result[field] = FieldValue{Operator: "=", Value: value}
	}

	return result
}

func HttpSortToORM(c *ghttp.Request) string {
	if c == nil {
		return ""
	}

	sort := c.Get("sort").String()
	values := strings.Split(sort, sortHTTPDelimiter)

	// We should have 2 values:
	// - field
	// - sort order
	if len(values) != 2 {
		return ""
	}

	field, order := values[0], values[1]

	if order != "asc" && order != "desc" {
		return ""
	}

	return fmt.Sprintf("%s %s", field, order)
}

func convertStringToInt64(value string) (int64, error) {
	if value == "" {
		return 0, nil
	}

	intValue, err := strconv.ParseInt(value, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("unable to convert string to uint64: %v", err)
	}

	return intValue, nil
}
