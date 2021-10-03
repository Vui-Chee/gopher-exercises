package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		shortUrl := r.RequestURI
		if longUrl, ok := pathsToUrls[shortUrl]; ok {
			http.Redirect(w, r, longUrl, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func parseYaml(data []byte) ([]pathUrl, error) {
	pathUrls := []pathUrl{}
	if err := yaml.Unmarshal(data, &pathUrls); err != nil {
		return nil, err
	}
	return pathUrls, nil
}

func YAMLHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(data)
	if err != nil {
		return nil, err
	}

	pathsToUrls := map[string]string{}
	for _, pathUrl := range pathUrls {
		pathsToUrls[pathUrl.Path] = pathUrl.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}
