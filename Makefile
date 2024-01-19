.PHONY: download-deps run translations frontend

all: download-deps translations frontend run

download-deps:
	go mod download

run:
	go run main.go serve

translations:
	mkdir -p translations/merged
	go run -mod=mod github.com/nicksnyder/go-i18n/goi18n merge -sourceLanguage de-de -format yaml -outdir translations/merged/ translations/src/*.json

frontend:
	cd frontend \
		&& npm ci \
		&& npx flamingo-carotene build
