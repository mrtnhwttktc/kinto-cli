.PHONY: generate_and_copy generate

generate_and_copy:
	go generate ./internal/translations/translations.go
	find ./internal/translations/locales -name out.gotext.json -execdir cp {} messages.gotext.json \;

generate:
	go generate ./internal/translations/translations.go