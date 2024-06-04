.PHONY: generate_and_copy generate

begin_translations:
	go generate ./internal/translations/translations.go
	find ./internal/translations/locales -name out.gotext.json -execdir cp {} messages.gotext.json \;

complete_translations:
	go generate ./internal/translations/translations.go