go-jsonschema \
	--schema-package="https://geojson.org/schema/GeoJSON.json#"=github.com/zeevallin/ninjs \
	--schema-output="https://geojson.org/schema/GeoJSON.json#"=geojson.go \
	--schema-package="http://www.iptc.org/std/ninjs/ninjs-schema_2.1.json#"=github.com/zeevallin/ninjs \
	--schema-output="http://www.iptc.org/std/ninjs/ninjs-schema_2.1.json#"=ninjs.go \
	schemas/geojson.schema.json \
	schemas/ninjs.schema.json