release-bug:
	bash .github/scripts/new_release.sh bug

release-feature:
	bash .github/scripts/new_release.sh feature

release-major:
	bash .github/scripts/new_release.sh major

build:
	bash -x ./.github/scripts/build.sh
