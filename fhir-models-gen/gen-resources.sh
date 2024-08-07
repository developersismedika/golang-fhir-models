#!/usr/bin/env bash

ver=R4

wget -O definitions.zip https://www.hl7.org/fhir/$ver/definitions.json.zip
unzip definitions.zip profiles-types.json valuesets.json -d fhir
rm definitions.zip
wget -O fhir/bundle.json http://hl7.org/fhir/$ver/bundle.profile.json
wget -O fhir/codesystem.json http://hl7.org/fhir/$ver/codesystem.profile.json
wget -O fhir/structuredefinition.json http://hl7.org/fhir/$ver/structuredefinition.profile.json
wget -O fhir/valueset.json http://hl7.org/fhir/$ver/valueset.profile.json

go generate ./fhir
