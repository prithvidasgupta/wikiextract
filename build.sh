#!/usr/bin/env bash

package=github.com/prithvidasgupta/wikiextract

if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi

package_split=(${package//\// })
package_name=${package_split[-1]}


platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

for platform in "${platforms[@]}"
do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}

      output_name=$package_name'-'$GOOS'-'$GOARCH

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi
	
	env GOOS=$GOOS GOARCH=$GOARCH go build -ldflags "-s -w" -o $output_name $package
	
done
