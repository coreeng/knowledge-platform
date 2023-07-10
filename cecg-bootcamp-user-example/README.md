## CECG Bootcamp

### Hugo

The CECG bootcamp utilizes the [Hugo](https://gohugo.io/) framework, and incorporates
the [learn](https://learn.netlify.app/en/) theme.

This `cecg-bootcamp-user` project serves as an example on how you can extend/configure the Open source CECG bootcamp.

### Structure explained

#### Sample Project layout:

```text
cecg-bootcamp-user
    customization:
        config.toml
        logo.png
        theme-cecg.css
    modules:
        user-specific-module-1
            submodule
                _index.md
                module-1-nesting.md
            _index.md
            background-info.md
            epic-module-1-example.md
            module-1.md
        user-specific-module-2
            _index.md
            background-info.md
            epic-module-2-example.md
            module-2.md
  Dockerfile
```

The bootcamp offers the flexibility to incorporate custom modules such as `user-specific-module-1`
and `user-specific-module-2`.

Additionally, you can leverage the hugo front-matter `weight` field to adjust the ordering of your custom content. More
information around weight and hugo's front-matter can be found [here](https://gohugo.io/content-management/front-matter/)

Inside the customisation folder, you'll discover samples of components that you might wish to alter, such as the logo (`logo.png`) or
the config (`config.toml`) file. Additionally, you'll find a `theme-cecg.css` file, which serves as a customizable CSS file.
This `theme-cecg.css` file is the default one utilized in the base image.

In the [config.toml](customisation/config.toml) file, you'll notice that the themeVariant is currently set to `cecg`. If you wish to 
replace the default `cecg` CSS file with your own custom `.css` file in the `Dockerfile`, it's important to follow the 
naming convention: `theme-<yourThemeName>.css`. Afterwards, you can reference your custom CSS file in the `themeVariant` field
of the `config.toml` by using the name after the `theme-` prefix. For example, if you name your file `theme-example.css`,
you would reference it as `themeVariant: example`.

For further details on configuring the config.toml, refer to the [Learn theme config docs](https://learn.netlify.app/en/basics/configuration/)

### Dockerfile

The [Dockerfile](Dockerfile) builds upon the CECG bootcamp base image by incorporating custom content, as well as
copying a logo and a config files.


### Running locally

You can easily run this project locally by navigating to the Dockerfile location and executing:

`docker build -t <your-tag> .`

`docker run -d -p 8080:8080 <your-tag>`

Your cecg-bootcamp instance will then be available on `localhost:8080`