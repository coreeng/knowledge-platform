#!/bin/bash

./gradlew clean build shadowJar

# Create bootcamp-initialize script
echo '#!/bin/bash' > bootcamp-initialize
echo 'java -cp build/libs/bootcamp-initialize.jar:* io.cecg.bootcamp.initialisation.Main "$@"' >> bootcamp-initialize

# Make bootcamp-initialize script executable
chmod +x bootcamp-initialize

output=$(./bootcamp-initialize --git-token=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 --org=coreeng --modules=p2p-fast-feedback --bootcampee-repo=bootcamp-initialisation-smoke-test)
echo $output

github_link=$(echo "$output" | grep -o "https://github.com/[[:alnum:]_-]*/[[:alnum:]_-]*/issues/")
# Extract repository name from the GitHub Link
repository=$(echo "$github_link" | awk -F'/' '{print $4 "/" $5}')

# GitHub API request to count issues
issues_count=$(curl -s -H "Authorization: Bearer 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" "https://api.github.com/repos/$repository/issues?state=all" | jq length)
echo "Number of issues: $issues_count"
