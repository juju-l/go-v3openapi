package v3

import (
	/**/ "embed"; "net/http"; "io/fs" /**/
)

//

//go:embed v3
var f embed.FS //
func OpenapiV3()http.Handler {
	m := http.NewServeMux()
	s,e := fs.Sub(f, "v3")
	if e != nil {
	panic(e)
	}
	m.Handle("/swg/openapi/",//*any",
	http.StripPrefix(
	"/swg/openapi/"/*any"*/,http.FileServer(http.FS(s)),/**/
	),
	/**/)
	return m
}

func init() {
  //
}