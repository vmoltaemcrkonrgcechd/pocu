package sqhelper

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"strings"
)

func MoreLess[T any](sb squirrel.SelectBuilder, sl []T, col string) squirrel.SelectBuilder {
	switch len(sl) {
	case 1:
		sb = sb.Where(fmt.Sprintf("%s >= ?", col), sl[0])
	case 2:
		sb = sb.Where(fmt.Sprintf("%s >= ? AND %s <= ?", col, col), sl[0], sl[1])
	}
	return sb
}

func OrderBy(sb squirrel.SelectBuilder, sl []string) squirrel.SelectBuilder {
	if len(sl) != 0 {
		sb = sb.OrderBy(strings.Join(sl, " "))
	}
	return sb
}

func LimitOffset(sb squirrel.SelectBuilder, limit, offset uint64) squirrel.SelectBuilder {
	if limit != 0 {
		sb = sb.Limit(limit)
	}

	if offset != 0 {
		sb = sb.Offset(offset)
	}

	return sb
}
