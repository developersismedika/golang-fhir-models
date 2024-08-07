#!/usr/bin/env bash

ver=R4

wget -O definitions.zip https://www.hl7.org/fhir/$ver/definitions.json.zip
unzip definitions.zip profiles-resources.json profiles-types.json valuesets.json -d fhir
rm definitions.zip

go generate ./fhir
