#!/bin/bash

../gradlew clean build shadowJar

# Create initialisation script
echo '#!/bin/bash' > initialisation-tool
echo 'java -cp build/libs/initialisation.jar:* io.cecg.initialisation.Main "$@"' >> initialisation-tool

# Make initialisation script executable
chmod +x initialisation-tool
