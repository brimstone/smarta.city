.PHONY: deploy
ifeq (${TRAVIS_PULL_REQUEST},false)
ifeq (${TRAVIS_BRANCH},master)
deploy:
	wget -qO- https://toolbelt.heroku.com/install-ubuntu.sh | sh
	heroku plugins:install heroku-container-registry
	heroku container:login
	docker push ${IMAGE_TAG}
else
deploy:
	@echo "Travis won't deploy when branch is not master"
endif
else
deploy:
	@echo "Travis won't deploy during pull requests"
endif

include ${PROJECTBUILDER}/Makefile
