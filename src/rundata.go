package src

type RouteComment struct {
	Route   string
	Comment string
}

type RunData struct {
	WelcomeMessage string
	RouteComments  []RouteComment
}

var welcomeMessage string
var routecomments []RouteComment

func init() {
	welcomeMessage = "Herbalife server on the go..."
	routecomments = []RouteComment{}
}

func AppendRouteComment(route, comment string) {
	routecomment := RouteComment{
		Route:   route,
		Comment: comment,
	}
	routecomments = append(routecomments, routecomment)
}

func GetRunData() RunData {
	return RunData{
		WelcomeMessage: welcomeMessage,
		RouteComments:  routecomments,
	}
}
