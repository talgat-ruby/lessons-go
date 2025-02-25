package shared

type SortByDirection string

const (
	SortByDirectionUnspecified    SortByDirection = "Unspecified"
	SortByDirectionAsc            SortByDirection = "Asc"
	SortByDirectionDesc           SortByDirection = "Desc"
	SortByDirectionAscNullsFirst  SortByDirection = "AscNullsFirst"
	SortByDirectionAscNullsLast   SortByDirection = "AscNullsLast"
	SortByDirectionDescNullsFirst SortByDirection = "DescNullsFirst"
	SortByDirectionDescNullsLast  SortByDirection = "DescNullsLast"
)

var SortByDirections = []SortByDirection{
	SortByDirectionUnspecified,
	SortByDirectionAsc,
	SortByDirectionDesc,
	SortByDirectionAscNullsFirst,
	SortByDirectionAscNullsLast,
	SortByDirectionDescNullsFirst,
	SortByDirectionDescNullsLast,
}

type SortBy[Field ~string] interface {
	GetSortBy() []SortByItem[Field]
}

type SortByItem[Field ~string] interface {
	GetDirection() SortByDirection
	GetField() Field
}
