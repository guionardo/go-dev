git tag -l --sort=refname "v*" | tail -n 1 | sed -e 's/^v//'
