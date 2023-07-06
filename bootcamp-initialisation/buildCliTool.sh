#!/bin/bash

./gradlew clean build shadowJar

# Create bootcamp-initialize script
echo '#!/bin/bash' > bootcamp-initialize
echo 'java -cp build/libs/bootcamp-initialize.jar:* io.cecg.bootcamp.initialisation.Main "$@"' >> bootcamp-initialize

# Make bootcamp-initialize script executable
chmod +x bootcamp-initialize
