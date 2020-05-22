###
SHELL := /bin/bash

.PHONY: init
init:
	git init
	git add .
	git remote add origin https://gitlab.mcsolutions.ru/find-psy/configmap.git
	git commit -m "Init"
	git push origin master
.PHONY: commit
commit:
	git add .
	git commit -m "next"
	git push -u origin master
