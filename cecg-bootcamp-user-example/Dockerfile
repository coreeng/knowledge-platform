FROM cecg-bootcamp

# Copies your custom content -- the destination (/site/content/bootcamp/modules) needs to remain the same as that is
# the location of the modules in the base image
COPY modules/user-specific-module-1 /site/content/bootcamp/modules/user-specific-module-1
COPY modules/user-specific-module-2 /site/content/bootcamp/modules/user-specific-module-2

# Copies your config over -- the destination (/site/) needs to remain the same as that is the base hugo location in the base image
COPY customisation/config.toml /site/

# If you'd like to change the logo -- the destination (/site/static/images/logo.png) needs to remain the same as that is
# the location hugo will be looking for a logo, named logo.png
COPY customisation/logo.png /site/static/images/logo.png

# If you'd like to change the default css settings -- the destination (/site/static/css/) needs to remain the same as that is
# the location hugo will be looking for these customisations.

#COPY customisation/theme-example.css /site/static/css/theme-example.css
